package main

import (
	"math"
	"math/rand"
	"fmt"
	"strconv"
	"os"
)

var dimension, num int
var k int

// Calculate the distance between two vectors
func vectorDistance(v1, v2 []float64) float64 {
	var summ float64

	for i := 0; i < dimension; i++ {
		summ += math.Pow((v1[i] - v2[i]), 2)
	}

	return summ
}

// Add two vectors position by position
func vectorAdd(v1, v2 []float64) {

	for i := 0; i < dimension; i++ {
		v1[i] += v2[i]
	}

}

// Divide each number in the vector by a specific number of count
func vectorDivide(v []float64, divider int) {

	for i := 0; i < dimension; i++ {
		v[i] /= float64(divider)
	}

}

// Assign new class label of a vector according to its closest centroid
func assignLabel(p *Protein, centroids []Protein) bool {
	least, temp, class := math.MaxFloat64, 0.0, 0

	for _, ctr := range centroids {
		temp = vectorDistance(p.feature, ctr.feature)
		if temp < least {
			least = temp
			class = ctr.class
		}
	}

	if p.class != class {
		p.class = class
		return false
	}

	return true
}

// Get the new centroids via calculating the mean of its related nodes
func calcCentroids(proteins []Protein, centroids []Protein) []Protein {
	// Clear the vector value in the centroids
	for i := 0; i < k; i++ {
		centroids[i].feature = make([]float64, dimension)
	}

	// Count the number of vectors in each class and calculate the
	// their centroids according to their positions
	count := make([]int, k)
	for _, p := range proteins {
		index := p.class - 1
		// fmt.Println(index)
		vectorAdd(centroids[index].feature, p.feature)
		count[index]++
	}

	// Divide the centroids vector by the count of its corresponding vectors
	for i := 0; i < k; i++ {
		vectorDivide(centroids[i].feature, count[i])
	}

	return centroids
}

// assign the label to each nodes and at the mean time check if the
// algorithm is coverged (which means the labels are not changed any more)
func checkConverge(proteins []Protein, centroids []Protein) bool {
	converged := true

	for i := 0; i < num; i++ {
		converged = assignLabel(&proteins[i], centroids) && converged
	}

	return converged
}

func printResult(proteins []Protein, centroids []Protein) {
	result := make(map[int][]string)

	for _, p := range proteins {
		class := p.class
		result[class] = append(result[class], p.label)
	}

	var temp string
	for i := 0; i < k; i++ {
		temp += ("Cluster" + strconv.Itoa(i+1))
		temp += " :{"
		for _, f := range centroids[i].feature {
			temp += strconv.FormatFloat(f, 'g', 1, 64) + " "
		}
		temp += "}\n"
		for j := range result[i+1] {
			temp += (result[i+1][j] + "\n")
		}
		temp += "\n"
	}

	fmt.Print(temp)

	fileName := "k_Means_Clustering.txt"
	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Fprintln(f, temp)
	fmt.Println("Successfully created the file ", fileName)
}

// Initialize the value of centroids random data points from the proteins
func initCentroid(k int, proteins []Protein) []Protein {
	centroids := make([]Protein, k)


	var strList []string
	for _, p := range proteins {
		str := ""
		for _, i := range p.feature {
			str += strconv.Itoa(int(i))
		}
		strList = append(strList, str)
	}

	/* Remove proteins with duplicate features as the candidate for centroids
	   Record their index first */
	encountered := map[string]bool{}
    nonDupN := []int{}

    for v := range strList {
		if encountered[strList[v]] == true {
		    // Do not add duplicate.
		} else {
		    // Record this element as an encountered element.
		    encountered[strList[v]] = true
		    // Append to result slice.
		    nonDupN = append(nonDupN, v)
		}
    }

    nonDup := []Protein{}
    for _, i := range nonDupN {
    	nonDup = append(nonDup, proteins[i])
    }

    // Generate a random permutation of num, and take its first k proteins as the initial centroids
	list := rand.Perm(len(nonDup))

	for i := 0; i < k; i++ {
		index := list[i]
		centroids[i] = nonDup[index]
		// assign the class labels of centroids to their index
		c := &centroids[i]
		c.class = i + 1
		c.label = "centroid"
	}

	return centroids
}

// Run the program with given vectors
func RunKmeans(proteins []Protein, D, N int) {
	dimension, num = D, N

	k = int(math.Sqrt(float64(N)))
	centroids := initCentroid(k, proteins)

	for !checkConverge(proteins, centroids) {
		calcCentroids(proteins, centroids)
	}

	printResult(proteins, centroids)
}
