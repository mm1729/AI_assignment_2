package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type GeneticAlgorithm struct {
	s       *SatProblem
	initPop int
	pop     []Speciman
}

type Speciman struct {
	score int
	value []byte
}

func NewGeneticAlgorithm(s *SatProblem, initPop int) *GeneticAlgorithm {
	//	rand.Seed(int64(initPop * 65))
	return &GeneticAlgorithm{s: s, initPop: initPop}
}

func (g *GeneticAlgorithm) Run() {
	g.createInitPopulation()
	g.eval()
	sort.Sort(ByScore(g.pop))
	fmt.Println(g.pop)

}

/*
   SORT BY SCORE FUNCTIONS
   *implements the Sort interface
*/
type ByScore []Speciman

func (s ByScore) Len() int           { return len(s) }
func (s ByScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByScore) Less(i, j int) bool { return s[i].score < s[j].score }

/*
   RANDOM GENERATOR FUNCTIONS
*/

var r1 *rand.Rand

func newRand() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 = rand.New(s1)
}

func randByte() byte {
	return byte(r1.Intn(2))
}

func initByteArr(arr []byte) {
	newRand()
	for i := range arr {
		arr[i] = randByte()
	}
}

/*
   GENETIC ALGORITHM FUNCTIONS
*/

func (g *GeneticAlgorithm) createInitPopulation() {
	pop := make([][]byte, g.initPop)
	g.pop = make([]Speciman, g.initPop)
	numVar := g.s.NumVar
	allByteArr := make([]byte, g.initPop*numVar)

	initByteArr(allByteArr)

	for i := range pop {
		pop[i], allByteArr = allByteArr[:numVar], allByteArr[numVar:]
		g.pop[i] = Speciman{score: 0, value: pop[i]}
	}

}

func (g *GeneticAlgorithm) eval() {
	index := 0
	for sindex, speciman := range g.pop {
		count := 0
		specimanArr := speciman.value
		for _, clause := range g.s.Map {
			j := byte(0)
			for _, i := range clause {
				if i < 0 {
					j |= (specimanArr[(-1*i)-1] ^ 0x01)
				} else if i > 0 {
					j |= specimanArr[i-1]
				}
			}

			if j == 1 {
				count++
			}
		}
		g.pop[sindex].score = count
		index++
	}
}
