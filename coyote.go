package coyote

import (
	"github.com/mediocregopher/radix.v2/pool"
)

var p pool.Pool

/*
func Connect() error {
	pool, err := pool.New("tcp", "localhost:6379", 10)
	return err
}
*/

func Connect(redisHost string) error {
	p, err := pool.New("tcp", redisHost, 10)
	_ = p
	return err
}

func IPEvent(ipaddress string) {

}

func UserEvent(userID string) {

}

func checkUser(userID string) error {
	conn, err := p.Get()
	if err != nil {
		return err
	}
	defer p.Put(conn)
	return nil
}
