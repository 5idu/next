package redis

import (
	"fmt"
	"testing"
)

func TestDial(t *testing.T) {
	var conf RedisConf
	conf.Addr = "localhost:6379"
	conf.Password = ""
	conf.DB = 0
	conn, err := Dial(&conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(conn.Get("name").String())
}
