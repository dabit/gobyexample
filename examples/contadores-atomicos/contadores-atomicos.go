// El mecanismo principal para manejar estados en Go es la comunicación
// mediante canales (channels).  Podemos ver esto, por ejemplo, con los
// [_pools de trabajadores_](/pool-de-trabajadores). Además de canales existen otras formas para
// manejar estados. En éste código conoceremos el uso del paquete
// `sync/atomic` para _contadores atómicos_, los cuales pueden ser accedidos
// por múltiples _goroutines_.

package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

    // Usaremos un entero sin signo para representar a un contador (que siempre
    // será positivo).
    var ops uint64 = 0

    // Para simular actualizaciones concurrentes iniciaremos 50 _goroutines_,
    // cada una de las cuales incrementará el contador, aproximadamente una vez
    // cada milisegundo.
    for i := 0; i < 50; i++ {
        go func() {
            for {
                // Para incrementar el contador de manera
                // [atómica](http://es.wikipedia.org/wiki/Atomicidad) usamos la
                // función `atomic.AddUint64`, pasando la dirección en memoria
                // del contador `ops` usando la sintaxis `&`.
                atomic.AddUint64(&ops, 1)

                // Permite continuar a otros _goroutines_.
                runtime.Gosched()
            }
        }()
    }

    // Espera un segundo. Esto permite que se acumulen algunos valores en
    // `ops`.
    time.Sleep(time.Second)

    // Para usar de manera segura el contador mientras está siendo actualizado
    // por otras _goroutines_, hacemos una copia del valor actual en la
    // variable `opsFinal`, usando la función `atomic.LoadUint64`. Tal como en
    // el bloque anterior necesitamos pasar la dirección de la variable `ops`,
    // es decir, `&ops`, de donde se copiará el valor.
    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}
