# go-dynamodb-sample

# init
**pull DynamoDB Local**

```shell
$ docker pull amazon/dynamodb-local
```

**run DynamoDB container**
```
$ docker run -d --name dynamodb -p 8000:8000 amazon/dynamodb-local
```

**create table**

```shell
$ aws dynamodb create-table --endpoint-url http://localhost:8000 --table-name MyFirstTable --attribute-definitions AttributeName=MyHashKey,AttributeType=S AttributeName=MyRangeKey,AttributeType=N --key-schema AttributeName=MyHashKey,KeyType=HASH AttributeName=MyRangeKey,KeyType=RANGE --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1
```

# Reference
https://github.com/guregu/dynamo#readme<br>
https://future-architect.github.io/articles/20200225/
