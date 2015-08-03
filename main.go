package main

import (
	"github.com/dr4ke616/gospeedtest/speedtest"
	"log"
)

func main() {
	st := speedtest.Speedtest{
		FileLocation: "http://download.thinkbroadband.com/200MB.zip",
		Verbos:       true,
	}
	log.Println(st.Start())
}
