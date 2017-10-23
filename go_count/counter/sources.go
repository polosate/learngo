package counter

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	file = "file"
	url  = "url"
)

// Source interface
type Source interface {
	Read() ([]byte, error)
	GetPath() string
	SetPath(string)
}

// URL describes url-source
type URL struct {
	Source
	path string
}

// File describe file-source
type File struct {
	Source
	path string
}

// NewSource creates new source
func NewSource(sType, path string) Source {
	var s Source
	switch sType {
	case strings.ToLower(url):
		s = &URL{}
	case strings.ToLower(file):
		s = &File{}
	default:
		log.Fatalf("Not supported source type: %q", sType)
		return nil
	}
	s.SetPath(path)
	return s
}

// Read does GET request and returns loaded data
func (this *URL) Read() ([]byte, error) {
	c := http.Client{}
	resp, err := c.Get(this.path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	return body, nil
}

// Read reads from file
func (this *File) Read() ([]byte, error) {
	f, err := os.Open(this.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bs []byte
	if bs, err = ioutil.ReadAll(f); err != nil {
		return nil, err
	}
	return bs, nil
}

func (this *URL) SetPath(path string) {
	this.path = path
}

func (this *File) SetPath(path string) {
	this.path = path
}

func (this *URL) GetPath() string {
	return this.path
}

func (this *File) GetPath() string {
	return this.path
}
