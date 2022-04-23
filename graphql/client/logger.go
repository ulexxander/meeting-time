package client

import "log"

type Logger interface {
	Query(query string, variables Variables)
	QueryResponse(res []byte)
}

type LoggerStd struct {
	Logger *log.Logger
}

func (l *LoggerStd) Query(query string, variables Variables) {
	l.Logger.Printf("GraphQL query: %s, variables: %v", query, variables)
}

func (l *LoggerStd) QueryResponse(res []byte) {
	l.Logger.Printf("GraphQL query response: %s", res)
}
