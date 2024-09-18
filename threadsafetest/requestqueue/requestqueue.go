package requestqueue

import (
	"net/http"
	"sync"
)

type RequestQueue struct{
	Chan chan *http.Request
	Mutex *sync.Mutex
}