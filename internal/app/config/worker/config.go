package config

import (
	"os"
	"strconv"
)

func New() *Worker {
	config := &Worker{}

	config.init()

	return config
}

func (w *Worker) init() {
	w.initEnv()
	w.initDefaultValue()
}

func (w *Worker) initEnv() {
	w.WorkerCounter, _ = strconv.Atoi(os.Getenv("WORKER_COUNTER"))
	w.ChanelLength, _ = strconv.Atoi(os.Getenv("WORKER_CHANEL_LENGTH"))
}

func (w *Worker) initDefaultValue() {
	if w.WorkerCounter == 0 {
		w.WorkerCounter = workerCounter
	}

	if w.ChanelLength == 0 {
		w.ChanelLength = chanelLength
	}
}
