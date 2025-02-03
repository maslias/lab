package tmux

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"tmux-sessionmaster/types"
)

func TmuxCmdPiping() (string, error) {
	if _, err := exec.LookPath("tmux"); err != nil {
		return "", err
	}

	if _, err := exec.LookPath("awk"); err != nil {
		return "", err
	}

	if _, err := exec.LookPath("sed"); err != nil {
		return "", err
	}

	var stdOut, stdErr bytes.Buffer

	cmdTmux := exec.Command("tmux", "ls")
	cmdAwk := exec.Command("awk", "{print \":\" $1 \"󰓩:\"$2\":\" $10}")
	cmdSed := exec.Command("sed", "-e", "s/:/ /g", "-e", "s/(/󰌹 /g", "-e", "s/)//g")

	cmdAwk.Stdin, _ = cmdTmux.StdoutPipe()
	cmdSed.Stdin, _ = cmdAwk.StdoutPipe()

	cmdSed.Stdout = &stdOut
	cmdSed.Stderr = io.MultiWriter(os.Stderr, &stdErr)

	if err := cmdTmux.Start(); err != nil {
		return "", err
	}

	if err := cmdAwk.Start(); err != nil {
		return "", err
	}

	if err := cmdSed.Start(); err != nil {
		return "", err
	}

	if err := cmdTmux.Wait(); err != nil {
		return "", err
	}

	if err := cmdAwk.Wait(); err != nil {
		return "", err
	}
	if err := cmdSed.Wait(); err != nil {
		return "", err
	}

	return stdOut.String(), nil
}

// func TmuxCmd() ([]types.OutResult, error) {
func TmuxCmd(wg *sync.WaitGroup, ch chan<- []types.OutResult) {
	defer wg.Done()
	if _, err := exec.LookPath("tmux"); err != nil {
		log.Fatal(err)
	}

	var stdOut, stdErr bytes.Buffer

	cmdTmux := exec.Command("tmux", "ls", "-F", "#{session_name} #{session_path} #{session_windows} #{session_attached}")

	cmdTmux.Stdout = &stdOut
	cmdTmux.Stderr = &stdErr // io.MultiWriter(os.Stderr, &stdErr)

	if err := cmdTmux.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmdTmux.Wait(); err != nil {
		if strings.HasPrefix(strings.TrimSpace(stdErr.String()), "no server running") {
			// return nil, nil
			return
		}
		log.Fatal(err)
	}

	list := strings.Split(strings.TrimSpace(stdOut.String()), "\n")
	ors := make([]types.OutResult, 0, len(list))
	for _, line := range list {
		fields := strings.Split(line, " ")

		var str bytes.Buffer
		str.WriteString(" ")
		str.WriteString(strings.Replace(fields[0], "_", ".", -1))
		str.WriteString(" 󰓩 ")
		str.WriteString(fields[2])
		or := types.OutResult{Prio: 2, Name: fields[0], Path: fields[1]}

		if fields[3] == "1" {
			str.WriteString(" 󰌹 attached")
			or.Prio = 1
		}
		or.Alias = str.String()

		ors = append(ors, or)

	}

	// return stdRes.String(), nil
	ch <- ors
	// return ors, nil
}
