// Go soporta _múltiples valores de retorno_.
// Esta característica se usa mucho en Go idiomático, por
// ejemplo, para regresar un resultado y valores de error
// de una función.

package main

import "fmt"

// El `(int, int)` en la signatura de esta función indica que
// regresa 2 valores de tipo `int`.
func vals() (int, int) {
    return 3, 7
}

func main() {

    // Aquí asignamos los dos valores de retorno que devuelve
    // la función con _asignación múltiple_.
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // Si sólo quieres utilizar uno de los valores que regresa la
    // función, puedes utilizar el identificador vacío `_`.
    _, c := vals()
    fmt.Println(c)
}
