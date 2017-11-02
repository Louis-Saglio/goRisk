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

	fmt.Println(nbrAttaquants, nbrDefenseurs)
	return "nil", 0, 0
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
	combattre(190, 720)
	fmt.Println(time.Now().UTC().UnixNano() / 1000000 - debut, "millisecondes")
}
