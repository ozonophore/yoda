package mqclient

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/webapp/internal/config"
	"github.com/yoda/webapp/internal/dao"
)

const token = "6232820413:AAGumULdIEF8ClJkTnaVJs5V6MXCl90vMwI"

type botLogger struct {
	logger *logrus.Logger
}

func (b botLogger) Println(v ...interface{}) {
	b.logger.Debug(v)
}

func (b botLogger) Printf(format string, v ...interface{}) {
	b.logger.Debugf(format, v)
}

var logger = logrus.New()
var bot *tgbotapi.BotAPI

func genearateMenuButtons(chatId int64) tgbotapi.InlineKeyboardMarkup {
	var button tgbotapi.InlineKeyboardButton
	if dao.ExistsTlgEventByChatIdAndEvent(chatId, "etl-info") {
		button = tgbotapi.NewInlineKeyboardButtonData("Информация о загрузке(откл)", "etl-info-off")
	} else {
		button = tgbotapi.NewInlineKeyboardButtonData("Информация о загрузке(вкл)", "etl-info")
	}
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			button,
		), tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("button3", "button3"),
			tgbotapi.NewInlineKeyboardButtonData("button4", "button4"),
		), tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("button5", "button5"),
		),
	)
}

func StartTgBot(context context.Context, config config.TelegramBot, repository *dao.Repository) error {
	// Start telegram bot
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	logger.Level, err = logrus.ParseLevel(config.LoggingLevel)
	if err != nil {
		return err
	}
	bot.Debug = true
	tgbotapi.SetLogger(botLogger{logger: logger})
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = config.UpdateTimeOut
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		select {
		case <-context.Done():
			logger.Info("Telegram bot stopped")
			return nil
		default:
			if update.Message != nil {
				handleMessage(update.Message, bot, repository)
			}
			if update.CallbackQuery != nil {
				handelCallbackQuery(update.CallbackQuery, bot, repository)
			}
		}
	}

	return nil
}

func handelCallbackQuery(query *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI, repository *dao.Repository) {
	switch query.Data {
	case "help":
		msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Help")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
	case "etl-info":
		if dao.ExistsTlgEventByChatIdAndEvent(query.Message.Chat.ID, "etl-info") {
			return
		}
		m, err := repository.CreateTlgMessage(query.Message.Chat.ID, query.Message.Chat.Type, query.Message.MessageID, &query.Data)
		if errorIf(err) {
			return
		}
		dao.CreateTlgEvent(query.Message.Chat.ID, "etl-info")
		repository.SetTlgMessageStatus(m.ID, model.TlgMessageStatusComplited)
	case "etl-info-off":
		if dao.ExistsTlgEventByChatIdAndEvent(query.Message.Chat.ID, "etl-info") {
			return
		}
		m, err := repository.CreateTlgMessage(query.Message.Chat.ID, query.Message.Chat.Type, query.Message.MessageID, &query.Data)
		if errorIf(err) {
			return
		}
		dao.DeleteTlgEvent(query.Message.Chat.ID, "etl-info")
		repository.SetTlgMessageStatus(m.ID, model.TlgMessageStatusComplited)
	}
}

func handleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI, repository *dao.Repository) {
	//var itemsSel = tgbotapi.NewReplyKeyboard(
	//	tgbotapi.NewKeyboardButtonRow(
	//		tgbotapi.NewKeyboardButton("button1"),
	//		tgbotapi.NewKeyboardButton("button2")),
	//)
	//msg := tgbotapi.NewMessage(message.Chat.ID, "Main Menu")
	user, _ := bot.GetMe()
	//m, err := repository.CreateTlgMessage(message.Chat.ID, message.Chat.Type, message.MessageID, &message.Text)
	if message.IsCommand() {
		switch message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to the bot!")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			//repository.SetTlgMessageStatus(m.ID, model.TlgMessageStatusPending)
			return
		case "help":
			msg := tgbotapi.NewMessage(message.Chat.ID, "Help")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			//repository.SetTlgMessageStatus(m.ID, model.TlgMessageStatusPending)
			return
		case "menu":
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(`Settings for @%s`, user.UserName))
			msg.ReplyMarkup = genearateMenuButtons(message.Chat.ID)
			bot.Send(msg)
			return
		}
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID
	bot.Send(msg)
	repository.CreateTlgMessage(message.Chat.ID, message.Chat.Type, message.MessageID, &message.Text)
}

func SendMessage(chatID int64, messageID int, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyToMessageID = messageID
	bot.Send(msg)
}

func errorIf(err error) bool {
	if err != nil {
		logger.Error(err)
		return true
	}
	return false
}
