package handlers

import (
	"fmt"
	"reminder-manager/internal/commander"
	"reminder-manager/internal/storage"
	"reminder-manager/utils"
	"strconv"
	"strings"
	"time"
)

const (
	helpCommand           = "help"
	addCommand            = "add"
	listCommand           = "list"
	removeOutdatedCommand = "rmoutdated"
	removeByIdCommand     = "remove"
	editCommand           = "edit"
	todayCommand          = "today"
	forDaysCommand        = "fordays"
)

var description = map[string]string{
	addCommand:            "[dd.mm.yy / today / tomorrow] [text] adds a new reminder",
	listCommand:           "shows all your plans in chronological order",
	removeOutdatedCommand: "removes outdated records",
	removeByIdCommand:     "[id] removes record with given id",
	editCommand:           "[id] [new text] changes the reminder text",
	todayCommand:          "shows today's activities",
	forDaysCommand:        "[count] shows records for next 'count' days",
	helpCommand:           "show this menu",
}

const BadArgumentResponse = "Bad argument, try one more time"
const SuccessResponse = "Success! =)"

func AddHandlers(cmd *commander.Commander) {
	cmd.RegisterHandler(listCommand, listFunc)
	cmd.RegisterHandler(addCommand, addFunc)
	cmd.RegisterHandler(removeOutdatedCommand, removeOutdatedFunc)
	cmd.RegisterHandler(removeByIdCommand, removeByIdFunc)
	cmd.RegisterHandler(editCommand, editFunc)
	cmd.RegisterHandler(todayCommand, todayFunc)
	cmd.RegisterHandler(forDaysCommand, forDaysFunc)

	var help string
	for name, desc := range description {
		help += fmt.Sprintf("/%s %s\n", name, desc)
	}

	cmd.RegisterHandler(helpCommand, func(string) string { return help })
}

func forDaysFunc(param string) string {
	cnt, err := strconv.Atoi(param)
	if err != nil || cnt < 1 {
		return BadArgumentResponse
	}
	rem := storage.AsStrings(storage.RemindersForDays(cnt))
	if rem == nil {
		return fmt.Sprintf("Nothing to do next %d days =(", cnt)
	} else {
		return fmt.Sprintf("%d things to do next %d days\n\n%s", len(rem), cnt, strings.Join(rem, "\n"))
	}
}

func todayFunc(string) string {
	rem := storage.AsStrings(storage.RemindersForDays(1))
	if rem == nil {
		return "Nothing to do today =("
	} else {
		return fmt.Sprintf("%d things to do today\n\n%s", len(rem), strings.Join(rem, "\n"))
	}
}

func editFunc(str string) string {
	params := strings.Split(str, " ")
	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil || len(params) < 2 {
		return BadArgumentResponse
	}
	if storage.Edit(id, strings.Join(params[1:], " ")) {
		return SuccessResponse
	} else {
		return "I can't find this id"
	}
}

func removeByIdFunc(params string) string {
	id, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		return BadArgumentResponse
	}
	if storage.RemoveById(id) {
		return SuccessResponse
	} else {
		return "I can't find this id"
	}
}

func removeOutdatedFunc(string) string {
	outdated := storage.RemoveOutdated()
	if outdated == 0 {
		return "There aren't outdated records"
	} else {
		return fmt.Sprintf("%d records were deleted", outdated)
	}
}

func listFunc(string) string {
	res := storage.AsStrings(storage.Data())
	if len(res) == 0 {
		return "You haven't planned anything yet"
	}
	oldCount := storage.OutdatedCount()
	var outdated string
	var actual string
	if oldCount > 0 {
		outdated += "There are outdated entries on your list\n\n"
		outdated += strings.Join(res[:oldCount], "\n")
	}
	if len(res)-int(oldCount) > 0 {
		actual += "Your actual plans\n\n"
		actual += strings.Join(res[oldCount:], "\n")
	}
	if len(outdated) != 0 && len(actual) != 0 {
		return outdated + "\n\n" + actual
	} else {
		return outdated + actual
	}
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
		date, err = time.Parse("02.01.06", strings.TrimSpace(params[0]))
		if err != nil || len(params) < 2 {
			return BadArgumentResponse
		}
	}
	storage.Add(storage.NewReminder(strings.Join(params[1:], " "), date))
	return SuccessResponse
}
