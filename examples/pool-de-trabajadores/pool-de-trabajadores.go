// En este ejemplo veremos cómo implementar un
// _pool de trabajadores_ usando goroutines y canales.

package main

import "fmt"
import "time"

// Este es el trabajador del cual correremos varias
// instancias concurrentes. Estos trabajadores recibirán
// trabajo desde el canal `jobs` y enviarán el resultado
// correspondiente en el canal `results`. Vamos a hacer
// una pausa de un segundo por trabajo para simular una
// tarea pesada.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("trabajador", id,
            "procesando trabajo", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {

    // Para usar nuestro pool de trabajadores necesitamos
    // enviarles trabajo y recolectar sus resultados.
    // Creamos dos canales para esto.
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Aquí iniciamos tres trabajadores, inicialmente
    // bloqueados porque no hay trabajos aún.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Aquí envíamos 9 _trabajos_ por el canal `jobs` y
    // luego lo cerramos para indicar que ya no hay más
    // trabajo por procesar.
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Finalmente recolectamos todos los resultados.
    for a := 1; a <= 9; a++ {
        <-results
    }
}
