package worklogs

import "time"

// TimeStamp is a type of logwork
type WorkLogs struct {
	StartTime time.Time `json:"startTime" dynamo:"startTime"`
	EndTime   time.Time `json:"endTime" dynamo:"endTime "`
}

// IsStartEqualsEnd ...
func (t WorkLogs) IsStartEqualsEnd() bool {
	return t.StartTime.Equal(t.EndTime)
}

// Duration ...
func (t WorkLogs) WorkLogDuration() time.Duration {
	if t.EndTime != (time.Time{}) {
		return t.EndTime.Sub(t.StartTime)
	}
	return -1
}
