package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"io/ioutil"
)

type InputReader struct {
}


func (InputReader) Read(bucket, source string) ([]byte,error){
	return readBuckets(bucket,source)
}

func ListBuckets(bucket, source string) {

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	result, err := svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		log.Println("Failed to list buckets", err)
		return
	}

	log.Println("Buckets:")

	for _, bucket := range result.Buckets {
		log.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
	}
}

func readBuckets(bucket, key string) ([]byte,error) {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	// Create S3 service client
	svc := s3.New(sess)

	output,err:=svc.GetObject(&s3.GetObjectInput{
		Bucket:&bucket,
		Key:&key,

	})
	if err !=nil{
		return nil,err
	}

	object,err:=ioutil.ReadAll(output.Body)
	if err !=nil{
		return nil,err
	}
	output.Body.Close()

	return object,err
}
