package main

import (
	"github.com/dr4ke616/gospeedtest/nw_speedtest"
	"log"
)

func main() {
	st := nw_speedtest.Speedtest{
		FileLocation: "http://download.thinkbroadband.com/200MB.zip",
		Verbos:       true,
	}
	rate, _ := st.Start()
	log.Println(rate)
}
