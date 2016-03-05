// Las _Funciones_ son una parte escencial de Go. Veamos algunos
// ejemplos para entender cómo se utilizan.

package main

import "fmt"

// Aquí tenemos una función que recibe dos valores `int`
// y regresa la suma de los mismos como `int`.
func plus(a int, b int) int {

    // Go requiere retornos de valor explícitos, i.e. no regresa
    // automáticamente el valor de la última expresión.
    return a + b
}

func main() {

    // Puedes llamar a una función tal y como lo esperarías, con
    // `funcion(args)`.
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
}
