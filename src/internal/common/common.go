package common

import (
	"math/rand"
	"time"
)

func SelectRandomString(s []string) string {
	rand.NewSource(time.Now().UnixNano())
	selectedString := s[rand.Intn(len(s))]
	return selectedString
}
