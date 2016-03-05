// Go soporta
// <a href="http://es.wikipedia.org/wiki/Recursión_(ciencias_de_computación)"><em>funciones recursivas</em></a>.
// Aquí tenemos un ejemplo clásico para calcular un factorial.

package main

// Esta función `fact` se llama a sí misma hasta que llega a la
// base de `fact(0)`.
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    ft.Println(fact(7))
}
