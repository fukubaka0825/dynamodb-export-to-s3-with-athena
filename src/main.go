package main

import (
	"context"
	"main.go/src/module"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event module.ExportEvent) error {
	return module.NewExportHandler(&event).Run()
}
