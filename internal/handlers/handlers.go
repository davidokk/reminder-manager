package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/internal/commander"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/storage"
	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

const (
	helpCommand       = "help"
	addCommand        = "add"
	listCommand       = "list"
	removeByIdCommand = "rm"
	editCommand       = "edit"
)

var description = map[string]string{
	addCommand:        "[dd.mm.yy / today / tomorrow] [text] adds a new reminder",
	listCommand:       "shows all your plans in chronological order",
	removeByIdCommand: "[id] removes record with given id",
	editCommand:       "[id] [new text] changes the reminder text",
	helpCommand:       "show this menu",
}

const badArgumentResponse = "Bad argument, try one more time"
const successResponse = "Success! =)"

// AddHandlers registers handlers for given Commander
func AddHandlers(cmd *commander.Commander) {
	cmd.RegisterHandler(listCommand, listFunc)
	cmd.RegisterHandler(addCommand, addFunc)
	cmd.RegisterHandler(removeByIdCommand, removeByIdFunc)
	cmd.RegisterHandler(editCommand, editFunc)

	var help string
	for name, desc := range description {
		help += fmt.Sprintf("/%s %s\n", name, desc)
	}

	cmd.RegisterHandler(helpCommand, func(string) string { return help })
}

func editFunc(str string) string {
	params := strings.Split(str, " ")
	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil || len(params) < 2 {
		return badArgumentResponse
	}
	if err := storage.Edit(id, strings.Join(params[1:], " ")); err != nil {
		return err.Error()
	}
	return successResponse
}

func removeByIdFunc(params string) string {
	id, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		return badArgumentResponse
	}
	if err := storage.RemoveById(id); err != nil {
		return err.Error()
	}
	return successResponse
}

func listFunc(string) string {
	res, err := storage.Data()
	if err != nil {
		return err.Error()
	}
	if len(res) == 0 {
		return "You haven't planned anything yet"
	}
	return fmt.Sprintf("Your plans\n\n%s", strings.Join(storage.AsStrings(res), "\n"))
}

func addFunc(str string) string {
	params := strings.Split(str, " ")
	var date time.Time
	if params[0] == "today" {
		date = utils.UpToDay(time.Now())
	} else if params[0] == "tomorrow" {
		date = utils.UpToDay(time.Now()).Add(time.Hour * 24)
	} else {
		var err error
		date, err = time.Parse("02.01.06", params[0])
		if err != nil || len(params) < 2 {
			return badArgumentResponse
		}
	}
	if err := storage.Add(storage.NewReminder(strings.Join(params[1:], " "), date)); err != nil {
		return err.Error()
	}
	return successResponse
}
