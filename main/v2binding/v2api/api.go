package main

import (
	"time"

	"github.com/robovpn/v2fly_core/main/v2binding"
)

func main() {
	v2binding.StartAPIInstance(".")
	for {
		time.Sleep(time.Minute)
	}
}
