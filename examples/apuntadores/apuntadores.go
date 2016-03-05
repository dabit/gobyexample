// Go soporta <em><a href="http://es.wikipedia.org/wiki/Puntero_(inform%C3%A1tica)">apuntadores</a></em>,
// lo cual nos permite pasar referencias de valores y datos entre
// las funciones de nuestro programa.

package main

import "fmt"

// Veamos cómo los apuntadores funcionan en contraste con los valores directos
// utilizando dos funciones: `zeroval` y `zeroptr`. `zeroval` recibe
// un parametro `int`, así que los argumentos serán pasados por valor.
// `zeroval` va a recibir una copia de `ival` distinta a la de la función
// donde se llama.
func zeroval(ival int) {
    ival = 0
}

// `zeroptr` recibe un parámetro `*int`, lo que significa
// que recibe un apuntador a un valor `int`. El código `*iptr` en el
// cuerpo de la función _dereferencía_ el apuntador de su dirección de
// memoria al valor actual de esa dirección.
// Si asignamos un valor a un apuntador dereferenciado se cambia el valor
// que se está almacenando en dicha dirección de memoria.
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // La sintaxis `&i` devuelve la dirección en memoria de `i`,
    // i.e. un apuntador a `i`.
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // Los apuntadores también pueden mostrarse en pantalla
    ft.Println("pointer:", &i)
}
