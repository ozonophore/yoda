package service

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type reportInterface interface {
	Print(date time.Time, finaName string) error
}

var (
	// Menu texts
	firstMenu  = "<b>Menu</b>\n\nОтчет."
	secondMenu = "<b>Menu 2</b>\n\nA better menu with even more shiny inline buttons."

	// Button texts
	nextButton     = "Next"
	backButton     = "Back"
	tutorialButton = "Tutorial"

	yesterdayButton    = "За вчера"
	twoDaysAgoButton   = "2 дня назад"
	threeDaysAgoButton = "3 дня назад"

	// Store bot screaming status
	screaming = false
	bot       *tgbotapi.BotAPI
	service   reportInterface

	// Keyboard layout for the first menu. One button, one row
	firstMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(yesterdayButton, yesterdayButton),
			tgbotapi.NewInlineKeyboardButtonData(twoDaysAgoButton, twoDaysAgoButton),
			tgbotapi.NewInlineKeyboardButtonData(threeDaysAgoButton, threeDaysAgoButton),
		),
	)

	// Keyboard layout for the second menu. Two buttons, one per row
	secondMenuMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(backButton, backButton),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL(tutorialButton, "https://core.telegram.org/bots/api"),
		),
	)
)

func StartBot(token string, ctx context.Context, srv reportInterface) {
	logrus.Info("Starting bot ", token)
	service = srv
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	logrus.Infof("Authorized on account %s", bot.Self.UserName)
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	// Pass cancellable context to goroutine
	receiveUpdates(ctx, updates)
}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update tgbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handleMessage(update.Message)
		break

	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
		break
	}
}

func handleButton(query *tgbotapi.CallbackQuery) {
	var text string

	//markup := tgbotapi.NewInlineKeyboardMarkup()
	message := query.Message

	//if query.Data == nextButton {
	//	text = secondMenu
	//	markup = secondMenuMarkup
	//} else if query.Data == backButton {
	//	text = firstMenu
	//	markup = firstMenuMarkup
	//} else
	var date time.Time
	if query.Data == yesterdayButton {
		text = "Отчет за вчера"
		date = time.Now().AddDate(0, 0, -1)
	} else if query.Data == twoDaysAgoButton {
		text = "Отчет за 2 дня назад"
		date = time.Now().AddDate(0, 0, -2)
	} else if query.Data == threeDaysAgoButton {
		text = "Отчет за 3 дня назад"
		date = time.Now().AddDate(0, 0, -3)
	} else {
		text = "Unknown button"
	}
	fileName := fmt.Sprintf(`%s%s.xlsx`, os.TempDir(), date.Format("2006-01-02"))
	defer func() {
		os.Remove(fileName)
	}()
	err := service.Print(date, fileName)
	var msg tgbotapi.Chattable
	if err == nil {
		doc := tgbotapi.NewDocument(message.Chat.ID, tgbotapi.FilePath(fileName))
		doc.Caption = text
		msg = doc
	} else {
		msg = tgbotapi.NewMessage(message.Chat.ID, err.Error())
	}

	//callbackCfg := tgbotapi.NewCallback(query.ID, "")
	//bot.Send(callbackCfg)
	//
	//// Replace menu text and keyboard
	//msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	//msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func handleMessage(message *tgbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	// Print to console
	logrus.Debugf("%s wrote %s", user.FirstName, text)

	var err error
	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, text)
	} else if screaming && len(text) > 0 {
		msg := tgbotapi.NewMessage(message.Chat.ID, strings.ToUpper(text))
		// To preserve markdown, we attach entities (bold, italic..)
		msg.Entities = message.Entities
		_, err = bot.Send(msg)
	} else {
		// This is equivalent to forwarding, without the sender's name
		copyMsg := tgbotapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
		_, err = bot.CopyMessage(copyMsg)
	}

	if err != nil {
		logrus.Errorf("An error occured: %s", err.Error())
	}
}

// When we get a command, we react accordingly
func handleCommand(chatId int64, command string) error {
	var err error

	switch command {
	case "/scream":
		screaming = true
		break

	case "/whisper":
		screaming = false
		break

	case "/menu":
		err = sendMenu(chatId)
		break
	}

	return err
}

func sendMenu(chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, firstMenu)
	msg.ParseMode = tgbotapi.ModeHTML
	msg.ReplyMarkup = firstMenuMarkup
	_, err := bot.Send(msg)
	return err
}
