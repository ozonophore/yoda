package service

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type reportInterface interface {
	Print(date time.Time, finaName string) error
}

type repositoryInterface interface {
	AddGroup(userName, groupName string, chatId int64) error
	DeleteGroup(userName, groupName string) error
}

var (
	mutax = &sync.Mutex{}
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
	screaming  = false
	bot        *tgbotapi.BotAPI
	service    reportInterface
	repository repositoryInterface
	baseDir    string

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

type Bot struct {
	ctx  context.Context
	name string
}

func NewBot(token, workDir string, ctx context.Context, srv reportInterface, rep repositoryInterface) *Bot {
	logrus.Info("Starting bot ", token)
	baseDir = workDir
	service = srv
	repository = rep
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	user, err := bot.GetMe()
	if err != nil {
		panic(err)
	}
	return &Bot{
		ctx:  ctx,
		name: user.UserName,
	}
}

func (b *Bot) GetSender() string {
	return b.name
}

func (b *Bot) StartBot() {
	logrus.Infof("Authorized on account %s", bot.Self.UserName)
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	// Pass cancellable context to goroutine
	receiveUpdates(b.ctx, updates)
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
	case update.MyChatMember != nil:
		handleMember(update.MyChatMember)
		break
	}
}

func handleMember(member *tgbotapi.ChatMemberUpdated) {
	if member.NewChatMember.Status == "member" && member.Chat.Type == "group" {
		handleNewChatMembers(member.Chat.ID, member.NewChatMember.User.UserName, member.Chat.Title)
	} else if member.NewChatMember.Status == "left" && member.Chat.Type == "group" {
		handleLeftChatMember(member.Chat.ID, member.NewChatMember.User.UserName, member.Chat.Title)
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

	msg := createReport(message.Chat.ID, text, date)

	//callbackCfg := tgbotapi.NewCallback(query.ID, "")
	//bot.Send(callbackCfg)
	//
	//// Replace menu text and keyboard
	//msg := tgbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	//msg.ParseMode = tgbotapi.ModeHTML
	bot.Send(msg)
}

func createReport(chatId int64, text string, date time.Time) tgbotapi.Chattable {
	mutax.Lock()
	defer mutax.Unlock()
	fileName := fmt.Sprintf(`%s.xlsx`, date.Format("2006-01-02"))
	absoluteFilePath := path.Join(baseDir, fileName)
	_, err := os.Stat(absoluteFilePath)
	if err != nil && os.IsNotExist(err) {
		//If file exists
		err = service.Print(date, absoluteFilePath)
	}

	var msg tgbotapi.Chattable
	if err == nil {
		doc := tgbotapi.NewDocument(chatId, tgbotapi.FilePath(absoluteFilePath))
		doc.Caption = text
		msg = doc
	} else {
		msg = tgbotapi.NewMessage(chatId, err.Error())
	}
	return msg
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
		command := strings.Split(text, "@")
		err = handleCommand(message.Chat.ID, command[0])
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

func handleLeftChatMember(chatId int64, userName, groupName string) error {
	msg := tgbotapi.NewMessage(chatId, "Пока, я буду скучать")
	repository.DeleteGroup(userName, groupName)
	_, err := bot.Send(msg)
	return err
}

func handleNewChatMembers(chatId int64, userName, groupName string) error {
	msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Привет %s, я бот для отчетов", groupName))
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return repository.AddGroup(userName, groupName, chatId)
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

func (*Bot) SendReport(date time.Time, chatIds *[]int64) error {
	if len(*chatIds) == 0 {
		return nil
	}
	for _, chatId := range *chatIds {
		msg := createReport(chatId, fmt.Sprintf("Отчет за %s", date.Format(time.DateOnly)), date)
		bot.Send(msg)
	}
	return nil
}
