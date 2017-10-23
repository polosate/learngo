package main

import (
	"flag"
	"fmt"
	"os"

	"learngo/go_count/counter"
)

func main() {
	sourceType := flag.String("type", "", "source_type [url, file]")
	flag.Parse()

	counter := counter.NewCounter(*sourceType)
	counter.Execute(os.Stdin)

	fmt.Print(counter.GetResult())
}
