package scheduler

import (
	"goweb/crawler/model"
)

type SimpleScheduler struct {
	workerChan chan model.Request
}

func (s *SimpleScheduler) ConfigWorkerChan(c chan model.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request model.Request) {
	go func() {s.workerChan <- request}()
}