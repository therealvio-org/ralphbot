package extract

import (
	"math/rand"
	"time"
)

func RandomString(s []string) string {
	rand.NewSource(time.Now().UnixNano())
	selectedString := s[rand.Intn(len(s))]
	return selectedString
}
