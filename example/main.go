package main

import (
	"fmt"

	"github.com/rxwen/xiaomi-go-sdk-push"
)

const (
	appSecret = "xxxxxxxxxxxx=="
)

func main() {
	c := push.NewClient(appSecret)
	r := push.Request{
		Title:       "hi",
		Description: "sth happened",
		Payload:     "hi there, checkout what happended",
		TargetType:  "all",
		//TargetType:  "regid",
		//TargetValue: "xxxxxxxxxxxx",
	}
	body, _ := c.SendRequest(&r)
	fmt.Println(body)
}
