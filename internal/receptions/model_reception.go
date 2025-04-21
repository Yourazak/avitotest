package receptions

import (
	"github.com/google/uuid"
	"time"
)

type Reception struct {
	ID       uuid.UUID `json:"id"`
	DataTime time.Time `json:"date_time"`
	PvzID    uuid.UUID `json:"pvz_id"`
	Status   string    `json:"status"`
}

//required: [dataTime,pvzId,status]

func NewReception(dataTime time.Time, pvzID uuid.UUID, status string) *Reception {
	return &Reception{
		DataTime: dataTime,
		PvzID:    pvzID,
		Status:   status,
	}
}
