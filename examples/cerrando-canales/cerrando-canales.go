// Al _cerrar_ un canal, indicamos que ya no se enviarán
// más valores por él. Esto puede ser útil para comunicar 
// a los canales receptores que el trabajo se ha completado.
package main

import "fmt"

// En este ejemplo vamos a usar el canal `jobs` para
// comunicar el trabajo que debe de ser hecho desde la  
// gorutina `main()` hacia la gorutina trabajadora. Cuando
// no haya más trabajos cerraremos el canal `jobs` con 
// la llamada built-in `close`.
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // Aquí esta la gorutina trabajadora. Recibe 
    // continuamente desde `jobs` con `j, more := <- jobs`.
    // En esta variante de recepción de 2 valores, el 
    // valor `more` será `false` si `jobs` ha sido cerrado
    // y todos los valores en este canal han sido recibidos.
    // Usamos esto para notificar en el canal `done` que 
    // ya hemos terminado con todos los trabajos.
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("trabajo recibido", j)
            } else {
                fmt.Println("todos los trabajos han sido"+
                            " recibidos")
                done <- true
                return
            }
        }
    }()

    // Aquí enviamos tres trabajos al trabajador por el canal
    // `jobs` y luego lo cerramos.
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("trabajo enviado", j)
    }
    close(jobs)
    fmt.Println("todos los trabajos han sido enviados")

    // Esperamos a que el trabajador termine usando 
    // la [sincronización](sincronizacion-de-canales) de
    // canales que vimos anteriormente
    <-done
}
