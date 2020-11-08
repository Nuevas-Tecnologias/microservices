package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	_ "github.com/go-sql-driver/mysql"
)

type MessageType struct {
	Type string `json:"type""`
}
type NewRequest struct {
	Name          string `json:"name""`
	CorrelationId int    `json:"correlationId""`
}
type FormRegistered struct {
	Id            int `json:"id""`
	CorrelationId int `json:"correlationId""`
}
type OrderData struct {
	OrderId       int `json:"techOrderId""`
	CorrelationId int `json:"correlationId""`
}
type ServiceData struct {
	ServiceIds    []int `json:"serviceIds""`
	CorrelationId int   `json:"correlationId""`
}
type MessageEvent struct {
	Id            int
	CorrelationId int
	Status        string
	Data          string
	Error         string
	CreatedAt     string
	UpdatedAt     string
}

const (
	//AWS Queues
	ServiceCenter   = "https://sqs.us-west-2.amazonaws.com/881619806726/assign-service-center-command.fifo"
	OrderService    = "https://sqs.us-west-2.amazonaws.com/881619806726/assign-service-center-command.fifo"
	RegisterService = "https://sqs.us-west-2.amazonaws.com/881619806726/create-service-register-bulk-command.fifo"
	Region          = "us-west-2"
	//DB credentials
	databaseHost     = "terraform-20201028182302976100000001.cgrpasjjlw1k.us-west-2.rds.amazonaws.com"
	databasePort     = "3306"
	databaseUser     = "newarchitectures"
	databasePassword = "newarchitectures"
	databaseName     = "orchestrator"
)

//Lambda handler
func handler(ctx context.Context, sqsEvent events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	if len(sqsEvent.Records) == 0 {
		return returnResponse(400, fmt.Sprintf("No SQS message passed to function.")), nil
	}
	messageType := &MessageType{}
	for _, msg := range sqsEvent.Records {
		err := json.Unmarshal([]byte(msg.Body), messageType)
		messageData, err := getMessageData(*messageType, msg.Body)
		if err != nil {
			return returnResponse(400, fmt.Sprintf("Error on strategy. %s", err.Error())), nil
		}
		return messageData, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello world"),
		StatusCode: 200,
	}, nil
}

//Strategies
func getMessageData(mt MessageType, msg string) (events.APIGatewayProxyResponse, error) {
	rsp := events.APIGatewayProxyResponse{}
	var err error
	var a string
	switch origin := mt.Type; origin {
	case "TechRequestReceived":
		err = newRequestStrategy(msg)
		if err != nil {
			a = err.Error()
		} else {
			a = "todo bien"
		}
		rsp = returnResponse(200, fmt.Sprintf("New request origin. %s", a))
	case "TechServiceCenter":
		err = formStrategy(msg)
		if err != nil {
			a = err.Error()
		} else {
			a = "todo bien"
		}
		rsp = returnResponse(200, fmt.Sprintf("Form origin. %s", a))
	case "TechOrderCreated":
		err = orderStrategy(msg)
		if err != nil {
			a = err.Error()
		} else {
			a = "todo bien"
		}
		rsp = returnResponse(200, fmt.Sprintf("Orders origin. %s", a))
	case "TechServiceCreated":
		err := serviceStrategy(msg)
		if err != nil {
			a = err.Error()
		} else {
			a = "todo bien"
		}
		rsp = returnResponse(200, fmt.Sprintf("Service origin. %s", a))
	default:
		rsp = returnResponse(400, fmt.Sprintf("No data."))
	}
	return rsp, err
}

//New Request incoming
func newRequestStrategy(msg string) error {
	messageData := NewRequest{}
	json.Unmarshal([]byte(msg), &messageData)
	lastId, rowCnt, err := insertMessageEvent("CREATED", "New request incoming", 0)
	if err != nil {
		return err
	}
	ids := fmt.Sprintf("%v - %v", lastId, rowCnt)
	fmt.Println("Ids: ", ids)
	err = updateMessageEventCorrelationId(lastId, lastId)
	if err != nil {
		return err
	}
	//err = writeMessageOnQueue(RegisterService, "data")
	if err != nil {
		return err
	}
	return nil
}

//Form request created
func formStrategy(msg string) error {
	messageData := FormRegistered{}
	json.Unmarshal([]byte(msg), &messageData)
	ans, err := searchMessageLog(messageData.CorrelationId)
	fmt.Println("Type: ", ans.Id)
	if err != nil {
		updateMessageEvent(messageData.CorrelationId, "ERROR", err.Error())
		return err
	} else if ans.Status != "CREATED" && ans.Status != "SERCENASSIGNED" {
		err = errors.New("Record has an invalid status")
		updateMessageEvent(messageData.CorrelationId, "ERROR", err.Error())
		return err
	} else if ans.Status == "SERCENASSIGNED" {
		err = errors.New("Record already in the current status")
		return err
	}
	//err = writeMessageOnQueue(RegisterService, "data")
	if err != nil {
		return err
	}
	lastId, rowCnt, err := insertMessageEvent("SERCENASSIGNED", "New request incoming", messageData.CorrelationId)
	if err != nil {
		return err
	}
	ids := fmt.Sprintf("%v - %v", lastId, rowCnt)
	fmt.Println("Ids: ", ids)
	//updateMessageEvent(messageData.CorrelationId, "SERCENASSIGNED", " ")
	return nil
}

//Order created
func orderStrategy(msg string) error {
	messageData := OrderData{}
	json.Unmarshal([]byte(msg), &messageData)
	ans, err := searchMessageLog(messageData.CorrelationId)
	if err != nil {
		updateMessageEvent(messageData.CorrelationId, "ERROR", err.Error())
		return err
	} else if ans.Status != "SERCENASSIGNED" && ans.Status != "ORDERCREATED" {
		err = errors.New("Record has an invalid status")
		updateMessageEvent(messageData.CorrelationId, "ERROR", err.Error())
		return err
	} else if ans.Status == "ORDERCREATED" {
		err = errors.New("Record already in the current status")
		return err
	}
	//err = writeMessageOnQueue(RegisterService, "data")
	if err != nil {
		updateMessageEvent(messageData.CorrelationId, "ERROR", err.Error())
		return err
	}
	lastId, rowCnt, err := insertMessageEvent("ORDERCREATED", "New request incoming", messageData.CorrelationId)
	if err != nil {
		return err
	}
	ids := fmt.Sprintf("%v - %v", lastId, rowCnt)
	fmt.Println("Ids: ", ids)
	//updateMessageEvent(messageData.CorrelationId, "ORDERCREATED", " ")
	return nil
}

//Service created
func serviceStrategy(msg string) error {
	messageData := ServiceData{}
	json.Unmarshal([]byte(msg), &messageData)
	ans, err := searchMessageLog(messageData.CorrelationId)
	if err != nil {
		updateMessageEvent(messageData.CorrelationId, "ERROR", err.Error())
		return err
	} else if ans.Status != "ORDERCREATED" && ans.Status != "SERVICECREATED" {
		err = errors.New("Record has an invalid status")
		updateMessageEvent(messageData.CorrelationId, "ERROR", err.Error())
		return err
	} else if ans.Status == "SERVICECREATED" {
		err = errors.New("Record already in the current status")
		return err
	}
	//err = writeMessageOnQueue(RegisterService, "data")
	if err != nil {
		err = updateMessageEvent(messageData.CorrelationId, "ERROR", err.Error())
		return err
	}
	lastId, rowCnt, err := insertMessageEvent("SERVICECREATED", "New request incoming", messageData.CorrelationId)
	if err != nil {
		return err
	}
	ids := fmt.Sprintf("%v - %v", lastId, rowCnt)
	fmt.Println("Ids: ", ids)
	//updateMessageEvent(messageData.CorrelationId, "SERVICECREATED", " ")
	return nil
}

//General response
func returnResponse(code int, message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: code,
	}
}

//Write message on queue sqs
func writeMessageOnQueue(QueueUrl string, message string) error {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// session := session.New(&aws.Config{
	// 	Region:      aws.String(Region),
	// 	Credentials: credentials.NewSharedCredentials(CredPath, CredProfile),
	// 	MaxRetries:  aws.Int(5),
	// })
	// session, err := session.NewSession(&aws.Config{
	// 	Region:      aws.String("us-east-1"),
	// 	Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
	// })
	svc := sqs.New(session)
	// Send message
	send_params := &sqs.SendMessageInput{
		MessageBody:  aws.String(message),  // Required
		QueueUrl:     aws.String(QueueUrl), // Required
		DelaySeconds: aws.Int64(3),         // (optional) 傳進去的 message 延遲 n 秒才會被取出, 0 ~ 900s (15 minutes)
	}
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("[Send message] \n%v \n\n", send_resp)

	return nil
}

//Search into DB by id
func searchMessageLog(id int) (MessageEvent, error) {
	messageEvent := MessageEvent{}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	connection, err := sql.Open("mysql", connectionString)
	if err != nil {
		return messageEvent, err
	}
	rows, err := connection.Query("SELECT * FROM `event_status` WHERE `correlation_id` = ? ORDER BY `id` DESC LIMIT 1;", id)
	if err != nil {
		return messageEvent, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&messageEvent.Id, &messageEvent.CorrelationId, &messageEvent.Status, &messageEvent.Data, &messageEvent.Error, &messageEvent.CreatedAt, &messageEvent.UpdatedAt)

		if err != nil {
			return messageEvent, err
		}
	}
	err = rows.Err()
	if err != nil {
		return messageEvent, err
	}
	return messageEvent, nil
}

//Insert into DB a record
func insertMessageEvent(status string, description string, correlation_id int) (int64, int64, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	connection, err := sql.Open("mysql", connectionString)
	if err != nil {
		return 0, 0, err
	}
	statement := fmt.Sprintf("INSERT INTO `event_status` (`status`, `correlation_id`, `description`, `error`, `createdAt`, `updatedAt`) VALUES(?,?,?,?,?,?)")
	stmt, err := connection.Prepare(statement)
	if err != nil {
		return 0, 0, err
	}
	date := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(status, correlation_id, description, " ", date, date)
	if err != nil {
		return 0, 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	defer connection.Close()
	//ans := fmt.Sprintf("%v", lastId)
	return lastId, rowCnt, nil
}

//Update DB record
func updateMessageEvent(id int, status string, reason string) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	connection, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	statement := fmt.Sprintf("UPDATE `event_status` SET `status` = ?, `error` = ? WHERE (`id` = ?);")
	stmt, err := connection.Prepare(statement)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(status, reason, id)
	if err != nil {
		return err
	}
	defer connection.Close()
	return nil
}

//Update DB record
func updateMessageEventCorrelationId(id int64, correlationId int64) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	connection, err := sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	statement := fmt.Sprintf("UPDATE `event_status` SET `correlation_id` = ? WHERE (`id` = ?);")
	stmt, err := connection.Prepare(statement)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(correlationId, id)
	if err != nil {
		return err
	}
	defer connection.Close()
	return nil
}

//Main function
func main() {
	lambda.Start(handler)
}
