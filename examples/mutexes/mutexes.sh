# Al ejecutar el programa vemos que se ejecutaron cerca
# de 3,500,000 operaciones a nuestro `estado` sincronizado
# con `mutex` 
$ go run mutexes.go
ops: 3598302
state: map[1:38 4:98 2:23 3:85 0:44]

# A continuación veremos cómo implementar el mismo 
# manejo de estado usando solamente gorutinas y
# canales.
