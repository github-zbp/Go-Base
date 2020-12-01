package main

import (
	"fmt"
	"os"
)

func main() {
	if f, err := os.Open(fname); err != nil { // compile error: unused: f
		return err
	}
	f.ReadByte()
	f.Close()
}

