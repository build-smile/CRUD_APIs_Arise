package main

import (
	"github.com/build-smile/CRUD_APIs_Arise/httpserve"
	"github.com/build-smile/CRUD_APIs_Arise/infrastructure"
)

func main() {
	infrastructure.InitConfig()
	infrastructure.ConnectDb()
	httpserve.Run()
}
