/**
 * User: Medcl
 * Date: 13-7-8
 * Time: 下午4:45
 */
package main

import (

	"log"
	"regexp"
    . "hunter"
	"flag"
)

var seed_url = flag.String("seed", "", "Seed URL")

func main() {
    log.Print("[gopa] is on.")

	flag.Parse()
	if *seed_url == "" {
		log.Fatal("no seed was given.")
	}

    curl := make(chan []byte)
    success := make(chan Task)
    failure := make(chan string)

    visited := make(map[string]int)

    regex := regexp.MustCompile("<a.*?href=[\"'](http.*?)[\"']")
//    download_regex := regexp.MustCompile("<a.*?href=[\"'](http.*?)[\"']")


    // Give our crawler a place to start.
    go Seed(curl,*seed_url)

    // Start the throttled crawling.
    go ThrottledCrawl(curl, success, failure, visited)

    // Main loop that never exits and blocks on the data of a page.
    for {
        site := <-success
        go GetUrls(curl, site, regex)
    }

}

