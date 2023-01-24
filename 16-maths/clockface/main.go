package main

import (
	"os"
	"time"

	clockface "learn-Go-with-tests/16-maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)

}
