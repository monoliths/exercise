package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"net/http/httptrace"
	"time"
	"bufio"
	"net/url"
	"strings"
	"sync"
	"strconv"
)

// Used for "pretty printing" final results for easy visual grepping
const (
	ColorReset = "\033[0m"
	ColorRed = "\033[31m"
	ColorGreen = "\033[32m"
)

// yank filename from a URL
func fileNameFromUrl(rawUrl string) string {
	fileUrl, err := url.Parse(rawUrl)
    if err != nil {
        panic(err)
    }
    path := fileUrl.Path
    segments := strings.Split(path, "/")
    fileName := segments[len(segments)-1]
	return fileName
}

// safe to read lines into memory, challage states "from 1 to 100s of URLs" 
func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

/// Downloads a file and notes results in ttfbs and totals slices
func downloadFile(downloadDir string, fileName string, url string, wg *sync.WaitGroup, client *http.Client, count int, ttfbs []int64, totals []int64) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	start := time.Now() 
	var ttfb time.Duration
	trace := &httptrace.ClientTrace{
        GotFirstResponseByte: func() {
			t := time.Now()
			ttfb = t.Sub(start)
			ttfbs[count] = ttfb.Milliseconds()
        },
    }
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
    }

	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(downloadDir + fileName)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
    }
	out.Close()

	total := time.Since(start)
	totals[count] = total.Milliseconds()
	wg.Done()
}

func printIndividualResults(url string, ttfb int64, targetTtfb int64, total int64, targetTotal int64) {
	result := fmt.Sprintf("%s \tTTFB: %vms \tTotal: %vms", url, ttfb, total)
	if (ttfb > targetTtfb || total > targetTotal) {
		fmt.Println(ColorRed, result, ColorReset)
	} else {
		fmt.Println(result)
	}
}

func average(collection []int64) int64 {
	collectionSize := len(collection)
	var sum int64 = 0
	for _,item := range collection {
		sum += item
	}
	// we dont need precision here since its in ms
	return (sum / int64(collectionSize))
}

func main() {
	urlFile := os.Args[1]
	downloadDir := os.Args[2] + "/"
	targetTtfb,_ := strconv.ParseInt(os.Args[3], 10, 64)
	targetTotal,_ := strconv.ParseInt(os.Args[4], 10, 64)

	urls, err :=  readLines(urlFile)
	if err != nil {
        panic(err)
    }

	urlCount := len(urls)
	ttfbs := make([]int64, urlCount)
	totals := make([]int64, urlCount)

	client := &http.Client{}
	var wg sync.WaitGroup

	count := 0
    for _,url := range urls {
		wg.Add(1)
		// This will allow me to provide unique downloaded filenames for the same urls
		fileName := fileNameFromUrl(url) + "-" + strconv.Itoa(count)
		go downloadFile(downloadDir, fileName, url, &wg, client, count, ttfbs, totals)
		count++
    }
	wg.Wait()

	fmt.Println("===================== Individual Results ====================")
	for i := 0; i < urlCount; i++ {
		printIndividualResults(urls[i], ttfbs[i], targetTtfb, totals[i], targetTotal)
	}
	
	fmt.Println("===================== Average Results ====================")
	ttfbAverage := average(ttfbs)
	totalAverage := average(totals)

	if (targetTtfb > ttfbAverage) {
		fmt.Println(ColorGreen, fmt.Sprintf("TTFB PASS @ %vms", ttfbAverage), ColorReset)
	} else {
		fmt.Println(ColorRed, fmt.Sprintf("TTFB FAIL @ %vms", ttfbAverage), ColorReset)
	}

	if (targetTotal > totalAverage) {
		fmt.Println(ColorGreen, fmt.Sprintf("TOTAL PASS @ %vms", totalAverage), ColorReset)
	} else {
		fmt.Println(ColorRed, fmt.Sprintf("TOTAL FAIL @ %vms", totalAverage), ColorReset)
	}
}	
