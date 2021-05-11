package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coordinate struct {
	x float64
	y float64
}

type Line struct {
	slope float64
	yint  float64
	x     float64
}

type LineEquation struct {
	A  float64
	B  float64
	C  float64
	Xs float64
	Xf float64
	Ys float64
	Yf float64
}

func pointExistsBetweenSegment(line LineEquation, x float64, y float64) bool {

	validX := math.Min(line.Xs, line.Xf) <= x && x <= math.Max(line.Xs, line.Xf)
	validY := math.Min(line.Ys, line.Yf) <= y && y <= math.Max(line.Ys, line.Yf)

	fmt.Println(validX, validY)
	return validX && validY
}

func findIntersectionCoordinates(l1, l2 LineEquation) (Coordinate, error) {
	determinant := l1.A*l2.B - l2.A*l1.B
	if determinant == 0 {
		// lines are parallel
		return Coordinate{}, errors.New("The lines are parallel")
	}

	x := (l2.B*l1.C - l1.B*l2.C) / determinant
	y := (l1.A*l2.C - l2.A*l1.C) / determinant

	validPoint := pointExistsBetweenSegment(l1, x, y) && pointExistsBetweenSegment(l2, x, y)

	if validPoint {
		return Coordinate{x, y}, nil
	}

	return Coordinate{}, errors.New("point exists outside of line segment")
}

func getCommandCoordinate(command string) Coordinate {
	com := command[0]
	multiplier, err := strconv.ParseFloat(command[1:], 64)
	if err != nil {
		fmt.Println(err)
	}
	coord := Coordinate{0, 0}
	switch com {
	case 'R':
		coord = Coordinate{1, 0}
	case 'L':
		coord = Coordinate{-1, 0}
	case 'D':
		coord = Coordinate{0, -1}
	case 'U':
		coord = Coordinate{0, 1}
	}
	coord.x *= multiplier
	coord.y *= multiplier
	return coord
}

func generatePath(path []string) []Coordinate {
	start := Coordinate{0, 0}
	var coordPath []Coordinate
	for _, s := range path {
		if len(coordPath) == 0 {
			coordPath = append(coordPath, start)
			coordPath = append(coordPath, getCommandCoordinate(s))
		} else {
			last := len(coordPath) - 1
			cmd := getCommandCoordinate(s)
			newPoint := Coordinate{coordPath[last].x + cmd.x, coordPath[last].y + cmd.y}
			coordPath = append(coordPath, newPoint)
		}
	}
	return coordPath
}

func createLineEquation(a, b Coordinate) LineEquation {
	A := b.y - a.y
	B := a.x - b.x
	C := A*a.x + B*b.y
	Xs := a.x
	Xf := b.x
	Ys := a.y
	Yf := b.y
	return LineEquation{A, B, C, Xs, Xf, Ys, Yf}
}

func createLines(path []Coordinate) []LineEquation {
	var lines []LineEquation
	for i := 0; i < len(path)-1; i++ {
		j := i + 1
		lines = append(lines, createLineEquation(path[i], path[j]))
	}
	return lines
}

func main() {
	fmt.Println("Starting application")

	// generate paths from input
	// example 2
	//path1 := generatePath(strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ","))
	//path2 := generatePath(strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ","))

	// example 1
	path1 := generatePath(strings.Split("R8,U5,L5,D3", ","))
	path2 := generatePath(strings.Split("U7,R6,D4,L4", ","))

	fmt.Println(path1)
	fmt.Println(path2)

	// generate lines
	lines1 := createLines(path1)
	lines2 := createLines(path2)

	var intersections []Coordinate

	for i := 0; i < len(lines1); i++ {
		for j := 0; j < len(lines2); j++ {
			coord, err := findIntersectionCoordinates(lines1[i], lines2[j])
			if err != nil {
				continue
			}
			intersections = append(intersections, coord)
		}
	}

	// TODO: intersesctions are coming out wrong because it things that each of the lines are infinite
	// you might want to look into using vectors instead since vectors have magnitude
	fmt.Println(intersections)

	// l1 := createLineEquation(Coordinate{0, 0}, Coordinate{75, 0})
	// l2 := createLineEquation(Coordinate{0, 0}, Coordinate{0, 62})
	// fmt.Println(l1)
	// fmt.Println(l2)

	// fmt.Println(findIntersectionCoordinates(l1, l2))
}
