package fzf

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"
	"tmux-sessionmaster/types"
)

func FzfCmd(outs []types.OutResult) (string, error) {
	var input bytes.Buffer
	for _, or := range outs {
		input.WriteString(or.Alias)
		input.WriteString("\n")
	}

	if _, err := exec.LookPath("fzf"); err != nil {
		return "", err
	}

	var stdoutBuf bytes.Buffer
	cmd := exec.Command("fzf", "--print-query", "--header", "C-D => delete | C-R => rename | C-V => nvim", "--bind", "ctrl-d:become(echo {} --delete)", "--bind", "ctrl-r:become(echo {} --rename)", "--bind", "ctrl-e:become(echo {} --edit)", "--bind", "ctrl-a:become(echo {} --add)")
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input.String())
	}()

	if err := cmd.Start(); err != nil {
		return "", errors.New("start error")
	}

	if err := cmd.Wait(); err != nil {
		list := strings.Split(strings.TrimSpace(stdoutBuf.String()), "\n")
		if len(list) == 1 && strings.TrimSpace(list[0]) != "" {
			// fmt.Println("fzf value check: " + list[0])

			return list[0], nil
		}
		return "", errors.New("wait error: " + stdoutBuf.String())
	}

	list := strings.Split(strings.TrimSpace(stdoutBuf.String()), "\n")
	return list[len(list)-1], nil

	// return strings.TrimSpace(stdoutBuf.String()), nil
}
