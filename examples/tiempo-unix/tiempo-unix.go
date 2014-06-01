// Una necesidad de programación común es obtener el número de segundos que han
// transcurrido desde el [_UNIX
// Epoch_](http://es.wikipedia.org/wiki/Tiempo_Unix). A continuación veremos
// como hacerlo en Go.

package main

import "fmt"
import "time"

func main() {

	// Usamos `time.Now` (regresa un tipo `time.Time`) con `time.Time.Unix` o
	// `time.Time.UnixNano` para obtener el número de segundos, milisegundos o
	// nanosegundos respectivamente, transcurridos desde el [_Unix
	// Epoch_](http://es.wikipedia.org/wiki/Tiempo_Unix).
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Notemos que no existe una función `time.Time.UnixMillis`, pero podemos
	// obtener los milisegundos si dividimos manualmente los nanosegundos entre
	// 1000000.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// También podemos convertir un valor en segundos o nanosegundos desde el
	// [_UNIX Epoch_](http://es.wikipedia.org/wiki/Tiempo_Unix) en una valor de
	// fecha de Go (`time.Time`).
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
