package models

import (
	"os"
	"sync"
	"time"
)

type Sigterm struct {
	KillChannel chan os.Signal
	WaitGroup   *sync.WaitGroup
	TimeOut     time.Duration
}

func NewSigtermStruct(kc chan os.Signal, wg *sync.WaitGroup) *Sigterm {
	return &Sigterm{kc, wg, 30 * time.Second}
}

//TODO: methods for accessing the wg (without )
