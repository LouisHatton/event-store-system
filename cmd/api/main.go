package main

import (
	"fmt"

	"github.com/LouisHatton/user-audit-saas/internal/store/dynamodb"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	db := dynamodb.New(*logger, "wave-api")

	var doc map[string]interface{}
	err := db.Get("test-item", &doc)
	if err != nil {
		logger.Error("failed to get document", zap.Error(err))
	}

	fmt.Println(doc)
}
