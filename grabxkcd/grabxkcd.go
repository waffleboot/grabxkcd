package grabxkcd

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

type appEnv struct {
	hc         http.Client
	comicNo    int
	saveImage  bool
	outputJSON bool
}

const LatestComic = 0

func (app *appEnv) fromArgs(args []string) error {
	// Shallow copy of default client
	app.hc = *http.DefaultClient
	fl := flag.NewFlagSet("xkcd-grab", flag.ContinueOnError)
	fl.IntVar(&app.comicNo, "n", LatestComic, "Comic number to fetch (default latest)")
	fl.DurationVar(&app.hc.Timeout, "t", 30*time.Second, "Client timeout")
	return nil
}

func (app *appEnv) run() error {
	return nil
}

func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}
