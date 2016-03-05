// Los _structs_ en go son colecciones de campos con tipos específicos.
// Son útiles para agrupar datos y formar registros.

package main

import "fmt"

// El struct `persona` tiene los campos `nombre` y `edad`.
type persona struct {
    nombre string
    edad   int
}

func main() {

    // Esta sintaxis crea una instancia de un struct.
    fmt.Println(persona{"Bob", 20})

    // Puedes definir los valores de los campos cuando inicializas un struct.
    fmt.Println(persona{name: "Alice", age: 30})

    // Los campos omitidos serán de valor cero.
    fmt.Println(persona{name: "Fred"})

    // El prefijo `&` devuelve el apuntador a un struct.
    fmt.Println(&persona{name: "Ann", age: 40})

    // Puedes acceder a los campos del struct con un punto.
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // También puedes usar puntos con apuntadores a struct -
    // los apuntadores son automáticamente dereferenciados.
    sp := &s
    fmt.Println(sp.age)

    // Los Structs son inmutables.
    sp.age = 51
    fmt.Println(sp.age)
}
