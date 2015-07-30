package main

import (
	"github.com/dr4ke616/go_speedtest/go_speedtest"
	"log"
)

func main() {
	st := go_speedtest.Speedtest{
		FileLocation: "http://download.thinkbroadband.com/200MB.zip",
		Verbos:       true,
	}
	log.Println(st.Start())
}
