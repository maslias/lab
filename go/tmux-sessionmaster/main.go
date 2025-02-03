package main

import (
	"log"
	"sync"
	"tmux-sessionmaster/find"
	"tmux-sessionmaster/fzf"
	"tmux-sessionmaster/tmux"
	"tmux-sessionmaster/types"
	"tmux-sessionmaster/zoxide"
)

func main() {
	runApp()
}

func runApp() {
	// starttime := time.Now()

	var wg sync.WaitGroup
	ch := make(chan []types.OutResult)

	wg.Add(3)
	go tmux.TmuxCmd(&wg, ch)
	go zoxide.ZoxideCmd(&wg, ch)
	go find.FindCmd(&wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	cmdOuts := [][]types.OutResult{<-ch, <-ch, <-ch}

	outs := types.SortOutResult(cmdOuts...)
	fzfOut, _ := fzf.FzfCmd(outs)

	if fzfOut == "" {
		return
	}

	arg, _, err := tmux.Session(fzfOut, outs)
	if err != nil {
		log.Fatal(err)
	}
	if arg == "rename" || arg == "delete" {
		runApp()
	}

	// fmt.Printf("peformance %v", time.Since(starttime))
}
