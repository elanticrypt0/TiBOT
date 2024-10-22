# Example of use

```go
var baseURL = "http://localhost:9001/api/telbot"

func main(){
    argsQty, args := me.getArgs(update.Message.Text)
	
	response := ""

	if argsQty != 3 {
		response = "No hay suficientes argumentos"
	}

	operator := args[0]
	amount := args[1]
	description := args[2]

	method := "/sub"

	if operator == "+" {
		method += "/add"
	}

	var httpC = httpcaller.New(baseURL + method)

	httpC.Params.Add("amount", amount)
	httpC.Params.Add("description", description)
	httpC.UpdateQuery()

	response = fmt.Sprintf("Calling to %q", httpC.GetQuery())
}
```