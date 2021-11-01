package main

import (
	"github.com/altairtt/pkg/service"
	"github.com/altairtt/pkg/utils"
	"log"
)

func main(){
	tools := utils.NewTools()

	srv := service.NewService(tools)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}