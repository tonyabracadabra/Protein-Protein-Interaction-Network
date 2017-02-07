package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var gpMap map[string]string // Map gene to protein name -> map[Gene]Protein
var gnMap map[string]int // Map gene to protein index -> map[Gene]Index
var pnMap map[string]int // Map protein name to protein index -> map[Protein]Index

// For K-Means clustering
type Protein struct {
	feature []float64
	class   int
	label   string
}

// Read nodes file that given by the user input
func readGenesInfoFromTxt() ([]Protein, int, int) {
	genesInfoFile, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println("Error: Something wrong with the proteins file.")
		os.Exit(3)
	}
	defer genesInfoFile.Close()

	// Read infomations from the file into a buffer
	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(genesInfoFile)

	for scanner.Scan() {
		// Append the content of the file to the lines slice
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Sorry: There was some kind of error during the file reading")
		os.Exit(3)
	}

	/* Parsing the file
	 * into protein infomations
	 */

	// Get number of genes and the dimension of each feature within the file
	num, dimension := len(lines)-1, len(strings.Split(lines[1], "	"))-1

	// Initialize the protein informations
	proteins := make([]Protein, num)

	for i := 0; i < num; i++ {
		featureArray := strings.Split(lines[i+1], "	") // tab deliminated
		p := &proteins[i]
		// Find protein based on gene with gpMap generated before
		p.label = gpMap[featureArray[0]]
		p.feature = make([]float64, dimension)
		for j := 0; j < dimension; j++ {
			val, _ := strconv.ParseFloat(featureArray[j+1], 64)
			p.feature[j] = val
		}
	}

	return proteins, dimension, num
}

/* The annotation file is based on genes, so we have to map those genes to their main coded proteins
 * Also, we have to map the numbers of the proteins for the better implementation of the algorithm
 */
func mapProteinsFromGenes() {
	mapFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: Something wrong with the mapping file.")
		os.Exit(3)
	}
	defer mapFile.Close()

	// Read infomations from the file into a buffer
	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(mapFile)

	for scanner.Scan() {
		// Append the content of the file to the lines slice
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Sorry: There was some kind of error during the file reading")
		os.Exit(3)
	}

	/* Parsing file */

	gpMap, gnMap, pnMap = make(map[string]string), make(map[string]int), make(map[string]int)

	for i := 0; i < len(lines)-1; i++ {
		pair := strings.Split(lines[i+1], "	")
		gpMap[pair[1]] = pair[0]
		gnMap[pair[1]] = i
		pnMap[pair[0]] = i
	}
	
}

func readEdgesFile(proteins []Protein) []Edge {
	edgesFile, err2 := os.Open(os.Args[3])
	if err2 != nil {
		fmt.Println("Error: Something wrong with the edge file.")
		os.Exit(3)
	}
	defer edgesFile.Close()

	// Read infomations from the file into a buffer
	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(edgesFile)

	for scanner.Scan() {
		// Append the content of the file to the lines slice
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Sorry: There was some kind of error during the file reading")
		os.Exit(3)
	}

	/* Parsing the file
	 * into edges infomations
	 */

	var edges []Edge

	// Each line is a pair of genes interacting each other, so store them in the edge
	for _, l := range lines[1:] {
		line := strings.Split(l, "	") // Deliminated by tab
		// Find the protein based on the index stored before
		index1, index2 := gnMap[line[0]], gnMap[line[1]]
		fromProtein, toProtein := proteins[index1], proteins[index2]
		// Calculating the weight of the edge based on the protein features
		weight := calcWeight(fromProtein.feature, toProtein.feature)
		edges = append(edges, CreateNewEdge(index1, index2, weight))
	}

	return edges
}

func calcWeight(v1, v2 []float64) float64 {
	var summ float64

	for i := 0; i < len(v1); i++ {
		summ += math.Pow((v1[i] - v2[i]), 2)
	}

	return math.Sqrt(summ)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Build protein-gene pair
	mapProteinsFromGenes()
	proteins, dimension, num := readGenesInfoFromTxt()
	
	edges := readEdgesFile(proteins)

	RunKmeans(proteins, dimension, num)
	RunDijkstra(num, edges, gnMap, pnMap, proteins)
}
