package main

import "bufio"
import "fmt"
import "io/ioutil"
import "log"
import "os"
import "strconv"
import "strings"

type SatProblem struct {
	NumVar     int
	NumClauses int
	Map        map[int][3]int
}

func main() {
	fmt.Println("Hello, world!")
	//readCFFiles()
	s := processCFFile("./uf/uf20/uf20-01.cnf")
	fmt.Println(s)
}

/***********************************************************
Genetic algorithm methods
************************************************************/

/***********************************************************
File reading methods
************************************************************/
func readCFFiles() {
	files, err := ioutil.ReadDir("./uf")
	check(err)

	for _, file := range files {
		fileName := file.Name()
		if strings.HasPrefix(fileName, "uf") {
			//process this file

		}
	}
}

func processCFFile(fileName string) (s *SatProblem) {
	f, err := os.Open(fileName)
	check(err)

	s = new(SatProblem)
	i := 1
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil || len(line) == 0 {
			break
		}

		skip, create, a1, a2, a3 := processCFLine(line)
		if skip {
			continue
		}
		if create {
			s.NumVar = a1
			s.NumClauses = a2
			s.Map = make(map[int][3]int)
		} else {
			s.Map[i] = [3]int{a1, a2, a3}
		}
		i++
	}
	f.Close()
	return
}

func processCFLine(line string) (skip bool, create bool, a1 int, a2 int, a3 int) {

	items := strings.Fields(line)
	lineType := string(items[0])

	switch lineType {
	case "0":
		fallthrough
	case "%":
		fallthrough
	case "c":
		return true, false, 0, 0, 0
	case "p":
		if len(items) != 4 {
			log.Fatal("unknown line format: " + line)
		}

		return false, true, Int(items[2]), Int(items[3]), 0
	default:
		if len(items) != 4 {
			log.Fatal("unknown line format: " + line)
		}

		return false, false, Int(items[1]),
			Int(items[2]), Int(items[3])
	}
}

func Int(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
