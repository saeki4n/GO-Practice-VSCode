package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	existInFile := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, existInFile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, existInFile)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			tmp := existInFile[line]
			fmt.Printf("%d\t%s\t%s\n", n, line, tmp)
		}
	}
}

func countLines(f *os.File, counts map[string]int, existInFile map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		//ファイル名を格納
		existFlag := false
		for _, item := range existInFile[input.Text()] {
			//ファイル名が既に記録済みかを確認
			if item == f.Name() {
				existFlag = true
			}
		}
		//まだ記録していなかったら保存
		if existFlag == false {
			existInFile[input.Text()] = append(existInFile[input.Text()], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
