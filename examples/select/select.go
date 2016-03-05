// La función _select_ de Go te permite esperar los resltados de las operaciones
// de varios canales.
// Combinar goroutines y canales con select
// es una de las características más poderosas de Go.

package main

import "time"
import "fmt"

func main() {

    // Para nuestro ejemplo vamos a utilizar select con dos canales
    c1 := make(chan string)
    c2 := make(chan string)

    // Cada canal va a recibir un valor después de cierto tiempo,
    // para simular operaciónes bloqueadas de procesos externos
    // en goroutines concurrentes.
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "uno"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "dos"
    }()

    // Vamos a usar `select` para esperar ambos valores
    // simultaneamente, y mostraremos cada uno conforme llegue.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("recibido", msg1)
        case msg2 := <-c2:
            fmt.Println("recibido", msg2)
        }
    }
}
