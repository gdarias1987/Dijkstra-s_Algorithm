package main

import (
	"fmt"
)

const w = 9
const h = 9
const points = w * h

// Point - es donde se guardaran las distancias
type Point struct {
	x    int
	y    int
	dist int
	vis  bool
}

// Infinity - Variable con valor "inifito"
const Infinity = int(^uint(0) >> 1)

func getPointsArray() [points]Point {
	var listado [points]Point
	count := 0

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			listado[count] = Point{i, j, Infinity, false}
			count++
		}
	}

	listado[0].vis = true
	listado[0].dist = 0

	return listado
}

func getNextPoint(listado *[points]Point) Point {

	min := Infinity
	pos := 0

	for i := 0; i < points; i++ {
		if !listado[i].vis && listado[i].dist < min {
			min = listado[i].dist
			pos = i
		}
	}

	if pos != 0 {
		listado[pos].vis = true
		return listado[pos]
	}

	return Point{0, 0, 0, false}
}

func setPointDist(x int, y int, dist int, L *[points]Point) {
	for i := 0; i < points; i++ {
		if L[i].x == x && L[i].y == y {
			L[i].dist = dist
		}
	}
}

func getPointDist(y int, x int, L *[points]Point) int {
	for i := 0; i < points; i++ {
		if L[i].x == x && L[i].y == y {
			return L[i].dist
		}
	}
	return 0
}

func printList(L *[points]Point) {
	for i := 0; i < points; i++ {
		fmt.Println(L[i])
	}
}

func main() {

	// var mInicial = [w][h]int{
	// 	{0, 2, 3, 4},
	// 	{133, 12, 1, 44},
	// 	{144, 1, 135, 24},
	// 	{1, 2, 3, 4},
	// }

	var mInicial = [w][h]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0},
	}

	var lVis = getPointsArray()

	var init Point = lVis[0]

	for ok := true; ok; ok = init.vis == true {

		X := init.x
		Y := init.y

		//Right
		if X < w-1 {
			var control = mInicial[Y][X+1] + init.dist
			var actual = getPointDist(Y, X+1, &lVis)
			if control < actual {
				setPointDist(X+1, Y, control, &lVis)
			}
		}

		//Down
		if Y < (h - 1) {
			var control = mInicial[Y+1][X] + init.dist
			var actual = getPointDist(Y+1, X, &lVis)
			if control < actual {
				setPointDist(X, Y+1, control, &lVis)
			}
		}

		//Down-Right
		if X < w-1 && Y < h-1 {
			var control = mInicial[Y+1][X+1] + init.dist
			var actual = getPointDist(Y+1, X+1, &lVis)
			if control < actual {
				setPointDist(X+1, Y+1, control, &lVis)
			}
		}

		//Down-Left
		if X >= 1 && Y < h-1 {
			var control = mInicial[Y+1][X-1] + init.dist
			var actual = getPointDist(Y+1, X-1, &lVis)
			if control < actual {
				setPointDist(X-1, Y+1, control, &lVis)
			}
		}

		//Left
		if X >= 1 {
			var control = mInicial[Y][X-1] + init.dist
			var actual = getPointDist(Y, X-1, &lVis)
			if control < actual {
				setPointDist(X-1, Y, control, &lVis)
			}
		}

		//Up
		if Y >= 1 {
			var control = mInicial[Y-1][X] + init.dist
			var actual = getPointDist(Y-1, X, &lVis)
			if control < actual {
				setPointDist(X, Y-1, control, &lVis)
			}
		}

		//Up-Left
		if X >= 1 && Y >= 1 {
			var control = mInicial[Y-1][X-1] + init.dist
			var actual = getPointDist(Y-1, X-1, &lVis)
			if control < actual {
				setPointDist(X-1, Y-1, control, &lVis)
			}
		}

		//Up-Right
		if X < w-1 && Y >= 1 {
			var control = mInicial[Y-1][X+1] + init.dist
			var actual = getPointDist(Y-1, X+1, &lVis)
			if control < actual {
				setPointDist(X+1, Y-1, control, &lVis)
			}
		}

		init = getNextPoint(&lVis)
	}

	// Impresion de listado completo, con todas las distancias
	printList(&lVis)

	//fmt.Printf("La distancia mas corta desde la fuente (0,0)\nal extremo opuesto (%d,%d), es %d.", h, w, lVis[points-1].dist)
}
