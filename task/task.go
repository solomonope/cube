package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
	"github.com/solomonope/cube/job"
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         job.State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
	Timestamp     time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     job.State
	Timestamp time.Time
	Task      Task
}
