package pyrunner

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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

func (me *PyRunner) getScriptsByEngine(engine string) ([]ScriptConfig, error) {

	scriptsAvailables := []ScriptConfig{}

	for _, scriptAux := range me.Config {
		if scriptAux.Engine == engine {
			scriptsAvailables = append(scriptsAvailables, scriptAux)
		}
	}

	return scriptsAvailables, nil

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

func (me *PyRunner) RunScript(engine string, params []string) (string, error) {

	log.Printf("	params > (%q) %q  %q \n", engine, params[0], params[1:])

	response, err := me.cmdRun(engine, params)
	if err != nil {
		return "", err
	}

	return response, nil
}

func (cr *PyRunner) cmdRun(engine string, args []string) (string, error) {

	var outb, errb bytes.Buffer
	var cmd *exec.Cmd

	// if engine is not python run this otherwise run OS
	if engine != "python" {
		commandOutput, err := cr.runCommandByOS(args[0], args[1:])
		if err != nil {
			return "", err
		}
		return commandOutput, nil
	}

	pythonCommand := os.Getenv("PYTHON_COMMAND")
	cmd = exec.Command(pythonCommand, args...)
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

// EJECUTA otros scripts

// runCommandByOS detecta el sistema operativo y obtiene el comando a ejecutar
func (cr *PyRunner) runCommandByOS(scriptPath string, args []string) (string, error) {
	osType := runtime.GOOS

	var cmd *exec.Cmd

	isScriptPathExtWildcard := filepath.Ext(scriptPath) == ".*"
	scriptPath2run := ""
	argsStr := strings.Join(args, " ")

	switch osType {
	case "linux", "darwin": // Linux o macOS
		fmt.Println("Ejecutando script en bash (Linux/macOS)")
		if isScriptPathExtWildcard {
			scriptPath2run = cr.replaceFileExtension(scriptPath, ".sh")
		}
		cmd = exec.Command("bash", "-c", fmt.Sprintf("%s %s", scriptPath2run, argsStr))
	case "windows": // Windows
		// debe buscar los archivos que existan con ese nombre
		if isScriptPathExtWildcard {
			scriptsAvailables, err := cr.listFilesWithDifferentExtensions(scriptPath)
			if err != nil {
				return "", err
			}
			// por defecto siempre devuelve el primer que encuentra
			scriptPath2run = scriptsAvailables[0]
		}
		// ejecuta dependiendo del tipo de archivo
		if cr.hasBatchExtension(scriptPath2run) {
			fmt.Println("Ejecutando script en Batch (Windows)")
			cmd = exec.Command("cmd", "/C", fmt.Sprintf("%s %s", scriptPath2run, argsStr))
		} else if cr.hasVBSExtension(scriptPath2run) {
			fmt.Println("Ejecutando script en VBS (Windows)")
			cmd = exec.Command("wscript", fmt.Sprintf("%s %s", scriptPath2run, argsStr))
		} else {
			return "", fmt.Errorf("Tipo de script no soportado para Windows: %s", scriptPath)
		}
	default:
		return "", fmt.Errorf("Sistema operativo no soportado: %s", osType)
	}

	// Ejecutar el comando
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Error al ejecutar el script: %v\nSalida: %s", err, string(output))
	}

	return string(output), nil

}

// hasBatchExtension verifica si el archivo tiene extensión .bat o .cmd
func (cr *PyRunner) hasBatchExtension(path string) bool {
	return len(path) > 4 && (path[len(path)-4:] == ".bat" || path[len(path)-4:] == ".cmd")
}

// hasVBSExtension verifica si el archivo tiene extensión .vbs
func (cr *PyRunner) hasVBSExtension(path string) bool {
	return len(path) > 4 && path[len(path)-4:] == ".vbs"
}

// replaceFileExtension reemplaza la extensión de un archivo con una nueva
func (cr *PyRunner) replaceFileExtension(filePath, newExtension string) string {
	// Obtener el directorio y el nombre base del archivo sin la extensión
	baseName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

	// Combinar el directorio con el nuevo nombre y la nueva extensión
	return filepath.Join(filepath.Dir(filePath), baseName+newExtension)
}

func (cr *PyRunner) listFilesWithDifferentExtensions(pathPattern string) ([]string, error) {
	// Utiliza filepath.Glob para encontrar archivos que coincidan con el patrón
	matches, err := filepath.Glob(pathPattern)
	if err != nil {
		return nil, fmt.Errorf("error al buscar archivos: %v", err)
	}

	return matches, nil
}