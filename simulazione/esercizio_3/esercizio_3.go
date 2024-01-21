package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

//utenze identificato da: numero e codice sim
// ad un numero possonon corrispondere + sim
// l'utenze di telefonia mobile attiva è quella con il codice più recente relativo alla sim

type Utenza struct {
	numero_telefono string
	codice_sim string
}

type RegistroTelefonico map[string][]string

func main() { //fine 2h 29min
	utenze := LeggiUtenze()
	registro := InizializzaRegistro()
	registro_riempito := riempi_registro(registro,utenze)
	// Println(registro_riempito)
	begin338(registro_riempito)
}

func begin338(registro RegistroTelefonico) {
	for numero_telefonico := range registro{
		if numero_telefonico[0:3] == "338"{
			Printf("Il numero %s è associato alla sim %s\n", numero_telefonico,Identifica(registro,numero_telefonico))
		}
	}
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {
	if len(registro[telefono])== 0{
		return ""
	}
	return registro[telefono][len(registro[telefono])-1]
}

func riempi_registro(registro RegistroTelefonico, utenze []Utenza) (registro_riempito RegistroTelefonico) {
	for i := 0; i < len(utenze); i++ {
		registro_riempito = AggiungiUtenza(registro,utenze[i])
	}
	return
}

func InizializzaRegistro() (registro RegistroTelefonico) {
	registro = make(RegistroTelefonico)
	return
}
//map[string][]string
func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) (registroAggiornato RegistroTelefonico) {
	
	registro[utenza.numero_telefono] = append(registro[utenza.numero_telefono],utenza.codice_sim)
	return registro
}

func LeggiUtenze() (utenze []Utenza) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line_split := strings.Split(scanner.Text(), ";")
		utenze = append(utenze, Utenza{line_split[0], line_split[1]})
	}
	return
}
