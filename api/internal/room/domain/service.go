package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type RoomService struct{}

func (rs *RoomService) Create(createdByUserId string, isPrivateRoom bool) (*Room, error) {
	passCode := ""

	if isPrivateRoom {
		passCode = generateRoomPassCode()
	}

	return &Room{
		Id:        RoomId(generateRoomCode()),
		IsPrivate: isPrivateRoom,
		PassCode:  passCode,
		OwnerId:   UserId(createdByUserId),
		CreatedAt: time.Now(),
	}, nil
}

// generateRoomCode generate random room code
func generateRoomCode() string {
	const chars = "abcdefghijkmnpqrstuvwxyz23456789"
	rand.NewSource(time.Now().UnixNano())

	// Kodun parçalarını oluştur
	part1 := make([]byte, 3)
	for i := range part1 {
		part1[i] = chars[rand.Intn(len(chars))]
	}

	part2 := make([]byte, 4)
	for i := range part2 {
		part2[i] = chars[rand.Intn(len(chars))]
	}

	part3 := make([]byte, 3)
	for i := range part3 {
		part3[i] = chars[rand.Intn(len(chars))]
	}

	return fmt.Sprintf("%s-%s-%s", part1, part2, part3)
}

// generateRoomPassCode generate 6 digits passcode
func generateRoomPassCode() string {
	rand.NewSource(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	return fmt.Sprintf("%d", code)

}
