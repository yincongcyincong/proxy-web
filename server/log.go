package server

import (
	"github.com/snail007/webtail"
	"log"
)

func InitShowLog(){
	address := ":8822"
	basedir := "./log"
	listener, err := webtail.Serve(address, basedir)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tail server on %s",(*listener).Addr())
}