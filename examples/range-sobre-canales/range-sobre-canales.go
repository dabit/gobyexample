// También podemos usar esta sintáxis para iterar los
// valores recibidos desde un canal.
package main

import "fmt"

func main() {

    // Vamos a iterar sobre 2 valores en el canal `queue`.
    queue := make(chan string, 2)
    queue <- "uno"
    queue <- "dos"
    close(queue)

    // Este `range` itera sobre cada elemento conforme es
    // recibido desde `queue`. Debido a que llamamos a
    // `close` arriba, la iteración termina después de
    // recibir los 2 elementos. Si no lo cerramos
    // entonces bloquearíamos esperando por un tercer
    // elemento en la iteración.
    for elem := range queue {
        fmt.Println(elem)
    }
}
