package domain

import "time"

type Participant struct {
	UserId   UserId
	JoinedAt time.Time
}
