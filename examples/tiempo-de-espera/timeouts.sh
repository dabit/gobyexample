# Al ejecutar este programa podemos ver que en
# la primera operación el tiempo de espera
# expira y que el segundo caso es exitoso.
$ go run tiempo-de-espera.go
tiempo de espera 1
resultado 2

# Utilizar este patrón de `select` requiere comunicar
# resultados a través de canales. Esta es una buena
# idea en general porque otras características importantes
# de Go están basadas en canales y `select`. Vamos a ver
# dos ejemplos a continuación: timers y tickers.
