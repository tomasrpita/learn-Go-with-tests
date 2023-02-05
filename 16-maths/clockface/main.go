package main

import (
	"os"
	"time"

	clockface "learn-Go-with-tests/16-maths/svg"
)

func main() {
	t := time.Now()
	clockface.Writer(os.Stdout, t)

}
