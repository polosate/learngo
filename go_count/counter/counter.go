package counter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	search          = "Go"
	maxWorkersCount = 5
	maxSources      = 128
)

// Counter describes workers pool object
type Counter struct {
	mu         sync.Mutex
	wg         sync.WaitGroup
	sourceType string
	sources    chan Source
	total      int
	result     chan string
}

// NewCounter constructor for Counter object
func NewCounter(sourceType string) *Counter {
	counter := &Counter{
		sourceType: sourceType,
		sources:    make(chan Source, maxSources),
	}
	return counter
}

// exec starts workers
func (this *Counter) exec() {
	sourcesCount := len(this.sources)
	this.result = make(chan string, sourcesCount)
	m := maxWorkersCount

	if sourcesCount < maxWorkersCount {
		m = sourcesCount
	}
	for i := 0; i < m; i++ {
		this.wg.Add(1)
		go this.worker(i)
	}
}

func (this *Counter) Execute(stdin *os.File) {
	this.stdinReader(stdin)

	this.exec()

	close(this.sources)
	this.wg.Wait()
	close(this.result)
}

// GetResult returns count of the specified word in each source and total count
func (this *Counter) GetResult() string {
	var result string
	for res := range this.result {
		result += fmt.Sprintf("%s\n", res)
	}
	result += fmt.Sprintf("Total: %d\n", this.total)
	return result
}

func (this *Counter) stdinReader(stdin *os.File) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Println("stdin is empty")
		return
	}

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		source := NewSource(this.sourceType, scanner.Text())
		if source == nil {
			return
		} else {
			this.sources <- source
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}
}

func (this *Counter) worker(i int) {
	defer this.wg.Done()
	for s := range this.sources {
		b, err := s.Read()
		if err != nil {
			return
		}
		count := strings.Count(string(b), search)

		this.mu.Lock()
		this.total += count
		this.mu.Unlock()

		this.result <- fmt.Sprintf("Count for %s: %d", s.GetPath(), count)
	}
}
