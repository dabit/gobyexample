// Por defecto, los canales no usan un buffer, lo que
// signifca que sólo aceptan envíos (`chan <-`) si hay un
// receptor correspondiente (`<- chan`) listo para recibir
// el valor enviado. Los canales _con buffer_ pueden aceptar
// un número limitado de valores sin un receptor correspondiente
// para esos valores.

package main

import "fmt"

func main() {

    // Aquí creamos un canal de cadenas con buffer para hasta
    // dos valores.
    mensajes := make(chan string, 2)

    // Como este canal utiliza un buffer, podemos enviar estos
    // valores al mismo sin un receptor correspondiente.
    mensajes <- "buffered"
    mensajes <- "channel"

    // Más adelante, podemos recibir estos dos valores.
    fmt.Println(<-mensajes)
    fmt.Println(<-mensajes)
}
