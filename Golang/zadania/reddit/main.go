package main

import (
	"context"
	"log"
	"os"
	"reddit/fetcher"
	"sync"
	"time"
)

func main() {
	log.Println("Starting Reddit Fetcher...")
	ufs := []urlAndFilename{
		{"https://www.reddit.com/r/golang.json", "golang.txt"},
		{"https://www.reddit.com/r/programmerhumor.json", "programmerhumor.txt"},
		{"https://www.reddit.com/r/programminghorror.json", "programminghorror.txt"},
		{"https://www.reddit.com/r/cscareerquestions.json", "cscareerquestions.txt"},
	}

	var wg sync.WaitGroup
	wg.Add(len(ufs))

	for _, uf := range ufs {
		go CreateAndRunFetcher(uf.url, uf.filename, &wg)
	}

	log.Println("Waiting for all fetchers to finish...")
	wg.Wait()
	log.Println("All fetchers finished.")
}

func CreateAndRunFetcher(url string, outputFilename string, wg *sync.WaitGroup) {
	defer wg.Done()

	var client fetcher.RedditFetcher
	client = fetcher.NewFetcher(url)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Fetch(ctx); err != nil {
		log.Printf("error while fetching data: %v", err)
	}

	file, err := os.Create(outputFilename)
	if err != nil {
		log.Printf("error while creating file: %v", err)
	}
	defer func() {
		_ = file.Close()
	}()

	if err = client.Save(file); err != nil {
		log.Printf("error while saving data: %v", err)
	}

	log.Printf("Data from %s saved to %s", url, outputFilename)
}

type urlAndFilename struct {
	url      string
	filename string
}
