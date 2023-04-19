package G_scheduler

import (
	G_mq "Gungnir/internal/master/G-mq"
	"log"
)

var Scheduler *GScheduler

type GScheduler struct{}

func New() {
	Scheduler = &GScheduler{}
}

func GetSchedulerOr() *GScheduler {
	if Scheduler == nil {
		New()
	}
	return Scheduler
}

func (g *GScheduler) Run() {
	go func() {
		for {
			jobDesc := <-G_mq.GetMqOr().JobStartSchedulerChan
			log.Println(jobDesc)
		}
	}()
}
