package main

import "fmt"
import "math/rand"

type CartaDaGioco struct {
	simbolo string
	seme string
}

type Mazzo struct {
	carte []CartaDaGioco
	ncarte int
}

//crea una carta da gioco: restituisce la carta da gioco dato il seme e il simbolo
func CreaCarta(simbolo, seme string) CartaDaGioco {
	return CartaDaGioco{simbolo, seme}
}

//crea e restituisce un mazzo di 40 carte da gioco
func CreaMazzo() Mazzo {
	var semi = [4]string{"quadri", "picche", "cuori", "fiori"}
	var simboli = [10]string{"asso", "due", "tre", "quattro", "cinque", "sei", "sette", "fante", "donna", "re"}
	tmp := make([]CartaDaGioco, 40)
	k := 0	
	for i:=0; i<4; i++ {
		for j:=0; j<10; j++{
			carta := CreaCarta(simboli[j], semi[i])
			tmp[k] = carta
			k++
		}
	}
	mazzo := Mazzo{carte:tmp, ncarte:40} 
	return mazzo
}


//mischia un mazzo di carte, usando il pacchetto "math/rand"
func (m Mazzo)Mischia(){
	for i := range m.carte {
		j := rand.Intn(i + 1)
		m.carte[i], m.carte[j] = m.carte[j], m.carte[i]
	}
}

//simula il prelievo di una carta dalla cima di un mazzo;
//restituisce la tripla (c,m,nil), dove c è la carta
//in cima al mazzo, m è un mazzo senza c, se il mazzo non è vuoto;
//restituisce (c,m,err), dove c ed m sono arbitrari, ed err è un
//valore diverso da nil, altrimenti
func (m Mazzo)Preleva() (CartaDaGioco, Mazzo, error) {
	if m.ncarte > 0 { 
		prelevata := m.carte[0]
		m.carte = m.carte[1:len(m.carte)]
		m.ncarte--
		return prelevata, m, nil
	} else {
		var empty_m Mazzo 
		return CartaDaGioco{"", ""}, empty_m, fmt.Errorf("Mazzo Terminato")
	}
}

//simula la posa di una carta in cima ad un mazzo
//restituisce la coppia (m,nil), dove m è un mazzo con l'aggiunta di
//c, se il mazzo non è pieno;
//restituisce (m,err), con m arbitrario ed err un valore diverso da
//nil, altrimenti
func (m Mazzo)Riponi(c CartaDaGioco) (Mazzo, error) {
	if m.ncarte < 40 {
		m.carte = append([]CartaDaGioco{c}, m.carte...)
		m.ncarte++
		return m, nil
	} else {
		return m, fmt.Errorf("Mazzo Pieno")
	}
}

func main() {
	
	mazzo := CreaMazzo()
	
	fmt.Println(mazzo)
	fmt.Println()
	
	mazzo.Mischia()
	fmt.Println(mazzo)

	carta, mazzo, _ := mazzo.Preleva()
	fmt.Println()
	fmt.Println(carta)
	fmt.Println(mazzo)

	mazzo, _ = mazzo.Riponi(carta)
	fmt.Println()
	fmt.Println(mazzo)

	mazzo, _ = mazzo.Riponi(carta)
	fmt.Println()
	fmt.Println(mazzo)
}
