package domain

import (
	"fmt"
	"time"
)

type UserId string
type RoomId string

type Room struct {
	Id           RoomId
	OwnerId      UserId
	CreatedAt    time.Time
	PassCode     string
	IsPrivate    bool
	Participants []Participant
}

func (r *Room) AddParticipant(userId UserId) error {
	if r.HasParticipant(userId) {
		return fmt.Errorf("user already in room")
	}

	r.Participants = append(r.Participants, Participant{
		UserId:   userId,
		JoinedAt: time.Now(),
	})
	return nil
}

func (r *Room) HasParticipant(userId UserId) bool {
	for _, p := range r.Participants {
		if p.UserId == userId {
			return true
		}
	}
	return false
}
