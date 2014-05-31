// Los envíos y recepciones básicos sobre un canal lo bloquean.
// Sin embargo, podemos usar la cláusula `select` con un cláusula `default`
// para implementar envios y recepciones sin bloqueo (non-blocking) e incluso
// `select`s multi-vía sin bloqueo.
package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    // Aquí hay un envío sin bloqueo. Si hay algún valor
    // disponible en `messages` entonces el `select` tomará
    // el `case` `<-messages` con ese valor. Sino
    // tomará inmediatamente el `case` `default`.
    select {
    case msg := <-messages:
        fmt.Println("mensaje recibido", msg)
    default:
        fmt.Println("ningún mensaje recibido")
    }

    // Un envío sin bloqueo funciona de manera similar.
    msg := "hola"
    select {
    case messages <- msg:
        fmt.Println("mensaje envíado", msg)
    default:
        fmt.Println("ningún mensaje envíado")
    }

    // Podemos usar múltiples `case` encima del `default`
    // para implementar un `select` multi-via sin bloqueo.
    // Aqui intentamos una recepción sin bloqueo
    // `messages` y `signals`.
    select {
    case msg := <-messages:
        fmt.Println("mensaje recibido", msg)
    case sig := <-signals:
        fmt.Println("señal recibida", sig)
    default:
        fmt.Println("sin actividad")
    }
}
