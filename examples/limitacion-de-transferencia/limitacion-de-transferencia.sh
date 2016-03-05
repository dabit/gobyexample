# Al ejecutar nuestro programa podemos ver que el primer
# bloque de peticiones es atendido cada ~200 milisegundos
# apróximadamente como lo deseamos.
$ go run limitacion-de-transferencia.go
peticiones 1 2014-07-16 17:58:36.733961487 +0000 UTC
peticiones 2 2014-07-16 17:58:36.933979229 +0000 UTC
peticiones 3 2014-07-16 17:58:37.133983308 +0000 UTC
peticiones 4 2014-07-16 17:58:37.333995394 +0000 UTC
peticiones 5 2014-07-16 17:58:37.534003928 +0000 UTC

# Para el segundo bloque de peticiones, servimos los
# primeros 3 inmediatamente usando el soporte de picos,
# y luego servimos los 2 restantes con un retraso de 
#  ~200 milisegundos
peticiones 1 2014-07-16 17:58:37.534072367 +0000 UTC
peticiones 2 2014-07-16 17:58:37.534085589 +0000 UTC
peticiones 3 2014-07-16 17:58:37.534094082 +0000 UTC
peticiones 4 2014-07-16 17:58:37.734343287 +0000 UTC
peticiones 5 2014-07-16 17:58:37.935017031 +0000 UTC
