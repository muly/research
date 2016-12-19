package main

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

const (
	redisSrv = "127.0.0.1:6379"
	index    = "ProcessedRecords"
	data     = "NEWS_XML"
)

func main() {
	conn, err := redis.Dial("tcp", redisSrv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println("Hello Redis")

	cleanup(conn)
	return

	var xmlfilename string = "00a5e539322693f39d9923e196727419"
	var xmldata string = "......................... 00a5e539322693f39d9923e196727419 ....................."

	resp, err := conn.Do("SISMEMBER", index, xmlfilename)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else if reflect.DeepEqual(resp, int64(1)) {
		fmt.Println("Record already exists, so skipped ", xmlfilename)
		return
	}
	fmt.Println("Response: ", resp)

	fmt.Printf("ret type: %T, ret: %v\n", resp, resp)

	resp, err = conn.Do("LPUSH", data, xmldata)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err = conn.Do("SADD", index, xmlfilename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func cleanup(conn redis.Conn) {
	//for {
	resp, err := conn.Do("SPOP", index, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Response: ", resp == nil)

	//}
}
