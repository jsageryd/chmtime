package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: chmtime <rfc3339nano> <path>")
	}

	t, err := time.Parse(time.RFC3339Nano, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	p := os.Args[2]

	stat, err := os.Lstat(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s old mtime: %s\n", stat.Name(), stat.ModTime().Format(time.RFC3339Nano))

	if err := os.Chtimes(p, t.UTC(), t.UTC()); err != nil {
		log.Fatal(err)
	}

	stat, err = os.Lstat(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s new mtime: %s\n", strings.Repeat(" ", len(stat.Name())), stat.ModTime().Format(time.RFC3339Nano))
}
