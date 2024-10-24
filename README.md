# TiBOT

**Telegram interactive BOT**

# Installation

Install go from official site: [https://go.dev](https://go.dev)

Donwload clone the repo:

```bash
git clone https://github.com/elanticrypt0/TiBOT
```

# Configuration

modify the env-example file and add your BOT API TOKEN

## Scripts

add your scripts to the file **config/scripts.json** 

```json
    {
        "handler":"handler_name",
        "script_path":"scripts/script_file.py",
        "engine":"python",
        "only_admin":true
    }
```

You cand run bash, batch or vbs 

```json
    {
        "handler":"handler_name",
        "script_path":"scripts/script_file.*",
        "engine":"os",
        "only_admin":true
    }
```

Use the wildcard extension to execute the script in the default os script

## users.json

add the users to the bot

```json
    {
        "username":"guest",
        "telegram_id":TELEGRAM_ID,
        "is_admin":false
    }
```

# Run the bot

```bash
go run .
```

## build

```bash
go build -o tibot.exe
```

# BOT defaults handlers

- **/start**: 
- **/args**: 
- **/myinfo**: shows info about the user 
- **/botinfo**: shows info about the bot
- **/run help**: shows scripts availables
- **/run [SCRIPT_NAME] [ARGS]**: run the chosen script if the user with the args
- **/die**: 