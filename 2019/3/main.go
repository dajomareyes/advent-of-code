package main

import (
	"errors"
	"fmt"
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
	A float64
	B float64
	C float64
}

func findIntersectionCoordinates(l1, l2 LineEquation) (Coordinate, error) {
	determinant := l1.A*l2.B - l2.A*l1.B
	if determinant == 0 {
		// lines are parallel
		return Coordinate{}, errors.New("The lines are parallel")
	}

	x := (l2.B*l1.C - l1.B*l2.C) / determinant
	y := (l1.A*l2.C - l2.A*l1.C) / determinant

	return Coordinate{x, y}, nil
}

func findIntersection(l1, l2 Line) (Coordinate, error) {
	if l1.slope == l2.slope {
		return Coordinate{}, errors.New("The lines don't intersect")
	}
	x := (l2.yint - l1.yint) / (l1.slope - l2.slope)
	if &l2.x != nil {
		x = l2.x
	}
	fmt.Println(l1.slope, l2.slope)
	y := l1.slope*x + l1.yint

	return Coordinate{x, y}, nil
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
	return LineEquation{A, B, C}
}

func createLine(a, b Coordinate) Line {
	slope := (a.y - b.y) / (a.x - b.x)
	yint := a.y - slope*a.x
	if (a.x - b.x) == 0 {
		return Line{slope, yint, a.x}
	}
	result := Line{}
	result.slope = slope
	result.yint = yint
	return result
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
