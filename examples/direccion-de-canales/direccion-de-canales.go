// Cuando utilizas un canal como parámetro de una función,
// puedes especificar si el canal es sólo para recibir o enviar
// valores. Esto nos permite incrementar la seguridad de tipos
// del programa.

package main

import "fmt"

// Esta función `ping` sólo acepta un canal para enviar valores.
// Se arrojaría un error de compilación si intentamos recibir
// un valor en este canal.
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// La función `pong` acepta un canal para recibir (`pings`) y
// un segundo canal para enviar (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "mensaje enviado")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
