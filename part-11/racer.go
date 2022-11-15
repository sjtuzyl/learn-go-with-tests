package racer

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func LogWriter(w io.Writer, val string) {
	fmt.Fprintf(w, val)
}

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// 串行
	// aDelta := measureResponseTime(a)
	// LogWriter(os.Stdout, a+": "+aDelta.String()+"\n")

	// bDelta := measureResponseTime(b)
	// LogWriter(os.Stdout, b+": "+bDelta.String()+"\n")

	// if aDelta < bDelta {
	// 	return a
	// }
	// return b

	// 并行
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
