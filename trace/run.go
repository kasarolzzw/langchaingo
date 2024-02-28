package trace

import "time"

type Run struct {
	ID             string
	Name           string
	TraceID        string
	ParentRunID    string
	RunType        RunType
	StartTime      time.Time
	EndTime        time.Time
	RunStatus      RunStatus
	ErrorCode      int64
	ErrorMsg       string
	Inputs         any
	Outputs        any
	InputTokens    int64
	OutputTokens   int64
	FirstTokenTime time.Time
}

type RunType string

const (
	RunTypeChain     RunType = "chain"
	RunTypeLLM       RunType = "llm"
	RunTypeRetriever RunType = "retriever"
	RunTypeTool      RunType = "tool"
)

type RunStatus int

const (
	RunStatusSuccess RunStatus = iota
	RunStatusError
)
