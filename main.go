package main

import (
	"log"

	"github.com/robfig/cron"
)

type stat struct {
	mac int
}

func (s *stat) Run() {
	log.Println(s.mac)
}

func main() {
	spec := "0-59/5 * * * * *"
	st := &stat{
		mac: 1000,
	}
	c := cron.New()
	c.AddJob(spec, st)
	c.Start()

	select {}
}
