package TaskList

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const TableName string = "Home-Dashboard"

func DBCreate(t any) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(t)
	if err != nil {
		log.Fatalf("Failed to marshal task item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Error calling PutItem: %s", err)
	}

	fmt.Println("Success: Added new item to table " + TableName)
}

func DBRead(userID string, section string) *dynamodb.QueryOutput {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	keyCond := expression.KeyAnd(expression.Key("UserID").Equal(expression.Value(userID)), expression.Key("SK").BeginsWith(section))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		log.Fatalf("Failed to build expression: %s", err)
	}

	input := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		TableName:                 aws.String(TableName),
	}

	result, err := svc.Query(input)
	if err != nil {
		log.Fatalf("Failed to query database: %s", err)
	}

	return result
}

func DBDelete(userID string, SK string) error {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				S: aws.String(userID),
			},
			"SK": {
				S: aws.String(SK),
			},
		},
		TableName:    aws.String(TableName),
		ReturnValues: aws.String("ALL_OLD"),
	}

	result, err := svc.DeleteItem(input)
	if err != nil {
		return err
	}
	fmt.Println("Success: Deleted " + result.String() + "from table " + TableName)

	return nil
}
