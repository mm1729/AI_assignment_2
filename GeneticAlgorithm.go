package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GeneticAlgorithm struct {
	s       *SatProblem
	initPop int
	pop     [][]byte
	scores  []int
}

var r1 *rand.Rand

func NewGeneticAlgorithm(s *SatProblem, initPop int) *GeneticAlgorithm {
	//	rand.Seed(int64(initPop * 65))
	return &GeneticAlgorithm{s: s, initPop: initPop,
		scores: make([]int, initPop, initPop)}
}

func (g *GeneticAlgorithm) Run() {
	g.createInitPopulation()
	g.eval()
	fmt.Println(g.pop)
	fmt.Println(g.scores)
}

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

func (g *GeneticAlgorithm) createInitPopulation() {
	g.pop = make([][]byte, g.initPop)
	numVar := g.s.NumVar
	allByteArr := make([]byte, g.initPop*numVar)

	initByteArr(allByteArr)

	for i := range g.pop {
		g.pop[i], allByteArr =
			allByteArr[:numVar], allByteArr[numVar:]
	}
}

func (g *GeneticAlgorithm) eval() {
	index := 0
	for _, speciman := range g.pop {
		count := 0

		for _, clause := range g.s.Map {
			j := byte(0)
			for _, i := range clause {
				if i < 0 {
					j |= (speciman[(-1*i)-1] ^ 0x01)
				} else if i > 0 {
					j |= speciman[i-1]
				}
			}

			if j == 1 {
				count++
			}
		}

		g.scores[index] = count
		index++
	}
}
