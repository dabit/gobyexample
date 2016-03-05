// Las _Interfaces_ son colecciones nombradas
// de signaturas de métodos.

package main

import "fmt"
import "math"

// Aquí podemos ver la interfaz básica de una
// figura geométrica.
type geometrica interface {
    area() float64
    perim() float64
}

// Para nuestro ejemplo vamos a implementar esta
// interfaz en los tipos `cuadro` y `circulo`.
type cuadro struct {
    ancho, altura float64
}
type circulo struct {
    radio float64
}

// Para implementar una interfaz en Go, sólo tenemos
// que implementar todos los métodos de la misma.
// Aquí implementamos `geometrica` en `cuadro`.
func (s cuadro) area() float64 {
    return s.ancho * s.altura
}
func (s cuadro) perim() float64 {
    return 2*s.ancho + 2*s.altura
}

// La implementación para `circulo`
func (c circulo) area() float64 {
    return math.Pi * c.radio * c.radio
}
func (c circulo) perim() float64 {
    return 2 * math.Pi * c.radio
}

// Si una variable es de tipo de una interfaz podemos
// llamar directamente los metodos definidos en dicha interfaz. Aquí tenemos
// un ejemplo genérico de la función `medida` que aprovecha
// esto para poder actuar sobre cualquier `geometrica`.
func medida(g geometrica) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    s := cuadro{ancho: 3, altura: 4}
    c := circulo{radio: 5}

    // Los tipos `circulo` y `cuadro` implementan la
    // interfaz `geometrica` así que podemos usar instancias de ambas
    // como parámetro de `medida`.
    medida(s)
    medida(c)
}
