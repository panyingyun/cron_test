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

type day struct {
	id int
}

func (d *day) Run() {
	log.Println(d.id)
}

//every 5 second run a job
var spec5s = "0-59/5 * * * * *"

//every day(00:00:00) run a job
var specday = "0 0 0 1-31 * *"

func main() {
	c := cron.New()

	st := &stat{
		mac: 1000,
	}

	c.AddJob(spec5s, st)

	day := &day{
		id: 9000,
	}

	c.AddJob(specday, day)

	c.Start()

	select {}
}
