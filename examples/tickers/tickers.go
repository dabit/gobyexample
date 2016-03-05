// Los [temporizadores](temporizadores) se usan cuando
// quieres hacer una cosa en el futuro - los _tickers_
// (por aquello de tic tac) - se usan cuando se quiere
// hacer algo repetidamente en intervalos regulares.
// Aquí hay un ejemplo de un ticker que hace tic
// periódicamente hasta que lo detenemos.

package main

import "time"
import "fmt"

func main() {

    // Los tickers usan un mecanismo similar a los
    // temporizadores: un canal al que se le envían
    // valores. Aquí vamos a usar el builtin `range` en
    // el canal para iterar cada 500ms los valores
    // conforme van llegando.
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tic a las", t)
        }
    }()

    // Los tickers pueden ser detenidos igual que los
    // temporizadores. Una vez que un ticker es detenido
    // ya no recibirá más valores en su canal.
    // Detendremos el nuestro después de 1500ms.
    time.Sleep(time.Millisecond * 1500)
    ticker.Stop()
    fmt.Println("Ticker detenido")
}
