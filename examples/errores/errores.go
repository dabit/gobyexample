// En Go es idiomático comunicar errores a través de
// un valor de retorno separado.
// Esto contrasta con las excepciones usadas en lenguajes
// como Java y Ruby y sobrecargar el resultado como en
// ocasiones se hace en C.
// Con la manera en que se hace en Go es más facil ver
// qué funciones regresan errores y manejarlos utilizando
// las mismas estructuras de control como lo hacemos con
// todas las demás tareas.

package main

import "errors"
import "fmt"

// Por convención, los errores siempre son el último valor
// que se regresa y son de tipo `error`, una interfaz que
// es parte del lenguaje.
func f1(arg int) (int, error) {
    if arg == 42 {

        // `errors.New` construye un valor de `error` básico
        // con el mensaje proporcionado.
        return -1, errors.New("can't work with 42")

    }

    // Un valor nulo en la posición de error indica que no hubo
    // ningún problema.
    return arg + 3, nil
}

// Es posible usar tipos personalizados como `error` simplemente
// implementando el método `Error()` en ellos. Aquí una
// variante del ejemplo anterior que utiliza un tipo personalizado
// para representar explícitamente un error de argumentos.
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {

        // En este caso usamos la sintaxis `&argError` para
        // construir un struct nuevo, proporcionando los valores
        // de los dos campos `arg` y `prob`.
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    // Estos dos ciclos prueban cada una de nuestras
    // funciones. Nota que el uso de la revisión de errores
    // en una única línea de `if` es un estilo común en Go
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    // Si quieres usar la información que es parte de un error personalizado
    // programáticamente, vas a necesitar asignar el error a una
    // una instancia del tipo personalizado de error por medio de la
    // aserción de tipo.
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
