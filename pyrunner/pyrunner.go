package pyrunner

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
)

type PyRunner struct {
	ConfigFilePath string
	Config         []ScriptConfig
	UserIsAdmin    bool
}

func New() PyRunner {
	pyrun := PyRunner{
		ConfigFilePath: "./config/scripts.json",
	}
	pyrun.LoadConfig()
	return pyrun
}

func (me *PyRunner) LoadConfig() {
	config := &[]ScriptConfig{}
	ReadAndParseJson(me.ConfigFilePath, config)
	me.Config = *config
}

func (me *PyRunner) GetScript(handler string) (ScriptConfig, error) {

	script, err := me.getScriptByHandler(handler)
	if err != nil {
		return ScriptConfig{}, err
	}

	return script, nil

}

func (me *PyRunner) getScriptByHandler(handler string) (ScriptConfig, error) {
	var script ScriptConfig

	for _, scriptAux := range me.Config {
		if scriptAux.Handler == handler {
			script = scriptAux
			break
		}
	}

	if script.Path == "" {
		log.Printf(" > %v\n", ErrorScriptNotFound)
		return ScriptConfig{}, errors.New(ErrorScriptNotFound)
	}

	log.Printf(" > selected script	=> %q\n", script.Path)

	return script, nil
}

func (me *PyRunner) RunScript(params []string) (string, error) {

	log.Printf("	params > %q \n", params[1:])
	response, err := me.cmdRun(params)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (cr *PyRunner) cmdRun(args []string) (string, error) {
	pyCmd := os.Getenv("PYTHON_COMMAND")

	var outb, errb bytes.Buffer

	cmd := exec.Command(pyCmd, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	log.Printf("> %s\n", cmd.String())
	// return cmd.Run()

	err := cmd.Run()
	if err != nil {
		// log.Fatal(err.Error())
		return "", errors.New(errb.String())
	}

	return outb.String(), nil

}

func (me *PyRunner) GetScriptsList() ([]string, error) {
	var scripts []string

	for _, script := range me.Config {
		scripts = append(scripts, script.Handler)
	}

	log.Printf(">> Python Scripts: %v", scripts)

	return scripts, nil
}
