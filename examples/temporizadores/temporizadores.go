// Con frequencia queremos ejecutar código Go en algún
// momento en el futuro, o repetidamente dentro de un
// intervalo. Los _timer_ y _ticker_ nativos de Go hacen
// fáciles ambas tareas. Veremos primero los temporizadores y
// luego los tickers.

package main

import "time"
import "fmt"

func main() {

    // Los temporizadores representan un evento único en
    // el futuro. Tú le dices al temporizador cuánto tiempo
    // quieres esperar, y él te proporcionará un canal que
    // será notificado en ese momento. Este temporizador
    // esperará 2 segundos.
    timer1 := time.NewTimer(time.Second * 2)

    // `<-timer1.C` bloquea el canal `C` del temporizador
    // hasta que envía un valor indicando que el
    // temporizador ha finalizado
    <-timer1.C
    fmt.Println("Temporizador 1 finalizado")

    // Si solo necesitas esperar, puedes usar
    // `time.Sleep`. Una de las razones por las que un
    // temporizador puede ser útil es que puedes
    // detenerlo antes de que finalice.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Temporizador 2 finalizado")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Temporizador 2 detenido")
    }
}
