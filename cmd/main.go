package main

import (
	"Project/api"
	"Project/dao"
)

func main() {
	dao.InitMySql()
	api.InitEngine()
}
