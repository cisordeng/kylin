package main

import (
	"github.com/cisordeng/beego/xenon"

	_ "kylin/model"
	_ "kylin/rest"
)

func main() {
	xenon.Run()
}
