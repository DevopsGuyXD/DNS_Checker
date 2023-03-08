package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

func CheckForNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckIfValidJson(jsondata []byte) bool{
	checkValidJson := json.Valid(jsondata)

	return checkValidJson
}

func InitEnvFile(){
	err := godotenv.Load(".env"); CheckForNil(err)
}

func InitAws(){

	_, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
		},
	})
	CheckForNil(err)
}