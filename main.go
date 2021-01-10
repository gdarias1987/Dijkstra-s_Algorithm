package main

import (
	"fmt"
	//	"math"
	//	"net/http"
	//	"strconv"
)

const w = 4
const h = 4
const points = w * h

type Point struct {
	x    int
	y    int
	dist int
	vis  bool
}

const Infinity = int(^uint(0) >> 1)

func getPathMatrix() [w][h]int {
	var mAux [w][h]int

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			mAux[i][j] = Infinity
		}
	}
	mAux[0][0] = 0
	return mAux
}

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

func printMPath(P *[h][w]int) {
	for i := 0; i < h; i++ {
		fmt.Println(P[i])
	}
}

func main() {

	var mInicial = [w][h]int{
		{0, 2, 3, 4},
		{133, 12, 1, 44},
		{144, 1, 135, 24},
		{1, 2, 3, 4},
	}

	// var mPath = getPathMatrix()
	var lVis = getPointsArray()

	// var init Point = Point{0, 0, 0, true}
	var init Point = lVis[0]

	for ok := true; ok; ok = init.vis == true {

		X := init.x
		Y := init.y

		//Right
		if X < w-1 {
			var control = mInicial[Y][X+1] + init.dist
			var actual = getPointDist(Y, X+1, &lVis) //mPath[Y][X+1]
			if control < actual {
				//mPath[Y][X+1] = control
				setPointDist(X+1, Y, control, &lVis)
			}
		}

		//Down
		if Y < (h - 1) {
			var control = mInicial[Y+1][X] + init.dist
			var actual = getPointDist(Y+1, X, &lVis) //mPath[Y+1][X]
			if control < actual {
				//mPath[Y+1][X] = control
				setPointDist(X, Y+1, control, &lVis)
			}
		}

		//Down-Right
		if X < w-1 && Y < h-1 {
			var control = mInicial[Y+1][X+1] + init.dist
			var actual = getPointDist(Y+1, X+1, &lVis) // mPath[Y+1][X+1]
			if control < actual {
				//mPath[Y+1][X+1] = control
				setPointDist(X+1, Y+1, control, &lVis)
			}
		}

		//Down-Left
		if X >= 1 && Y < h-1 {
			var control = mInicial[Y+1][X-1] + init.dist
			var actual = getPointDist(Y+1, X-1, &lVis) // mPath[Y+1][X-1]
			if control < actual {
				//mPath[Y+1][X-1] = control
				setPointDist(X-1, Y+1, control, &lVis)
			}
		}

		//Left
		if X >= 1 {
			var control = mInicial[Y][X-1] + init.dist
			var actual = getPointDist(Y, X-1, &lVis) // mPath[Y][X-1]
			if control < actual {
				//mPath[Y][X-1] = control
				setPointDist(X-1, Y, control, &lVis)
			}
		}

		//Up
		if Y >= 1 {
			var control = mInicial[Y-1][X] + init.dist
			var actual = getPointDist(Y-1, X, &lVis) // mPath[Y-1][X]
			if control < actual {
				//mPath[Y-1][X] = control
				setPointDist(X, Y-1, control, &lVis)
			}
		}

		//Up-Left
		if X >= 1 && Y >= 1 {
			var control = mInicial[Y-1][X-1] + init.dist
			var actual = getPointDist(Y-1, X-1, &lVis) // mPath[Y-1][X-1]
			if control < actual {
				//mPath[Y-1][X-1] = control
				setPointDist(X-1, Y-1, control, &lVis)
			}
		}

		//Up-Right
		if X < w-1 && Y >= 1 {
			var control = mInicial[Y-1][X+1] + init.dist
			var actual = getPointDist(Y-1, X+1, &lVis) // mPath[Y-1][X+1]
			if control < actual {
				//mPath[Y-1][X+1] = control
				setPointDist(X+1, Y-1, control, &lVis)
			}
		}

		init = getNextPoint(&lVis)
	}

	//printList(&lVis)
	fmt.Printf("La distancia mas corta desde la fuente (0,0)\nal extremo opuesto (%d,%d), es %d.", h, w, lVis[points-1].dist)
}
