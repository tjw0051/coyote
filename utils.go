package coyote

import (
	"strings"
)

const (
	USER_EVENT_PREFIX = "USER_EVENT"
	IP_EVENT_PREFIX   = "IP_EVENT"
)

func redisStringToID(redisString string) string {
	userID := strings.Replace(redisString, USER_EVENT_PREFIX+":", "", 1)
	return userID
}

func idToRedisString(id string) string {
	return USER_EVENT_PREFIX + ":" + id
}
