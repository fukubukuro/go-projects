package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(bytes []byte) (int, error) {
	n, e := r.r.Read(bytes)
	for idx := 0; idx < len(bytes); idx++ {
		b := bytes[idx]
		if b < 'A' || b > 'z' {
			continue
		}
		if b < 'N' || b < 'n' {
			bytes[idx] = b + 13
		} else {
			bytes[idx] = b - 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, &r)
	if err != nil {
		return
	}
}
