package zoxide

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"tmux-sessionmaster/types"
)

// func ZoxideCmd() ([]types.OutResult, error) {
func ZoxideCmd(wg *sync.WaitGroup, ch chan<- []types.OutResult) {
	defer wg.Done()
	if _, err := exec.LookPath("zoxide"); err != nil {
		return
	}

	var stdOut, stdErr bytes.Buffer
	cmdZoxide := exec.Command("zoxide", "query", "-ls")

	cmdZoxide.Stdout = &stdOut
	cmdZoxide.Stderr = io.MultiWriter(os.Stderr, &stdErr)

	if err := cmdZoxide.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmdZoxide.Wait(); err != nil {
		log.Fatal(err)
	}

	list := strings.Split(strings.TrimSpace(stdOut.String()), "\n")
	ors := make([]types.OutResult, 0, len(list))

	for _, line := range list {
		fields := strings.Split(strings.Trim(line, " "), " ")
		var str bytes.Buffer
		str.WriteString("󰰶 ")
		str.WriteString(fields[1])

		name := strings.Replace(filepath.Base(fields[1]), ".", "_", -1)
		if fields[1] == "/" {
			name = "root"
		}
		ors = append(ors, types.OutResult{Path: fields[1], Prio: 0, Name: name, Alias: str.String()})
	}

	// return stdRes.String(), nil
	ch <- ors
}

func AddToZoxide(path string) (string, error) {
	if strings.HasPrefix(path, "/") == false {
		return "", errors.New("no path")
	}
	p, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("can't add %q path to zoxide", path)
	}
	if _, err := exec.LookPath("zoxide"); err != nil {
		return "", nil
	}
	cmd := exec.Command("zoxide", "add", p)
	_, err = cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to add %q to zoxide: %w", p, err)
	}
	name := strings.Replace(filepath.Base(path), ".", "_", -1)
	return name, nil
}
