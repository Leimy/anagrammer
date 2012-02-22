package anagrammer

import (
	"fmt"
	"os"
	"sort"
	"bufio"
	ss "gofun/sortstring"
	"regexp"
)

var wordValidator = regexp.MustCompile("^[a-zA-Z]+$")

func AnagramsFromFile(filename string) map[string][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic("Couldn't open file")
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var nextTok func () (string, error) 

	nextTok = func () (string, error) {
		line, notDone, err := reader.ReadLine()
		if err != nil {
			return "", err
		} else if notDone {
			restline, err2 := nextTok() // possible problem
			if err2 != nil {
				return "", err2
			}
			return string(line) + string(restline), nil
		}
		return string(line), err
	}


	anagrams := make(map[string][]string)

	line, linerr := nextTok()

	for linerr == nil {
		if !wordValidator.MatchString(line) { // skip non validating lines
			line, linerr = nextTok()
			continue
		}
		ss := ss.NewSortString(line)
		sort.Sort(ss)
		index := ss.String()
		slice := anagrams[index]
		if slice == nil {
			slice = make([]string, 1)
		}
		anagrams[index] = append(slice, line)
		line, linerr = nextTok()
	}

	return anagrams
}

func DumpAnagrams(anagrams *map[string][]string) {
	for k, v := range *anagrams {
		fmt.Printf("======\n= %s\n======\n%v\n\n", k, v)
	}

	//	fmt.Printf("%v\n", anagrams)
}
