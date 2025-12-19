package main

import (
	"aoc/shared"
	"cmp"
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/*
Input:
- 1 junction box per line -> x, y, z coord per line
- We want to use these the coordinates to find the two closest junction boxes and connect them
- Once connected, junction boxes become one circuit
- An already connected junction box can still be considered for future connections
- You want to just find the next shortest connection

Output:
- Product of the number of junction boxes in the three largest circuits after doing 1000 of the shortest connections
- An already connected junction box can still be considered for future connections

Brainstorming:
- Definitely need to keep track of circuits -> graphs? arrays? ...
- Will need to use euclidian distance formula for getting distance
- Just need the first 1000 shortest distances but there are literally a thousand coordinates

First Approach:
- Get distance from every point to every other point (1000 x 1000) size map
- Sort map by their distance
- Get the first 1000 distances and connect those nodes if there isn't a connect already
- Traverse each graph and count how many nodes there are in each graph. Multiply the three biggest numbers
- Time: get distance for every point + sort the distances + connect the nodes in first 1000 + traverse every possible graph and get their number of nodes + sort by distances + get output
- Time: O(n^2) + O(nlogn) + O(n) + O(n^2) + O(nlogn) + O(1) = O(n^2)
- Space: O(n) -> if we are just using node graphs

Other considerations:
- What if we use an adjacency list instead -> could just sort by length of adjacency list -> would skip that node traversal thing
- We try both
- Wait, even with nodes and pointers, nodes can still have multiple connections making it O(n^2) space, so I think an adjacency list is fine
Might not work, also realized something. It is impossible to not have a fully connected graph if your edges is equal to your nodes nvm
Algorithm:
- Create list of distances
- Do a nested for loop to calculate the distances between every possible pair of points. Store this in array of {key: distance: value: (coord, coord)}
- Do a sorted func that sorts this by distance
- For the first 1000 instances, create nodes for the coords, store them in the node map if not already there, and make a connection
- For every node in the node map, do a graph traversal and count the amount of nodes that are visited -> store this value in array
- Keep on traversing unvisited nodes, until there are none.
- Sort length array and get the top 3
*/

type Coord struct {
	X int
	Y int
	Z int
}

type Junction struct {
	Location Coord
	Edges    []*Junction
}

type DistancePair struct {
	Distance float64
	Pair     [2]*Junction
}

func compareDistance(a DistancePair, b DistancePair) int {
	return cmp.Compare(a.Distance, b.Distance)
}

func getDistance(junction Junction, nextJunction Junction) float64 {
	return math.Sqrt(math.Pow(float64(nextJunction.Location.X-junction.Location.X), 2) + math.Pow(float64(nextJunction.Location.Y-junction.Location.Y), 2) + math.Pow(float64(nextJunction.Location.Z-junction.Location.Z), 2))
}

func part1(input string) {
	//Create list of junctions
	junctions := []*Junction{}
	for junction := range strings.SplitSeq(input, "\n") {
		coords := strings.Split(junction, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		newJunction := Junction{Location: Coord{X: x, Y: y, Z: z}, Edges: []*Junction{}}
		junctions = append(junctions, &newJunction)
	}

	//NOTE: Made it into a set for better time complexity
	connections := make(map[[2]Coord]struct{})
	distances := []DistancePair{}
	for _, junction := range junctions {
		for _, nextJunction := range junctions {
			_, inConnections := connections[[2]Coord{nextJunction.Location, junction.Location}]
			if nextJunction.Location == junction.Location || inConnections {
				continue //NOTE: Has already made this connection
			} else {
				distance := getDistance(*junction, *nextJunction)
				distances = append(distances, DistancePair{Distance: distance, Pair: [2]*Junction{junction, nextJunction}})
				connections[[2]Coord{junction.Location, nextJunction.Location}] = struct{}{}
			}
		}
	}

	slices.SortFunc(distances, compareDistance)
	distances = distances[:1000]

	for _, distance := range distances {
		junctions := distance.Pair
		j1 := junctions[0]
		j2 := junctions[1]
		j1.Edges = append(j1.Edges, j2)
		j2.Edges = append(j2.Edges, j1)
		// fmt.Println(*&j1.Location, *&j2.Location)
	}

	visited := map[Coord]struct{}{}
	sizes := []int{}
	currSize := 0

	var dfs func(junction *Junction)
	dfs = func(junction *Junction) {
		if _, ok := visited[junction.Location]; ok {
			return
		}

		visited[junction.Location] = struct{}{}
		currSize += 1

		for _, junction := range junction.Edges {
			dfs(junction)
		}
	}

	numCircuits := 0
	for _, junction := range junctions {
		currSize = 0
		fmt.Println("Graph starting at junction", junction.Location, currSize)
		dfs(junction)
		if currSize > 0 {
			fmt.Println("Ended this traversal", currSize)
			sizes = append(sizes, currSize)
			numCircuits += 1
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j] // note >
	})

	output := sizes[0] * sizes[1] * sizes[2]

	fmt.Println(sizes)
	fmt.Println(numCircuits)
	fmt.Println(output)
}

/*
Update:
- need to keep connecting based on shortest distance, until all junction boxes form one big circuit
- To do this, just need to keep track of a visited map while you are calculating the distances.
- As soon as the length of the map is equal to the length of the junctions, you have connected all of the junctions
- nevermind, just because you made a connection with every node, does not mean they man one big graph
- if you traverse, and there is more than one circuit you know that it has not reached the state yet
- so what if we traverse the graphs on every distance addition, as soon as you get more than one circuit you stop traversing, but, if you do an entire traversal, there is only one circuit you return the output?
- Might not work, also realized something. It is impossible to not have a fully connected graph if your edges is equal to your nodes nvm
- What if we connect every node, then remove edges in reverse distance order until the number of edges is equal to the number of junctions
*/

func part2(input string) {
	//Create list of junctions
	junctions := []*Junction{}
	for junction := range strings.SplitSeq(input, "\n") {
		coords := strings.Split(junction, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		newJunction := Junction{Location: Coord{X: x, Y: y, Z: z}, Edges: []*Junction{}}
		junctions = append(junctions, &newJunction)
	}

	//NOTE: Made it into a set for better time complexity
	connections := make(map[[2]Coord]struct{})
	distances := []DistancePair{}
	for _, junction := range junctions {
		for _, nextJunction := range junctions {
			_, inConnections := connections[[2]Coord{nextJunction.Location, junction.Location}]
			if nextJunction.Location == junction.Location || inConnections {
				continue //NOTE: Has already made this connection
			} else {
				distance := getDistance(*junction, *nextJunction)
				distances = append(distances, DistancePair{Distance: distance, Pair: [2]*Junction{junction, nextJunction}})
				connections[[2]Coord{junction.Location, nextJunction.Location}] = struct{}{}
			}
		}
	}

	slices.SortFunc(distances, compareDistance)

	visited := map[Coord]struct{}{}

	var dfs func(junction *Junction)
	dfs = func(junction *Junction) {
		if _, ok := visited[junction.Location]; ok {
			return
		}

		visited[junction.Location] = struct{}{}

		for _, junction := range junction.Edges {
			dfs(junction)
		}
	}

	for _, distance := range distances {
		junctionPair := distance.Pair
		j1 := junctionPair[0]
		j2 := junctionPair[1]
		j1.Edges = append(j1.Edges, j2)
		j2.Edges = append(j2.Edges, j1)
		visited = map[Coord]struct{}{}
		dfs(junctions[0])
		if len(visited) == len(junctions) {
			output := j1.Location.X * j2.Location.X
			fmt.Println(output)
			break
		}
	}
}

func main() {
	input := utils.GetInputString("test.txt")
	part2(input)
}
