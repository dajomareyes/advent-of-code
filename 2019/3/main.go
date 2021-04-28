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

func findIntersection(l1, l2 Line) (Coordinate, error) {
	if l1.slope == l2.slope {
		return Coordinate{}, errors.New("The lines don't intersect")
	}
	x := (l2.yint - l1.yint) / (l1.slope - l2.slope)
	if &l2.x != nil {
		x = l2.x
	}
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
	for i, s := range path {
		fmt.Println(i, s)
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
	fmt.Println(coordPath)
	return coordPath
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

func main() {
	fmt.Println("Starting application")
	path1 := strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ",")
	path2 := strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ",")
	fmt.Println(path1, path2)
	// cord1 := generatePath(path1)
	// cord2 := generatePath(path2)
	// TODO: Find intersections between path

	l1 := createLine(Coordinate{0, 0}, Coordinate{0, 75})
	l2 := createLine(Coordinate{75, 2}, Coordinate{-75, 2})
	fmt.Println(l1)
	fmt.Println(l2)

	fmt.Println(findIntersection(l1, l2))
}
