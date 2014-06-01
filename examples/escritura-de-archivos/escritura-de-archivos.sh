# Ejecutamos el programa para demostrar la escritura de archivos.
$ go run writing-files.go
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# Verificamos el contenido de los archivos escritos.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# Ahora veremos cómo implementar algunas de las ideas de entrada y salida (I/O)
# con archivos, hasta ahora sólo hemos visto el uso de `stdin` y `stdout`.
