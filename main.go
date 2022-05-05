package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func main() {
	dynamoDbRegion := "ap-northeast-1"
	dynamoDbEndpoint := "http://localhost:8000"
	db := dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(true),
	})

	table := db.Table("MyFirstTable")

	// Create Item
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "MyFirstText",
	}
	if err := table.Put(item).Run(); err != nil {
		fmt.Printf("faile to put item %v\n", err)
	}

	// Read Item
	var readItem Item
	err := table.Get("MyHashKey", item.MyHashKey).Range("MyRangeKey", dynamo.Equal, item.MyRangeKey).One(&readItem)
	if err != nil {
		fmt.Printf("failed to get item %v\n", err)
	}

	fmt.Printf("readItem.MyHashKey: %s, readItem.MyRangeKey: %v, readItem.MyText: %s\n", readItem.MyHashKey, readItem.MyRangeKey, readItem.MyText)

	// Update Item
	var updateItem Item
	err = table.Update("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Set("MyText", "My Update Text").Value(&updateItem)
	if err != nil {
		fmt.Printf("failed to get item %v\n", err)
	}
	fmt.Printf("updateItem.MyHashKey: %s, updateItem.MyRangeKey: %v, updateItem.MyText: %s\n", updateItem.MyHashKey, updateItem.MyRangeKey, updateItem.MyText)

	// Delete Item
	err = table.Delete("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Run()
	if err != nil {
		fmt.Printf("Failed to delete item %v\n", err)
	}
}

type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}
