package coyote

import (
	"testing"
)

func TestIDToRedisString(t *testing.T) {
	id := "rick"
	expect := USER_EVENT_PREFIX + ":" + id
	redisString := idToRedisString(id)
	if redisString != expect {
		t.Error("Expected " + expect + " got " + redisString)
	}
}

func TestRedisStringToID(t *testing.T) {
	redisString := USER_EVENT_PREFIX + ":rick"
	expect := "rick"
	id := redisStringToID(redisString)
	if id != expect {
		t.Error("Expected " + expect + " got " + id)
	}
}
