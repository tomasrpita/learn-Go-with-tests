package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdiwnStart = 3

func Countdown(out io.Writer) {
	for i := countdiwnStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)

	}
	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
