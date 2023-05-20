package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/LouisHatton/insight-wave/internal/events"
	"github.com/LouisHatton/insight-wave/internal/eventstore"
	"github.com/LouisHatton/insight-wave/internal/eventstore/tinybird"
	lambdaevents "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.uber.org/zap"
)

func handler(ctx context.Context, sqsEvent lambdaevents.SQSEvent) error {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	eventStore := tinybird.New(*logger)

	for _, message := range sqsEvent.Records {
		handleMessage(logger, message, eventStore)
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}

	return nil
}

func handleMessage(log *zap.Logger, message lambdaevents.SQSMessage, store eventstore.Storer) error {

	i, err := strconv.ParseInt(message.Attributes["ApproximateFirstReceiveTimestamp"], 10, 64)
	var timestamp time.Time

	if err != nil {
		timestamp = time.Now()
	}
	timestamp = time.Unix(i/1000, 0)

	var bodymap map[string]string
	json.Unmarshal([]byte(message.Body), &bodymap)
	body, _ := json.Marshal(bodymap)

	event := events.Event{
		Timestamp: &timestamp,
		ProjectId: "testing",
		Version:   events.V1,
		Payload:   string(body),
	}

	err = store.AddEvent(&event)
	if err != nil {
		log.Error("Failed to add event to store: ", zap.Error(err))
		return err
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
