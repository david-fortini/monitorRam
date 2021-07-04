package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func getMemoryMb(proc *process.Process) uint64 {
	memInfo, _ := proc.MemoryInfo()
	rss := memInfo.RSS
	rssMb := rss / (1024 * 1024)
	return rssMb
}

func main() {
	if len(os.Args) < 2 {
		errors.New("A filepath argument is required")
	}
	MAX_MB := flag.Uint64("max_mb", 3000, "Maximum mb allowed")
	flag.Parse()

	for true {

		processes, _ := process.Processes()
		for x := range processes {
			proc := processes[x]
			procName, _ := proc.Name()
			isNotPython := !strings.Contains(procName, "python")
			if isNotPython {
				continue
			}
			rssMb := getMemoryMb(proc)
			if rssMb > *MAX_MB {
				proc.Kill()
				fmt.Printf("killing %s", procName)
			}

		}
		time.Sleep(2 * time.Second)
		fmt.Printf("\n new loop...")
	}
}
