package process

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

var processItems struct {
	sync.Mutex
	data []float64
}

type Process struct {
	Index     int
	Total     uint64
	InProcess uint64
}

func NewProcess(index int, total uint64) *Process {
	i := Process{Index: index, Total: total * 8}
	processItems.Lock()
	processItems.data = append(processItems.data, 0)

	processItems.Unlock()

	return &i
}

func (wc *Process) Write(p []byte) (int, error) {
	n := len(p)
	wc.InProcess += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc *Process) PrintProgress() {
	pChange := wc.percentage(wc.InProcess, wc.Total)
	processItems.Lock()
	if len(processItems.data) > wc.Index && pChange != processItems.data[wc.Index] {
		processItems.data[wc.Index] = wc.percentage(wc.InProcess, wc.Total)
		fmt.Printf("%v\n", wc.parseInfo())
	}
	processItems.Unlock()
}

func (wc *Process) percentage(old, new uint64) float64 {
	return math.Ceil((float64(old) / float64(new)) * 100)
}

func (wc *Process) parseInfo() (res string) {
	var sSlice []string
	for _, info := range processItems.data {
		if info > 0 {
			sSlice = append(sSlice, fmt.Sprintf("%v%v", info, "%"))
		}
	}
	return strings.Join(sSlice, " ")
}
