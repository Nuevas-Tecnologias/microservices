package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	lamda_event "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type MessageData struct {
	Type             string `json:"Type""`
	MessageID        string `json:"MessageId"`
	TopicArn         string `json:"TopicArn"`
	Message          string `json:"Message"`
	Timestamp        string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
}
type Content struct {
	Content string `json:"content"`
}
type InputData struct {
	Company     string `json:"Company"`
	Environment string `json:"Environment"`
	InitialDate string `json:"InitialDate"`
	FinalDate   string `json:"FinalDate"`
}
type Payload struct {
	Body string `json:"body"`
}
type Body struct {
	Message string `json:"message"`
	Log     string `json:"log"`
}
type lambdaResponse struct {
	Headers    lamdaResponseHeaders `json:"headers"`
	Body       string               `json:"body"`
	StatusCode int                  `json:"statusCode"`
}
type lamdaResponseHeaders struct {
	ContentType string `json:"Content-Type"`
}
type lambdaResponseBody struct {
	NextResults bool   `json:"NextResults"`
	Result      string `json:"Result"`
	Error       bool   `json:"Error"`
	Type        string `json:"Type"`
	CompanyId   string `json:"CompanyId"`
}
type messageLog struct {
	Consecutive string `json:"consecutive"`
	Module      string `json:"module"`
	eventTypeId string `json:"event_type_id"`
	Detail      string `json:"detail"`
	Date        string `json:"date"`
}
type businessPartner struct {
	CardCode                string    `json:"CardCode"`
	CardName                string    `json:"CardName"`
	CardType                string    `json:"CardType"`
	GroupCode               int       `json:"GroupCode"`
	CreditLimit             float64   `json:"CreditLimit"`
	Phone1                  string    `json:"Phone1"`
	PayTermsGrpCode         int       `json:"PayTermsGrpCode"`
	FederalTaxID            string    `json:"FederalTaxID"`
	SalesPersonCode         int       `json:"SalesPersonCode"`
	EmailAddress            string    `json:"EmailAddress"`
	Valid                   string    `json:"Valid"`
	SubjectToWithholdingTax string    `json:"SubjectToWithholdingTax"`
	PriceListNum            int       `json:"PriceListNum"`
	U_HBT_RegTrib           string    `json:"U_HBT_RegTrib"`
	U_HBT_TipDoc            string    `json:"U_HBT_TipDoc"`
	U_HBT_MunMed            string    `json:"U_HBT_MunMed"`
	U_HBT_TipEnt            string    `json:"U_HBT_TipEnt"`
	U_HBT_Nacional          string    `json:"U_HBT_Nacional"`
	U_HBT_TipExt            string    `json:"U_HBT_TipExt"`
	DebitorAccount          string    `json:"DebitorAccount"`
	Currency                string    `json:"Currency"`
	DiscountPercent         float64   `json:"DiscountPercent"`
	VatLiable               string    `json:"VatLiable"`
	BPAddresses             []address `json:""BPAddresses"`
}
type address struct {
	AddressName string `json:"AddressName"`
	Street      string `json:"Street"`
	City        string `json:"City"`
	County      string `json:"County"`
	Country     string `json:"Country"`
	U_HBT_DirMM string `json:"U_HBT_DirMM"`
	AddressType string `json:"AddressType"`
	BPCode      string `json:"BPCode"`
	RowNum      int    `json:"RowNum"`
}

func main() {
	lamda_event.Start(handler)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	if len(sqsEvent.Records) == 0 {
		fmt.Println("No SQS message passed to function.")
		return nil
	}
	messageData := &MessageData{}
	content := &Content{}
	inputData := &InputData{}

	for _, msg := range sqsEvent.Records {
		//inserLog("abc123", "Trial detail")
		err := json.Unmarshal([]byte(msg.Body), messageData)
		if err != nil {
			fmt.Println("Error on unmarshaling message body. ", err.Error())
			return nil
		}
		err = json.Unmarshal([]byte(messageData.Message), content)
		if err != nil {
			fmt.Println("Error on unmarshaling message content. ", err.Error())
			return nil
		}
		err = json.Unmarshal([]byte(content.Content), inputData)
		if err != nil {
			fmt.Println("Error on unmarshaling message data. ", err.Error())
			return nil
		}
		//Inicialización de varibales
		businessPartnersBody := lambdaResponseBody{}
		businessPartnerArray := []businessPartner{}
		Peticion := 0
		//Loop de consultas
		for {
			erp_request := fmt.Sprintf(
				`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ws="http://ws.integracion.technologies.seidor.com/">
					<soapenv:Header/>
					<soapenv:Body>
						<ws:GetBusinessPartner>
							<Object>
							{
								"Company": "%s", 
								"FechaInicial": "%s", 
								"FechaFinal": "%s",
								"Peticion": %v
							}
							</Object>
						</ws:GetBusinessPartner>
					</soapenv:Body>
				</soapenv:Envelope>`,
				inputData.Environment,
				inputData.InitialDate,
				inputData.FinalDate,
				Peticion,
			)
			erp_body, err := json.Marshal(map[string]interface{}{
				"http_method":       string(erpHttpMethod),
				"xml_request":       string(erp_request),
				"integration_point": string("getBusinesspartners"),
			})
			if err != nil {
				fmt.Println("Error marshalling xml body")
				return nil
			}
			lambdaResponse, err := excecuteLambda(erpRegion, erp_body, erpLambda)
			if err != nil {
				fmt.Println("Error on lambda excecution. ", err.Error())
				return nil
			}
			lambdaResponseBody := lambdaResponseBody{}
			err = json.Unmarshal([]byte(lambdaResponse.Body), &lambdaResponseBody)
			if err != nil {
				fmt.Println("Error unmarshalling xml body response", err)
				return nil
			}
			if lambdaResponseBody.Error {
				fmt.Println("Failed to get suppliers", lambdaResponseBody.Result)
				return nil
			}
			bp := []businessPartner{}
			err = json.Unmarshal([]byte(lambdaResponseBody.Result), &bp)
			if err != nil {
				fmt.Println("Error unmarshalling business partners array", err)
				return nil
			}
			for _, v := range bp {
				businessPartnerArray = append(businessPartnerArray, v)
			}
			if lambdaResponseBody.NextResults {
				Peticion += 1
				fmt.Println("There are more results, request: ", Peticion)
			} else {
				completeBusinessPartners, err := json.Marshal(businessPartnerArray)
				if err != nil {
					fmt.Println("Error marshalling all the business partners")
					return nil
				}
				fmt.Println("These are all the results: ", string(completeBusinessPartners))
				businessPartnersBody = lambdaResponseBody
				businessPartnersBody.Result = string(completeBusinessPartners)
				businessPartnersBody.CompanyId = inputData.Company
				break
			}
		}
		body, err := json.Marshal(businessPartnersBody)
		if err != nil {
			fmt.Println("Error marshalling xml body. ", err)
			return nil
		}
		api_body, err := json.Marshal(map[string]interface{}{
			"body": string(body),
		})
		lambdaResponse, err := excecuteLambda(apiRegion, api_body, apiLambda)
		if err != nil {
			fmt.Println("Error on asynchronous lambda excecution. ", err.Error())
			return nil
		}
		fmt.Println(lambdaResponse)
	}
	return nil
}

/**
 * Excecutes lambda calling
 */
func excecuteLambda(region string, body []byte, functionName string) (lambdaResponse, error) {
	lambdaResponse := lambdaResponse{}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := lambda.New(sess, &aws.Config{Region: aws.String(region)})
	p := Payload{
		Body: string(body),
	}
	payload, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling request")
		return lambdaResponse, err
	}
	input := &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      payload,
	}
	result, err := client.Invoke(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeResourceNotFoundException:
				fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
			case lambda.ErrCodeInvalidRequestContentException:
				fmt.Println(lambda.ErrCodeInvalidRequestContentException, aerr.Error())
			case lambda.ErrCodeRequestTooLargeException:
				fmt.Println(lambda.ErrCodeRequestTooLargeException, aerr.Error())
			case lambda.ErrCodeUnsupportedMediaTypeException:
				fmt.Println(lambda.ErrCodeUnsupportedMediaTypeException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			case lambda.ErrCodeEC2UnexpectedException:
				fmt.Println(lambda.ErrCodeEC2UnexpectedException, aerr.Error())
			case lambda.ErrCodeSubnetIPAddressLimitReachedException:
				fmt.Println(lambda.ErrCodeSubnetIPAddressLimitReachedException, aerr.Error())
			case lambda.ErrCodeENILimitReachedException:
				fmt.Println(lambda.ErrCodeENILimitReachedException, aerr.Error())
			case lambda.ErrCodeEC2ThrottledException:
				fmt.Println(lambda.ErrCodeEC2ThrottledException, aerr.Error())
			case lambda.ErrCodeEC2AccessDeniedException:
				fmt.Println(lambda.ErrCodeEC2AccessDeniedException, aerr.Error())
			case lambda.ErrCodeInvalidSubnetIDException:
				fmt.Println(lambda.ErrCodeInvalidSubnetIDException, aerr.Error())
			case lambda.ErrCodeInvalidSecurityGroupIDException:
				fmt.Println(lambda.ErrCodeInvalidSecurityGroupIDException, aerr.Error())
			case lambda.ErrCodeInvalidZipFileException:
				fmt.Println(lambda.ErrCodeInvalidZipFileException, aerr.Error())
			case lambda.ErrCodeKMSDisabledException:
				fmt.Println(lambda.ErrCodeKMSDisabledException, aerr.Error())
			case lambda.ErrCodeKMSInvalidStateException:
				fmt.Println(lambda.ErrCodeKMSInvalidStateException, aerr.Error())
			case lambda.ErrCodeKMSAccessDeniedException:
				fmt.Println(lambda.ErrCodeKMSAccessDeniedException, aerr.Error())
			case lambda.ErrCodeKMSNotFoundException:
				fmt.Println(lambda.ErrCodeKMSNotFoundException, aerr.Error())
			case lambda.ErrCodeInvalidRuntimeException:
				fmt.Println(lambda.ErrCodeInvalidRuntimeException, aerr.Error())
			case lambda.ErrCodeResourceConflictException:
				fmt.Println(lambda.ErrCodeResourceConflictException, aerr.Error())
			case lambda.ErrCodeResourceNotReadyException:
				fmt.Println(lambda.ErrCodeResourceNotReadyException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return lambdaResponse, err
	}
	err = json.Unmarshal(result.Payload, &lambdaResponse)
	if err != nil {
		fmt.Println("Error unmarshalling payload response", err)
		return lambdaResponse, err
	}
	if lambdaResponse.StatusCode != 200 {
		fmt.Println("Error getting business partners, StatusCode: " + strconv.Itoa(lambdaResponse.StatusCode))
		return lambdaResponse, err
	}
	fmt.Println(lambdaResponse)
	return lambdaResponse, nil
}

/**
 * Excecutes lambda calling asynchronously
 */
func excecuteLambdaAsynchronously(region string, body []byte, functionName string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := lambda.New(sess, &aws.Config{Region: aws.String(region)})
	p := Payload{
		Body: string(body),
	}
	payload, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling request")
		return err
	}
	input := &lambda.InvokeInput{
		FunctionName:   aws.String(functionName),
		Payload:        payload,
		InvocationType: aws.String("Event"),
	}
	result, err := client.Invoke(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeResourceNotFoundException:
				fmt.Println(lambda.ErrCodeResourceNotFoundException, aerr.Error())
			case lambda.ErrCodeInvalidRequestContentException:
				fmt.Println(lambda.ErrCodeInvalidRequestContentException, aerr.Error())
			case lambda.ErrCodeRequestTooLargeException:
				fmt.Println(lambda.ErrCodeRequestTooLargeException, aerr.Error())
			case lambda.ErrCodeUnsupportedMediaTypeException:
				fmt.Println(lambda.ErrCodeUnsupportedMediaTypeException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			case lambda.ErrCodeEC2UnexpectedException:
				fmt.Println(lambda.ErrCodeEC2UnexpectedException, aerr.Error())
			case lambda.ErrCodeSubnetIPAddressLimitReachedException:
				fmt.Println(lambda.ErrCodeSubnetIPAddressLimitReachedException, aerr.Error())
			case lambda.ErrCodeENILimitReachedException:
				fmt.Println(lambda.ErrCodeENILimitReachedException, aerr.Error())
			case lambda.ErrCodeEC2ThrottledException:
				fmt.Println(lambda.ErrCodeEC2ThrottledException, aerr.Error())
			case lambda.ErrCodeEC2AccessDeniedException:
				fmt.Println(lambda.ErrCodeEC2AccessDeniedException, aerr.Error())
			case lambda.ErrCodeInvalidSubnetIDException:
				fmt.Println(lambda.ErrCodeInvalidSubnetIDException, aerr.Error())
			case lambda.ErrCodeInvalidSecurityGroupIDException:
				fmt.Println(lambda.ErrCodeInvalidSecurityGroupIDException, aerr.Error())
			case lambda.ErrCodeInvalidZipFileException:
				fmt.Println(lambda.ErrCodeInvalidZipFileException, aerr.Error())
			case lambda.ErrCodeKMSDisabledException:
				fmt.Println(lambda.ErrCodeKMSDisabledException, aerr.Error())
			case lambda.ErrCodeKMSInvalidStateException:
				fmt.Println(lambda.ErrCodeKMSInvalidStateException, aerr.Error())
			case lambda.ErrCodeKMSAccessDeniedException:
				fmt.Println(lambda.ErrCodeKMSAccessDeniedException, aerr.Error())
			case lambda.ErrCodeKMSNotFoundException:
				fmt.Println(lambda.ErrCodeKMSNotFoundException, aerr.Error())
			case lambda.ErrCodeInvalidRuntimeException:
				fmt.Println(lambda.ErrCodeInvalidRuntimeException, aerr.Error())
			case lambda.ErrCodeResourceConflictException:
				fmt.Println(lambda.ErrCodeResourceConflictException, aerr.Error())
			case lambda.ErrCodeResourceNotReadyException:
				fmt.Println(lambda.ErrCodeResourceNotReadyException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return err
	}
	if *result.StatusCode != int64(202) {
		fmt.Println("Error getting items, StatusCode: " + strconv.Itoa(int(*result.StatusCode)))
		return err
	}
	return nil
}

/**
 * Inserts log
 */
func inserLog(consecutive string, detail string) error {
	messageLog := messageLog{}
	messageLog.Consecutive = consecutive
	messageLog.Module = "get_supplier_integration_point"
	messageLog.eventTypeId = "1"
	messageLog.Detail = detail
	messageLog.Date = time.Now().Format("2006-01-02 15:04:05")
	body, err := json.Marshal(messageLog)
	if err != nil {
		fmt.Println("Error marshaling log message body. ", err.Error())
		return err
	}
	_, err = excecuteLambda(logRegion, body, logLambda)
	if err != nil {
		fmt.Println("Error executing log lambda. ", err.Error())
		return err
	}
	return nil
}
