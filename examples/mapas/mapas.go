// Los _Mapas_ son el tipo de [datos asociativo](http://es.wikipedia.org/wiki/Vector_asociativo) de Go.
// (también conocidos como Diccionarios o Hashes en otros lenguajes).

package main

import "fmt"

func main() {

    // Para crear un mapa vacío, se utiliza `make`:
    // `make(map[key-type]val-type)`.
    m := make(map[string]int)

    // Se pueden establecer los pares de clave/valor utilizando
    // la sintaxis típica `name[key] = val`.
    m["k1"] = 7
    m["k2"] = 13

    // Si se presenta el mapa con e.g. `Println` se muestran
    // todos sus pares de clave/valor.
    fmt.Println("map:", m)

    // Se obtiene el valor de una clave con la sintaxis `name[key]`.
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // La función `len` regresa el número de pares clave/valor cuando
    // se utiliza con un mapa.
    fmt.Println("len:", len(m))

    // La función `delete` elimina pares clave/valor de un mapa.
    delete(m, "k2")
    fmt.Println("map:", m)
    
    // El segundo valor de regreso (opcional) cuando obtienes un valor de
    // el mapa indica si la clave estaba presente. Este valor puede
    // ser usado para separar valores de claves que no existen y
    // valores de claves con valor cero, como `0` o `""`.
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // También puedes declarar e inicializar un mapa nuevo
    // en una única línea con la sintaxis:
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}
