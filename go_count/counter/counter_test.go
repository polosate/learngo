package counter

import (
	"os"
	"reflect"
	"testing"
)

func TestStdinRead(t *testing.T) {
	f, err := os.Open("test_files/file_path.txt")
	if err != nil {
		t.FailNow()
	}
	defer f.Close()

	counter := NewCounter("file")
	counter.stdinReader(f)
	close(counter.sources)
	counter.wg.Wait()
	sources := []string{}
	expected := []string{"test_files/test01.txt", "test_files/test02.txt"}
	for s := range counter.sources {
		sources = append(sources, s.GetPath())
	}

	if !reflect.DeepEqual(sources, expected) {
		t.Errorf("Wrong read from stdin")
		t.FailNow()
	}
}

func TestReadFromURLs(t *testing.T) {
	f, err := os.Open("test_files/url_path.txt")
	if err != nil {
		t.FailNow()
	}
	defer f.Close()

	counter := NewCounter("url")
	counter.Execute(f)

	if counter.total != 27 {
		t.Errorf("Wrong read from URLs")
		t.FailNow()
	}
}

func TestReadFromFiles(t *testing.T) {
	f, err := os.Open("test_files/file_path.txt")
	if err != nil {
		t.FailNow()
	}
	defer f.Close()

	counter := NewCounter("file")
	counter.Execute(f)

	if counter.total != 4 {
		t.Errorf("Wrong result: %d. Expected: 4", counter.total)
		t.FailNow()
	}

	r := counter.GetResult()
	expected := "Count for test_files/test01.txt: 3\nCount for test_files/test02.txt: 1\nTotal: 4\n"
	if r != expected {
		t.Errorf("Wrong result: \n%s\nExpected: \n%s\n", r, expected)
		t.FailNow()
	}
}
