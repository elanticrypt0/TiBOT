package pyrunner

type ScriptConfig struct {
	Handler   string `json:"handler"`
	Path      string `json:"script_path"`
	Engine    string `json:"engine"`
	OnlyAdmin bool   `json:"only_admin"`
}
