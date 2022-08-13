package commander

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/internal/storage"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/storage/local"

	"gitlab.ozon.dev/davidokk/reminder-manager/utils"
)

const (
	helpCommand   = "help"
	addCommand    = "add"
	listCommand   = "list"
	removeCommand = "rm"
	updateCommand = "update"
	getCommand    = "get"
)

var description = map[string]string{
	addCommand:    "[dd.mm.yy / today / tomorrow] [text] adds a new reminder",
	listCommand:   "shows all your plans in chronological order",
	removeCommand: "[id] removes record with given id",
	updateCommand: "[id] [new text] changes the reminder text",
	getCommand:    "[id] shows info about given id",
	helpCommand:   "show this menu",
}

const badArgumentResponse = "Bad argument, try one more time"
const successResponse = "Success! =)"

// AddHandlers registers handlers for given Commander
func AddHandlers(cmd *Commander) {
	cmd.RegisterHandler(listCommand, listFunc)
	cmd.RegisterHandler(addCommand, addFunc)
	cmd.RegisterHandler(removeCommand, removeFunc)
	cmd.RegisterHandler(updateCommand, updateFunc)
	cmd.RegisterHandler(getCommand, getFunc)

	var help string
	for name, desc := range description {
		help += fmt.Sprintf("/%s %s\n", name, desc)
	}

	cmd.RegisterHandler(helpCommand, func(string, storage.RemindersStorage) string { return help })
}

func getFunc(str string, storage storage.RemindersStorage) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return badArgumentResponse
	}
	rem, err := storage.GetReminder(ctx, id)
	if err != nil {
		return err.Error()
	}
	return rem.String()
}

func updateFunc(str string, storage storage.RemindersStorage) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	params := strings.Split(str, " ")
	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil || len(params) < 2 {
		return badArgumentResponse
	}
	if err := storage.UpdateReminder(ctx, id, strings.Join(params[1:], " ")); err != nil {
		return err.Error()
	}
	return successResponse
}

func removeFunc(params string, storage storage.RemindersStorage) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	id, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		return badArgumentResponse
	}
	if err := storage.RemoveReminder(ctx, id); err != nil {
		return err.Error()
	}
	return successResponse
}

func listFunc(str string, storage storage.RemindersStorage) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := storage.ListReminders(ctx)
	if err != nil {
		return err.Error()
	}
	if len(res) == 0 {
		return "You haven't planned anything yet"
	}
	return fmt.Sprintf("Your plans\n\n%s", strings.Join(local.AsStrings(res), "\n"))
}

func addFunc(str string, storage storage.RemindersStorage) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
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
	if _, err := storage.CreateReminder(ctx, date, strings.Join(params[1:], " ")); err != nil {
		return err.Error()
	}
	return successResponse
}
