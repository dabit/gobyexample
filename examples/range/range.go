// La función _range_ itera por los elementos de la mayoría
// de las estructuras de datos. Veamos como podríamos usar
// `range` con algunas de las que ya conocemos.

package main

import "fmt"

func main() {

    // Aquí utilizamos `range` para sumar los números de un slice.
    // En los arreglos funciona de manera similar.
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // Cuando se usa `range` en un arreglo o slice, se obtiene
    // el valor del elemento y el índice donde se encuentra.
    // En el ejemplo anterior no utilizamos el índice, así que
    // lo ignoramos asignándolo a un _identificador vacío_, `_`.
    // Sin embargo, en algunas ocasiones sí vamos a necesitar los índices.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // Al utilizar `range` en un mapa, este itera por
    // los pares clave/valor.
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // Si se utiliza `range` en una cadena, se itera por los
    // caracteres Unicode de la misma. El primer valor es
    // el índice del byte inicial del primer símbolo y el
    // símbolo mismo.
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
