package racer

import "net/http"

// func Racer(a, b string) (winner string) {
// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration > bDuration {
// 		return b
// 	}
// 	return a
// }

// func measureResponseTime(url string) time.Duration {
// 	startA := time.Now()
// 	http.Get(url)
// 	return time.Since(startA)
// }

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
