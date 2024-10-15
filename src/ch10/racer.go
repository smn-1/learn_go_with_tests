package main

import (
	"fmt"
	"net/http"
	"time"
)

var fiveSecondTimeout = 5 * time.Second

func Racer(a, b string) (winner string, err error) {
	return configurableRacer(a, b, fiveSecondTimeout)
}

func configurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
