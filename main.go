package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)



func combattre(nbrAttaquants int, nbrDefenseurs int) (string, int, int) {

	for nbrAttaquants > 1 && nbrDefenseurs > 0 {

		desAttaquant := make([]int, 0, 3)
		desDefenseur := make([]int, 0, 3)

		nbrDesAttaquant := getNbrDesAttaquant(nbrAttaquants)
		for i := 0; i < nbrDesAttaquant; i++ {
			rand.Seed(int64(time.Now().UTC().UnixNano()))
			desAttaquant = append(desAttaquant, rand.Intn(5)+1)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(desAttaquant)))

		nbrDesDefenseur := getNbrDesDefenseur(desAttaquant, nbrDefenseurs)
		for i := 0; i < nbrDesDefenseur; i++ {
			desDefenseur = append(desDefenseur, rand.Intn(5)+1)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(desDefenseur)))

		for i := 0; i < int(math.Min(float64(len(desAttaquant)), float64(len(desDefenseur)))); i++ {

			if desAttaquant[i] > desDefenseur[i] {
				nbrDefenseurs -= 1
			} else {
				nbrAttaquants -= 1
			}
		}
	}

	vainqueur := "defenseur"
	if nbrDefenseurs == 0 {
		vainqueur = "attaquant"
	}

	return vainqueur, nbrAttaquants, nbrDefenseurs
}


func resultatMoyen(nbrAttaquants int, nbrDefenseurs int, nbrSimulations int) (string, float64, float64, float64, float64){
	totalVictoiresAttaquant, totalVictoiresDefenseur, totalResteAttaquant, totalResteDefenseur := 0, 0, 0, 0
	for i := 0; i < nbrSimulations; i++ {
		vainqueur, resteAttaquant, resteDefenseur := combattre(nbrAttaquants, nbrDefenseurs)
		if vainqueur == "attaquant" {
			totalVictoiresAttaquant += 1
		} else {
			totalVictoiresDefenseur += 1
		}
		totalResteAttaquant += resteAttaquant
		totalResteDefenseur += resteDefenseur
	}
	probaVictoireAttaquant := float64(totalVictoiresAttaquant) / float64(nbrSimulations)
	probaVictoireDefenseur := float64(totalVictoiresDefenseur) / float64(nbrSimulations)
	vainqueur := "defenseur"
	if probaVictoireAttaquant > probaVictoireDefenseur {
		vainqueur = "attaquant"
	}
	return vainqueur, float64(totalResteAttaquant) / float64(nbrSimulations), float64(totalResteDefenseur) / float64(nbrSimulations), probaVictoireAttaquant, probaVictoireDefenseur
}


func trouverCombienGagnentContre(nbrDefenseurs int) (int, float64) {
	nbrAttaquants := 1
	vainqueur := "defenseur"
	var probabiliteVictoireAttaquant float64
	for vainqueur == "defenseur" {
		nbrAttaquants += 1
		vainqueur, _, _, probabiliteVictoireAttaquant, _ = resultatMoyen(nbrAttaquants, nbrDefenseurs, 100)
	}
	return nbrAttaquants, probabiliteVictoireAttaquant
}

func getNbrDesDefenseur(desAttaquant []int, nbrDefenseurs int) int {
	deuxiemeDe := 0
	if len(desAttaquant) > 1 {
		deuxiemeDe = desAttaquant[1]
	}
	if nbrDefenseurs < 2 || desAttaquant[0]+deuxiemeDe > 7 {
		return 1
	}
	return 2
}


func getNbrDesAttaquant(nbrAttaquants int) int {
	return int(math.Trunc(float64(nbrAttaquants-1)))
}


func main() {
	debut := time.Now().UTC().UnixNano() / 1000000

	for i := 0; i < 100; i++ {
		nbrAttaquants, probabilite := trouverCombienGagnentContre(i)
		fmt.Println(nbrAttaquants, "gagnent contre", i, "probabilité :", probabilite)
	}

	fmt.Println(time.Now().UTC().UnixNano() / 1000000 - debut, "millisecondes")
}
