# `zeroval` no cambia el valor de `i` en `main`,
# mientras que `zeroptr` sí ya que tiene una referencia 
# a la dirección en memoria de esa variable.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
