package tmux

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"tmux-sessionmaster/types"
	"tmux-sessionmaster/zoxide"
)

func Session(path string, outs []types.OutResult) (string, string, error) {
	splitPath := strings.SplitN(path, "--", 2)
	input := strings.TrimSpace(splitPath[0])

	envEditor := os.Getenv("EDITOR")

	var argument string
	if len(splitPath) == 2 {
		argument = strings.TrimSpace(splitPath[1])
	}

	index, found := types.SearchOutResult(input, outs)
	_, isAttached := types.SearchOutResultAttached(outs)

	opts := [][]string{}

	switch found {
	case false:

		name := input
		opt := []string{"new-session", "-ds"}
		if nameZox, err := zoxide.AddToZoxide(input); err != nil {
			if strings.HasPrefix(strings.TrimSpace(err.Error()), "no path") == false {
				return "", "", err
			}
			opt = append(opt, name)

		} else {
			name = nameZox
			opt = append(opt, nameZox, "-c", input)
		}

		if argument == "edit" {
			opt = append(opt, envEditor, "./")
		}

		opts = append(opts, opt)

		if isAttached == true {
			opts = append(opts, []string{"switch", "-t", name})
		} else {
			opts = append(opts, []string{"attach", "-t", name})
		}

	case true:
		out := outs[index]
		switch argument {
		case "delete":
			opts = append(opts, []string{"kill-session", "-t", out.Name})
		case "rename":
			newName, err := readInput(out.Name)
			if err == nil && newName != "" {
				opts = append(opts, []string{"rename-session", "-t", out.Name, newName})
			}
		default:
			switch out.Prio {
			case 1:
				opts = append(opts, []string{"detach"})
			case 2:
				if isAttached == true {
					opts = append(opts, []string{"switch", "-t", out.Name})
				} else {
					opts = append(opts, []string{"attach", "-t", out.Name})
				}
			case 0:
				opt := []string{"new-session", "-ds", out.Name, "-c", out.Path}
				// opts = append(opts, []string{"new-session", "-ds", out.Name, "-c", out.Path})
				if argument == "edit" {
					opt = append(opt, envEditor, "./")
				}
				opts = append(opts, opt)

				if isAttached == true {
					opts = append(opts, []string{"switch", "-t", out.Name})
				} else {
					opts = append(opts, []string{"attach", "-t", out.Name})
				}
			}
		}
	}

	if len(opts) > 1 {
		for _, opt := range opts[:len(opts)-1] {
			if _, err := sessionCmd(opt); err != nil {
				return "", "", err
			}
		}
	}

	if len(opts) == 0 {
		return argument, "nothing to do for tmux", nil
	}

	res, err := sessionCmd(opts[len(opts)-1])
	if err != nil {
		return "", "", err
	}

	return argument, res, nil
}

func sessionCmd(opt []string) (string, error) {
	var stdOut, stdErr bytes.Buffer
	cmd := exec.Command("tmux", opt...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = &stdOut
	cmd.Stderr = io.MultiWriter(&stdErr, os.Stderr)

	if err := cmd.Start(); err != nil {
		return "", errors.New("start error")
	}
	if err := cmd.Wait(); err != nil {
		return "", errors.New("wait error sessioncmd:" + stdOut.String())
	}

	return stdOut.String(), nil
}

func readInput(currentName string) (string, error) {
	fmt.Printf("Rename Session %s in: ", currentName)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("An error occured while reading input.")
	}

	return strings.TrimSpace(strings.TrimSuffix(input, "\n")), nil
}

// func nvimCmd(opt []string) (string, error) {
// 	envEditor := os.Getenv("EDITOR")
// 	cmd := exec.Command(envEditor, opt...)
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
//
// 	if err := cmd.Start(); err != nil {
// 		return "", err
// 	}
//
// 	if err := cmd.Wait(); err != nil {
// 		return "", err
// 	}
//
// 	return "", nil
// }
