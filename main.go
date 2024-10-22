package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"tibot/src/app"
	"tibot/src/tibot"
	"tibot/src/usermanager"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

var userManager = usermanager.NewUserManager()

// Send any text message to the bot after the bot has been started

func main() {

	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}

	// Load config
	config := app.NewConfig()

	// Add admin user
	userManager.AddUser(config.AdminUserId, "admin", true)

	ibot := tibot.New()
	ibot.SetUserManager(userManager)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(ibot.Handlers.Default),
	}

	// init telegram bot
	tbot, err := bot.New(config.TelegramApiKey, opts...)
	if err != nil {
		// panic(err)
		log.Printf("ERROR:\n %s\n", err)
	}

	// set botInfo
	botinfo, err := tbot.GetMe(ctx)
	ibot.SetBotInfo(botinfo)
	// set the defaulf handlertype
	ibot.SetDefaultHandlerType(bot.HandlerTypeMessageText)
	// set default python run
	ibot.SetRunPythonFunc(ibot.Handlers.RunPython)

	// TiBOT default handlers
	// add slug + handler
	ibot.AddRule(ibot.GetDefaultHandlerType(), `^/start`, ibot.Handlers.MyStart)
	ibot.AddRule(ibot.GetDefaultHandlerType(), `^/args`, ibot.Handlers.GetArgs)
	ibot.AddRule(ibot.GetDefaultHandlerType(), `^/myinfo`, ibot.Handlers.GetUserInfo)
	ibot.AddRule(ibot.GetDefaultHandlerType(), `^/botinfo`, ibot.Handlers.GetBotInfo)
	ibot.AddRule(ibot.GetDefaultHandlerType(), `^/run help`, ibot.Handlers.PythonHelp)
	ibot.AddRule(ibot.GetDefaultHandlerType(), `^/run`, ibot.Handlers.RunPython)
	ibot.AddRule(ibot.GetDefaultHandlerType(), `^/die`, ibot.Handlers.Die)

	ibot.InitRules(tbot)
	// ibot.InitPythonRules(tbot)

	if err != nil {
		fmt.Println(err)
	}

	app.AppBanner(ibot.GetBotInfo())

	tbot.Start(ctx)
}
