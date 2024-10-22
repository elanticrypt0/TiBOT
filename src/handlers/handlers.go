package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"tibot/src/app"
	"tibot/src/pyrunner"
	"tibot/src/usermanager"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Handlers struct {
	UserManager *usermanager.UserManager
	BotSettings BotSettings
}

func New() Handlers {
	return Handlers{}
}

func (me *Handlers) SetUserManager(usermanagerInstace *usermanager.UserManager) {
	me.UserManager = usermanagerInstace
}

func (me *Handlers) setBotSettings(ctx context.Context, b *bot.Bot, update *models.Update) {
	me.BotSettings = NewBotSettings(ctx, b, update)
}

func (me *Handlers) SendResponseSimple(message string) {
	me.BotSettings.Bot.SendMessage(me.BotSettings.Ctx, &bot.SendMessageParams{
		ChatID: me.BotSettings.Update.Message.Chat.ID,
		Text:   message,
	})
}

func (me *Handlers) getArgs(text string) (int, []string) {
	args := strings.Split(text, " ")
	return len(args) - 1, args[1:]

}

func (me *Handlers) Default(ctx context.Context, b *bot.Bot, update *models.Update) {

	// fmt.Printf("%#v -> %s (%s %s) \n", update.Message.From.ID, update.Message.From.Username, update.Message.From.FirstName, update.Message.From.LastName)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "What?",
	})
}

func (me *Handlers) MyStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	userInfo := update.Message.From
	response := fmt.Sprintf("Hi %s , %s", userInfo.Username, update.Message.Text)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}

// TODO: Falta devolver la información del usuaario.
func (me *Handlers) GetUserInfo(ctx context.Context, b *bot.Bot, update *models.Update) {
	userInfo := update.Message.From
	userContact := update.Message.Contact
	response := fmt.Sprintf("Your info \n %#v \n %#v", userInfo, userContact)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}

func (me *Handlers) GetArgs(ctx context.Context, b *bot.Bot, update *models.Update) {

	argsQty, args := me.getArgs(update.Message.Text)

	response := fmt.Sprintf("Args (%d): %#v", argsQty, args)
	fmt.Println(response)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   response,
	})
}

// TODO: Falta devolver la información del bot
func (me *Handlers) GetBotInfo(ctx context.Context, b *bot.Bot, update *models.Update) {

	info, err := b.GetMe(ctx)
	if err != nil {
		log.Fatal(err)
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   info.Username,
	})
}

// TODO: matar al bot
func (me *Handlers) Die(ctx context.Context, b *bot.Bot, update *models.Update) {

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Me muero",
	})
}

// TODO: Ejecutar script de python y obtener la respuesta
func (me *Handlers) RunPython(ctx context.Context, b *bot.Bot, update *models.Update) {

	fmt.Println("------------------------")
	userID := strconv.Itoa(int(update.Message.From.ID))

	me.setBotSettings(ctx, b, update)

	me.SendResponseSimple("📜 Running script")

	log.Printf("::: Running Script :::\n")
	log.Printf("Excecuted by User: %q(%s)\n", update.Message.From.Username, userID)

	argsQty, args := me.getArgs(update.Message.Text)

	if argsQty > 0 {
		pyr := pyrunner.New()
		script, err := pyr.GetScript(args[0])
		me.SendResponseSimple(">>> " + script.Path)

		if err != nil {
			log.Printf("🛑 Error:\n %s", err)
			me.SendResponseSimple(fmt.Sprintf("🛑 Error: %s", err))
			return
		}

		if script.OnlyAdmin && !app.IsAdmin(userID) {
			errMsg := "Just for admin"
			log.Printf("🛑 Error: %s", errMsg)
			me.SendResponseSimple(fmt.Sprintf("🛑 Error: %s", errMsg))
			return
		}

		args[0] = script.Path
		response, err := pyr.RunScript(script.Engine, args)

		if err != nil {
			log.Printf("🛑 Error: %s", err)
			me.SendResponseSimple(fmt.Sprintf("🛑 Error:\n %s", err))
			return
		}

		fmt.Printf("\nRESPONSE:\n%v", response)

		me.SendResponseSimple(response)

		fmt.Println("------------------------")
		return
	} else {
		me.SendResponseSimple("No script selected")
		return
	}

	return
}

func (me *Handlers) PythonHelp(ctx context.Context, b *bot.Bot, update *models.Update) {

	me.setBotSettings(ctx, b, update)

	me.SendResponseSimple("📜 Script availables")

	pyr := pyrunner.New()

	scriptsStr := ""
	for _, script := range pyr.Config {
		scriptsStr += fmt.Sprintf("\n%s (%q)\n", script.Handler, script.Engine)
	}
	me.SendResponseSimple(fmt.Sprintf("		%s\n", scriptsStr))

	return
}
