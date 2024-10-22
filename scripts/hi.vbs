' Verifica si se ha proporcionado un argumento
If WScript.Arguments.Count = 0 Then
    WScript.Echo "Por favor, proporciona un nombre."
Else
    nombre = WScript.Arguments(0)
    WScript.Echo "Hi, " & nombre
End If
