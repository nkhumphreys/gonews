package main

import "sync"

const nuviNewsURL = "http://feed.omgili.com/5Rh5AMTrc4Pv/mainstream/posts/"

var wg sync.WaitGroup

func main() {
	fileNames := getFileNames()

	c := make(chan string)

	wg.Add(1)

	go getFiles(fileNames, c)
	processFiles(c)

	wg.Wait()
}