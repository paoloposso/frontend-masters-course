package main

import (
	"fmt"
	"math/rand"
	"time"
)

const block = "#"
const goal = "$"
const current = "|"
const seen = "."

type Point struct {
	x int
	y int
}

type Maze [][]string

func generateRandomMaze(width int, height int) (maze [][]string, start Point) {
	for i := 0; i < width; i++ {
		line := []string{}

		for j := 0; j < height; j++ {
			line = append(line, " ")
		}
		maze = append(maze, line)
	}

	end := getRandomPoint(width, height)
	maze[end.x][end.y] = goal

	start = getRandomPoint(width, height)
	maze[start.x][start.y] = current

	return maze, start
}

func printMaze(maze [][]string) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			fmt.Printf("[%s]", maze[i][j])
		}
		fmt.Println()
	}
}

func getRandomPoint(maxX, maxY int) Point {

	result := Point{}

	rand.Seed(time.Now().UnixNano())
	result.x = rand.Intn(maxX)

	rand.Seed(time.Now().UnixNano())
	result.y = rand.Intn(maxY)

	return result
}

func getDescendants(point Point) []Point {
	descendants := []Point{}
	descendants = append(descendants, Point{x: point.x + 1, y: point.y})
	descendants = append(descendants, Point{x: point.x - 1, y: point.y})
	descendants = append(descendants, Point{x: point.x, y: point.y + 1})
	descendants = append(descendants, Point{x: point.x, y: point.y - 1})
	return descendants
}

func solve(maze [][]string, start Point) {
	path := []Point{}
	dfs(maze, start, path)
}

func dfs(maze [][]string, current Point, path []Point) bool {
	if current.x < 0 || current.x >= len(maze) || current.y < 0 || current.y >= len(maze[0]) {
		return false
	}
	if maze[current.x][current.y] == seen {
		return false
	}
	if maze[current.x][current.y] == goal {
		return true
	}

	descendants := getDescendants(current)

	for _, des := range descendants {
		path = append(path, des)
		maze[current.x][current.x] = seen
		if dfs(maze, des, path) {
			return true
		}
	}

	path = path[:len(path)-1]

	return false
}

func main() {
	maze, start := generateRandomMaze(10, 10)
	printMaze(maze)

	solve(maze, start)

	// fmt.Println(path)
}
