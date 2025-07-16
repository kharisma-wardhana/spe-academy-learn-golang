package consumer

import (
	"context"
	"fmt"

	"github.com/kharisma-wardhana/spe-academy-learn-golang/try-consumer-rabbitmq/todo-api/entity"
)

type LogQueue struct {
	ctx context.Context
}

type LogConsumer interface {
	ProcessSyncLog(payload map[string]interface{}) error
}

func NewLogConsumer(
	ctx context.Context,
) LogConsumer {
	return &LogQueue{ctx}
}

func (l *LogQueue) ProcessSyncLog(payload map[string]interface{}) error {
	var params entity.Log
	params.LoadFromMap(payload)

	fmt.Println("SYNC SUCCESS!")
	fmt.Println(params)

	return nil
}
