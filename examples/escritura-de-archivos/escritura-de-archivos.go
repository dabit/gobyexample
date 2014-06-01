// La escritura de archivos en Go sigue patrones similares a los que vimos
// anteriormente para lectura.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Aquí vemos como guardar una cadena de bytes ([]bytes) a un archivo.
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Para escribir de manera más controladas podemos intentar abrir un
	// archivo en modo de escritura.
	f, err := os.Create("/tmp/dat2")
	check(err)

	// Es idiomático postergar un `os.File.Close` inmediatamente después de
	// abrir un archivo, para esto usamos `defer`.
	defer f.Close()

	// Es posible escribir una secuencia de `byte` (`[]byte`) usando
	// `os.File.Write`.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// La función `os.File.WriteString` permite escribir usando tipos `string`
	// en vez de `[]byte`.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Usamos `os.File.Sync` para asegurarnos que las escrituras solicitadas
	// han sido ejecutadas.
	f.Sync()

	// El paquete `bufio` provee un búfer para escritura, muy útil en conjunto
	// con los búfers de lectura que conocimos anteriormente.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Finalmente, usamos `bufio.Write.Flush` para asegurarnos que todas las
	// operaciones en búfer han sido aplicadas al escritor correspondiente.
	w.Flush()

}
