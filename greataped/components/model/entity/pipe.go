package entity

import . "time"

type pipeEntity struct {
	Pipe           int    `json:"pipe"`
	Source         string `json:"source"`
	Editor         int64  `json:"editor"`
	QueueTimestamp Time   `json:"queued_at"`
}

func (entity *pipeEntity) GetPipe() int {
	return entity.Pipe
}

func (entity *pipeEntity) GetSource() string {
	return entity.Source
}

func (entity *pipeEntity) GetEditor() int64 {
	return entity.Editor
}

func (entity *pipeEntity) GetQueueTimestamp() Time {
	return entity.QueueTimestamp
}
