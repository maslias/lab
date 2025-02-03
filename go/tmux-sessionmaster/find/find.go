package find

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"tmux-sessionmaster/types"
)

// func FindCmd() ([]types.OutResult, error) {
func FindCmd(wg *sync.WaitGroup, ch chan<- []types.OutResult) {
	defer wg.Done()
	if _, err := exec.LookPath("find"); err != nil {
		// return nil, nil
		return
	}

	envHome := os.Getenv("HOME")
	opt := []string{
		envHome,
		envHome + "/.config",
		"-mindepth", "1",
		"-maxdepth", "1",
		"-type", "d",
	}
	var stdOut, stdErr bytes.Buffer

	cmdFind := exec.Command("find", opt...)

	cmdFind.Stdout = &stdOut
	cmdFind.Stderr = io.MultiWriter(os.Stderr, &stdErr)

	if err := cmdFind.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmdFind.Wait(); err != nil {
		log.Fatal(err)
	}

	list := strings.Split(strings.TrimSpace(stdOut.String()), "\n")
	ors := make([]types.OutResult, 0, len(list))

	for _, line := range list {
		var str bytes.Buffer
		str.WriteString(" ")
		str.WriteString(line)
		name := strings.Replace(filepath.Base(line), ".", "_", -1)
		ors = append(ors, types.OutResult{Path: line, Prio: 0, Name: name, Alias: str.String()})
	}

	// return stdRes.String(), nil
	ch <- ors
	// return ors, nil
}
