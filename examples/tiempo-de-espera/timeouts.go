// Los _Tiempos de Espera_ son una parte importante de los programas
// que se conectan a recursos externos o que
// necesitan limitar el tiempo de ejecución. La implementación
// de tiempos de espera en Go es fácil y elegante gracias
// a los canales y select.

package main

import "time"
import "fmt"

func main() {

    // Para este ejemplo, supongamos que estamos ejecutando
    // una llamada externa que regresa su resultado en el
    // canal `c1` después de 2s.
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    // Aquí tenemos el `select` implementando un tiempo de espera.
    // `res := <-c1` espera el resultado y `<-Time.After`
    // espera el valor que será enviado después del tiempo de
    // espera de 1s. Como `select` procede con el primer mensaje
    // recibido que esté listo, tomaremos el caso del tiempo de
    // espera si la operación toma más tiempo del permitido (1s).
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("tiempo de espera 1")
    }

    // Si permitimos un valor mayor a 3s, entonces c2 recibirá el
    // mensaje a tiempo y mostrará el resultado.
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "resultado 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("tiempo de espera 2")
    }
}

// todo: cancellation?
