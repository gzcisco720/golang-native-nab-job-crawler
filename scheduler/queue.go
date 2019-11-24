package scheduler

import (
	"goweb/model"
)

type QueuedScheduler struct {
	requestChan chan model.Request
	workerChan chan chan model.Request
}

func (s *QueuedScheduler) ConfigWorkerChan(c chan model.Request) {

}

func (s *QueuedScheduler) WorkerReady(c chan model.Request) {
	s.workerChan <- c
}

func (s *QueuedScheduler) Submit(request model.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan model.Request)
	s.requestChan = make(chan model.Request)
	go func() {
		var requestQueue []model.Request
		var workerQueue []chan model.Request
		for {
			var activeRequest model.Request
			var activeWorker chan model.Request
			if len(requestQueue)>0 && len(workerQueue)>0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
				case r := <- s.requestChan:
					requestQueue = append(requestQueue, r)
				case w := <- s.workerChan:
					workerQueue = append(workerQueue, w)
				case activeWorker <- activeRequest:
					workerQueue = workerQueue[1:]
					requestQueue = requestQueue[1:]
			}
		}
	}()
}