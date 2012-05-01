package anagrammer

import (
	"io"
	"bufio"
	"fmt"
	ss "github.com/Leimy/sortstring"
	"os"
	"regexp"
	"sort"
)

var wordValidator = regexp.MustCompile("^[a-zA-Z]+$")

func AnagramsFromReader(in io.Reader) map[string][]string {
	reader := bufio.NewReader(in)
	var nextTok func() (string, error)
	anagrams := make(map[string][]string)

	nextTok = func() (string, error) {
		line, notDone, err := reader.ReadLine()
		if err != nil {
			return "", err
		} 
		for notDone {
			var nextChunk []byte
			nextChunk, notDone, err = reader.ReadLine()
			if err != nil {
				return "", err
			}
			line = append(line, nextChunk...)
		}
		return string(line), err
	}
	
	// nextTok = func() (string, error) {
	// 	line, notDone, err := reader.ReadLine()
	// 	if err != nil {
	// 		return "", err
	// 	} else if notDone {
	// 		restline, err2 := nextTok() // possible problem
	// 		if err2 != nil {
	// 			return "", err2
	// 		}
	// 		return string(line) + string(restline), nil
	// 	}
	// 	return string(line), err
	// }
	
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

func AnagramsFromFile(filename string) map[string][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic("Couldn't open file")
	}

	defer file.Close()
	return AnagramsFromReader(file)
}

func DumpAnagrams(anagrams *map[string][]string) {
	for k, v := range *anagrams {
		fmt.Printf("======\n= %s\n======\n%v\n\n", k, v)
	}

	//	fmt.Printf("%v\n", anagrams)
}
