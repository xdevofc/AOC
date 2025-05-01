package main

import (
	"sort"
)

func SmallPeimeter(L int, A int, H int) int {

	lista := []int{(2 * (L + A)), (2 * (A + H)), (2 * (L + H))}

	//ordenar de menor a mayor
	sort.Ints(lista)

	// obtener el perimetro mas chico
	minor := lista[0]
	volumen := L * A * H
	res := minor + volumen
	return res
}

func SecondChallenge(L int, A int, H int) int {
	return SmallPeimeter(L, A, H)
}
