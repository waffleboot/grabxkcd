package grabxkcd

import (
	"os"
)

type appEnv struct {
}

func (app *appEnv) fromArgs([]string) error {
	return nil
}

func CLI(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	return 0
}
