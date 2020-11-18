package output

import (
	"sync"
)

type CoreInterface interface {
	Receiving(ch <-chan SuiteResult, wg *sync.WaitGroup)
}

func Init(name string) CoreInterface {
	return New(name)
}
