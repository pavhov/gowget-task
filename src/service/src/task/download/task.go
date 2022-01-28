package download

import (
	"fmt"
	"gowget/src/task/process"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Interface interface {
	Start()
}

type Task struct {
}

func New() Interface {
	var i = Task{}
	return &i
}

func (t *Task) Start() {
	urls := os.Args[1:]

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for index, url := range urls {
		urlSlice := strings.Split(url, "/")
		file := urlSlice[len(urlSlice)-1]

		fmt.Print(fmt.Sprintf("%s ", file))
		go t.download(index, url, file, &wg)
	}
	fmt.Println("")
	wg.Wait()
}

func (t *Task) download(index int, url, file string, wg *sync.WaitGroup) {
	defer wg.Done()
	wd, _ := os.Getwd()

	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(err.Error())
	}
	out, err := os.Create(fmt.Sprintf("%s/%s", wd, file))
	if err != nil {
		panic(err.Error())
	}
	defer out.Close()
	total, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	counter := process.NewProcess(index, uint64(total/8))
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		panic(err.Error())
	}
}
