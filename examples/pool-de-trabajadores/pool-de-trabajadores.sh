# La salida de nuestro programa muestra que 9 trabajos
# son ejecutados por varios trabajadores. El programa
# solo necesita cerca de 3 segundos para realizar 
# cerca de 9 segundos de trabajo total, ya que
# hay 3 trabajadores operando de manera concurrente.
$ time go run examples/worker-pools/worker-pools.go  
trabajador 1 procesando trabajo 1
trabajador 3 procesando trabajo 3
trabajador 2 procesando trabajo 2
trabajador 1 procesando trabajo 4
trabajador 3 procesando trabajo 5
trabajador 2 procesando trabajo 6
trabajador 1 procesando trabajo 7
trabajador 3 procesando trabajo 8
trabajador 2 procesando trabajo 9

real	0m3.380s
