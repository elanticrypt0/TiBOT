package tibot

import (
	"regexp"
	"tibot/src/handlers"
	"tibot/src/usermanager"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type TiBOT struct {
	botInfo            *models.User
	UserManager        *usermanager.UserManager
	Handlers           handlers.Handlers
	Rules              []Rule
	PythonRules        []Rule
	DefaultHandlerType bot.HandlerType
	RunPythonFunc      bot.HandlerFunc
}

func New() *TiBOT {
	return &TiBOT{
		Handlers: handlers.New(),
	}
}

func (me *TiBOT) SetBotInfo(botinfo *models.User) {
	me.botInfo = botinfo
}

func (me *TiBOT) GetBotInfo() *models.User {
	return me.botInfo
}

func (me *TiBOT) SetUserManager(userManager *usermanager.UserManager) {
	me.UserManager = userManager
}

func (me *TiBOT) SetHandlers(handlerIntance handlers.Handlers) {
	me.Handlers = handlerIntance
}

func (me *TiBOT) SetDefaultHandlerType(handlerType bot.HandlerType) {
	me.DefaultHandlerType = handlerType
}

func (me *TiBOT) GetDefaultHandlerType() bot.HandlerType {
	return me.DefaultHandlerType
}

func (me *TiBOT) SetRunPythonFunc(handlerFunc bot.HandlerFunc) {
	me.RunPythonFunc = handlerFunc
}

func (me *TiBOT) InitRules(b *bot.Bot) {

	for _, rule := range me.Rules {
		b.RegisterHandlerRegexp(rule.HandlerType, regexp.MustCompile(rule.Regex), rule.HandlerFunc)
	}

}

func (me *TiBOT) AddRule(handlerType bot.HandlerType, regex string, handlerFunc bot.HandlerFunc) {

	rule := Rule{
		HandlerType: handlerType,
		Regex:       regex,
		HandlerFunc: handlerFunc,
	}

	me.Rules = append(me.Rules, rule)
}

// func (me *TiBOT) AddPythonRule(regex string) {

// 	pyrule := Rule{
// 		HandlerType: me.DefaultHandlerType,
// 		Regex:       regex,
// 		HandlerFunc: me.RunPythonFunc,
// 	}

// 	me.PythonRules = append(me.PythonRules, pyrule)
// }

// func (me *TiBOT) InitPythonRules(b *bot.Bot) {

// 	pyconfig := me.Handlers.PyRunner.Config

// 	for _, pyscript := range pyconfig {
// 		me.AddPythonRule(pyscript.Handlers)
// 	}

// 	for _, pyrule := range me.PythonRules {
// 		b.RegisterHandlerRegexp(pyrule.HandlerType, regexp.MustCompile(pyrule.Regex), pyrule.HandlerFunc)
// 	}

// }
