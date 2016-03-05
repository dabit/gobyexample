# Cuando corremos el programa, el mensaje `"ping"` se
# pasa de una goroutine a otra a través de nuestro
# canal.
$ go run channels.go
ping

# Por defecto la recepción y los envíos se bloquean hasta
# que tanto receptor como transmisor están listos. Esta
# propiedad nos permite esperar el mensaje
# `"ping"` hasta el final del programa sin tener
# que usar ningún otro tipo de sincronización.
