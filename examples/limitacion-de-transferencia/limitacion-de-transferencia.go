// La limitación de tasa de transferencia es un mecanísmo
// importante para controlar la utilización de un
// recurso y mantener la calidad del servicio. Go lo
// soporta elegantemente usando [gorutinas](/goroutines), [canales](/canales) y
// [tickers](/tickers).

package main

import "time"
import "fmt"

func main() {

    // Primero veamos una limitación básica. Supongamos
    // que queremos limitar el número de peticiones
    // entrantes que podemos manejar. Serviremos estas
    // peticiones desde un canal con el mismo nombre.
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // Este canal `limiter` recibirá un valor cada 200
    // milisegundos. Este es el regulador en nuestro
    // esquema limitador de transferencia.
    limiter := time.Tick(time.Millisecond * 200)

    // Al bloquear durante la recepción del canal
    // `limiter` antes de servir cada petición, nos
    // autolimitamos a una petición cada 200 milisegundos
    for req := range requests {
        <-limiter
        fmt.Println("peticiones", req, time.Now())
    }

    // Podríamos permitir pequeños picos de peticiones
    // en nuestro esquema de limitación y seguir
    // conservando el límite general. Para lograrlo
    // podemos bufferear nuestro canal `limiter`. Este
    // canal `burstyLimiter` nos permitirá tener picos
    // de hasta 3 eventos.
    burstyLimiter := make(chan time.Time, 3)

    // Llenamos el canal para representar los picos.
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    // Cada 200 milisegundos intentaremos agregar un
    // nuevo valor a `burstyLimiter` hasta su límite.
    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t
        }
    }()

    // Ahora simularemos 5 peticiones más. La primera
    // de estas 3 se beneficiará de la capacidad de
    // soportar picos del canal `burstyLimiter`.
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("peticiones", req, time.Now())
    }
}
