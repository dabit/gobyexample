// Las URLs determinan una [manera uniforme de identificar recursos de manera
// única](http://adam.heroku.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// A continuación vemos como analizar URLs en Go.

package main

import "fmt"
import "net/url"
import "strings"

func main() {

	// Analicemos ésta URL de ejemplo, la cual incluye un esquema, información
	// de autenticación, servidor, puerto, ruta, parámetros de petición y un
	// fragmento.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Analizamos la URL y nos aseguramos de que no se generaron errores.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Obtener el esquema es muy fácil, sólo tenemos que leer la propiedad
	// `url.URL.Scheme`.
	fmt.Println(u.Scheme)

	// La propiedad [User](http://golang.org/pkg/net/url/#Userinfo) contiene
	// toda la información de autenticación dada al inicio.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// La propiedad `Host` contiene tanto el nombre de servidor como el puerto
	// (en caso de estar presentes). Es posible usar `strings.Split` a `Host`
	// manualmente para extraer el puerto.
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	// A continuación extraemos la ruta y el fragmento después de `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Para obtener los parámetos de la petición en formato `k=v` usamos
	// RawQuery. También es posible analizar los parámetros los parámetros en
	// un mapa. Los parámetros se mapean en un slice de strings, de tal manera
	// que podemos usar el índice [0] para obtener el primer valor.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
