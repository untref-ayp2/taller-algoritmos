package main

import (
	"bufio"
	"os"
	"time"
)

func main() {
	lector := bufio.NewReader(os.Stdin)

	n := LeerCantidadReinas(lector, os.Stdout)
	AnimarReinas(n, lector, os.Stdout, 40*time.Millisecond)
}
