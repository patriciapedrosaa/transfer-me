package memory

import "time"

type Transfer struct {
	id                   string
	accountOriginID      string
	accountDestinationID string
	amount               int
	createdAt            time.Time
}
