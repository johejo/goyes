package main

import (
	"bufio"
	"os"
)

func main() {
	var s string
	if len(os.Args) == 1 {
		s = "y"
	} else {
		s = os.Args[1]
	}
	s += "\n"

	const bufLen = 1024 * 32
	buf := make([]byte, 0, bufLen)
	for len(buf) <= bufLen {
		buf = append(buf, s...)
	}

	w := bufio.NewWriter(os.Stdout)
	for {
		w.Write(buf)
	}
}
