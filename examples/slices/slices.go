// Los _Slices_ son un tipo de datos en Go que proporcionan
// una interfaz más poderosa a las secuencias que los arreglos.

package main

import "fmt"

func main() {

    // A comparación de los arreglos, los slices son solo del tipo
    // de los elementos que contienen (no del numero de elementos).
    // Para crear un slice de tamaño cero, se usa la sentencia `make`.
    // En este ejemplo creamos un slice de `string`s de tamaño `3`
    // (inicializado con valores cero).
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // Podemos establecer y obtener valores just como con los arreglos.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len` regresa el tamaño del slice.
    fmt.Println("len:", len(s))

    // Aparte de estas operaciones básicas, los slices
    // soportan muchas más que los hacen más funcionales
    // que los arreglos. Una de ellas es `append`, la que
    // regresa un slice que contiene uno o más valores nuevos.
    // Nota que necesitamos asignar el valor de regreso de
    // append tal como lo haríamos con el valor de un slice nuevo.
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // Los Slices pueden ser copiados utilizando `copy`.
    // Aquí creamos un slice vacío `c` del mismo tamaño que
    // `s` y copiamos el contenido de `s` a `c`.
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // Los Slices soportan un operador de rango con la sintaxis
    // `slice[low:high]`. Por ejemplo, esto regresa un slice
    // de los elementos `s[2]`, `s[3]`, y `s[4]`.
    l := s[2:5]
    fmt.Println("sl1:", l)

    // Esto regresa los elementos hasta antes de `s[5]`.
    l = s[:5]
    fmt.Println("sl2:", l)

    // y esto regresa los elementos desde `s[2]`.
    l = s[2:]
    fmt.Println("sl3:", l)

    // Podemos declarar e inicializar una variable para el slice
    // en una sola línea también.
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Los slices pueden ser compuestos de estructuras multi dimensionales.
    // A diferencia de los arreglos, el tamaño de los slices interiores
    // puede variar.
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
