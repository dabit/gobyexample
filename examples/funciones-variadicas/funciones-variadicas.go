// Las _Funciones Variádicas_ pueden ser llamadas con
// cualquier numero de argumentos. Un ejemplo común de
// una función variádica es `fmt.Println`.

package main

import "fmt"

// Aquí declaramos una función que va a recibir un número
// arbitrario de valores tipo `int` como argumentos.
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // Las funciones variádicas pueden ser llamadas de la misma
    // manera con valores individuales.
    sum(1, 2)
    sum(1, 2, 3)

    // También puedes tener los argumentos en un slice
    // y aplicarlos a la función variádica utilizando la forma
    // `func(slice...)`:
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
