package domain

import (
	"time"
)

type UserId string
type RoomId string

type Room struct {
	Id           RoomId `bson:"_id"`
	OwnerId      UserId
	CreatedAt    time.Time
	PassCode     string
	IsPrivate    bool
	Participants []Participant
}

func (r *Room) AddParticipant(userId UserId) error {
	if r.HasParticipant(userId) {
		return nil
	}

	r.Participants = append(r.Participants, Participant{
		UserId:   userId,
		JoinedAt: time.Now(),
	})
	return nil
}

func (r *Room) HasParticipant(userId UserId) bool {
	if r.OwnerId == userId {
		return true
	}

	if len(r.Participants) == 0 {
		return false
	}

	for _, p := range r.Participants {
		if p.UserId == userId {
			return true
		}
	}
	return false
}
