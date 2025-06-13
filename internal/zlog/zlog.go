package zlog

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

type Logger struct {
	nats      *nats.Conn
	ZapLogger *zap.Logger
}

func newLogger(nats *nats.Conn, logger *zap.Logger) *Logger {
	return &Logger{
		nats:      nats,
		ZapLogger: logger,
	}
}

func Provide() fx.Option {
	return fx.Provide(newLogger)
}

type LogData struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"projectID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Removed     bool      `json:"removed"`
	Timestamp   time.Time `json:"timestamp"`
}

func (log *Logger) InfoNats(id, projectID int, name, description string, priority int, removed bool) {
	//log.ZapLogger.Info("Sending to Nats",
	//	zap.Int("id", id),
	//	zap.Int("projectID", projectID),
	//	zap.String("name", name),
	//	zap.String("description", description),
	//	zap.Int("priority", priority),
	//	zap.Bool("removed", removed))

	logData := LogData{
		ID:          id,
		ProjectID:   projectID,
		Name:        name,
		Description: description,
		Priority:    priority,
		Removed:     removed,
		Timestamp:   time.Now(),
	}

	data, err := json.Marshal(logData)
	if err != nil {
		log.ZapLogger.Error("Failed to marshal log data", zap.Error(err))
		return
	}

	err = log.nats.Publish("logs", data)
	if err != nil {
		log.ZapLogger.Error("Failed to publish to NATS", zap.Error(err))
	}
}
