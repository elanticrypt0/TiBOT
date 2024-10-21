# TODO

## Privilegios
[] Leer de un archivo json o csv los usuarios aceptados. Hacer que la app acepte a una lista de usuarios.
[] agregar la opción para que se recupere de un panic
[] Si el usuario NO está dentro de la lista entonces no hace nada.
[] si está habilitadad la regla allow_anyone entonces no chequea los usuarios, sólo al admin.

## Scripts
[x] La carga de los archivos python es dinámica. Entonces cada vez que se llama a la función carga el archivo y busca el comando.
[x] Carga el comando de ejecución de python en los handlers
[x] Carga la configuración de los comandos de python
[x] Ejecuta el comando de python y devuelve la salida al telegram bot
[x] Pasa los parámetros a la función de python
[x] Comprueba si el script es solo ejecutable por un admin y si el usuario es admin
