package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	mySession := session.Must(session.NewSession(aws.NewConfig().WithCredentials(
		credentials.NewStaticCredentials(
			"fa9fd7d56e75b89c19505c39935fb8d5",
			"173cf1df0f9fa7039206a461c43e09de578337e67e92e578049ae7a14f6fca34",
			"",
		),
	)))
	cli := s3.New(
		mySession,
		aws.NewConfig().
			WithRegion("auto").
			WithEndpoint("https://ab71136bc4c645df4dfbb7ded7e53cfd.r2.cloudflarestorage.com/chenyq"),
	)

	spew.Dump(cli.PutBucketCors(&s3.PutBucketCorsInput{
		Bucket: toPtr("chenyq"),
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: []*s3.CORSRule{{
				AllowedOrigins: []*string{toPtr("https://chenyq.com")},
				AllowedMethods: []*string{toPtr("GET"), toPtr("HEAD")},
			}},
		},
	}))

	spew.Dump(cli.GetBucketCors(&s3.GetBucketCorsInput{
		Bucket: toPtr("chenyq"),
	}))
}

func toPtr(s string) *string {
	return &s
}
