package cli

import (
	"errors"
	"fmt"
)

//Controller ...
type Controller func([]string) (error, bool)

type command struct {
	name string
	property []string
	function Controller
}

type App struct {
	listsOfCommands []command
}

//NewApp ...
func NewApp() App {
	app := App{
		listsOfCommands: []command{},
	}
	return app
}


//AddCommand ...
func (app *App) AddCommand(name string, foo Controller) {
	c := command{
		name: name,
		property: []string{},
		function: foo,
	}
	app.listsOfCommands = append(app.listsOfCommands, c)
}


//Run ...
func (app *App) Run(name []string) (error, bool) {
	if len(name) == 0 {
		return errors.New("no argument given"), false
	}

	for i := 0; i < len(app.listsOfCommands); i++ {
		if app.listsOfCommands[i].name == name[0] {
			return app.listsOfCommands[i].function(name[1:])
		}
	}

	return errors.New(fmt.Sprintf("no command found with name : %v", name)), false
}