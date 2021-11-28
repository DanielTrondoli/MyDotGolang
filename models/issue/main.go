package issue

import (
	"github.com/DanielTrondoli/MyDotGolang/models/worklogs"
)

// Issue ...
type Issue struct {
	UUID     string              `json:"uuid" dynamo:"uuid,hash"`
	URL      string              `json:"url" dynamo:"url"`
	Title    string              `json:"title" dynamo:"title"`
	WorkLogs []worklogs.WorkLogs `json:"workLogs" dynamo:"workLogs"`
	Hide     bool                `json:"hide" dynamo:"hide"`
}
