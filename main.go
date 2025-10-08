package main

import (
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/onyxg/ytsgo/pkg/ytsgoclient"
)

var (
	ytsURL          = flag.String("url", ytsgoclient.DefaultBaseURL, "Base URL of yts.lt API")
	downloadType    = flag.String("t", "", "Type: bluray, web")
	downloadQuality = flag.String("q", "", "Type: 1080p, 720p")
	tailEnabled     = flag.Bool("f", false, "tail the api")
)

func main() {
	flag.Parse()
	c, err := ytsgoclient.New(ytsgoclient.BaseURL(*ytsURL))
	if err != nil {
		log.Fatalf("Failed to create ytsgo client: %v", err)
	}

	switch flag.CommandLine.Arg(0) {
	case "movie":
		id, err := strconv.Atoi(flag.CommandLine.Arg(1))
		if err != nil {
			log.Fatalf("Failed to parse movie ID: %v", err)
		}
		m, err := c.Movie(id)
		if err != nil {
			log.Fatalf("Failed to fetch movie id:%v :%v", id, err)
		}
		fmt.Println(movieStr(m, "", ""))
	case "list":
		opts := []ytsgoclient.ListMoviesOption{
			ytsgoclient.LMSearch(flag.CommandLine.Arg(1)),
			ytsgoclient.LMLimit(50),
		}
		if *downloadQuality != "" {
			opts = append(opts, ytsgoclient.LMQuality(*downloadQuality))
		}

		mvs, err := c.ListMovies(opts...)

		if *tailEnabled {
			page := 1
			for {
				fmt.Printf("************** page: %d\n", page)
				opts = append(opts, ytsgoclient.LMPage(uint(page)))
				mvsNext, err := c.ListMovies(opts...)
				if err != nil {
					log.Printf("Failed to fetch movie list: %v", err)
					break
				}
				if mvsNext == nil || len(mvsNext.Movies) == 0 {
					log.Printf("completed listing")
					break
				}

				mvs.Movies = append(mvs.Movies, mvsNext.Movies...)

				page++
			}
		}

		if err != nil {
			log.Fatalf("Failed to search movies %q :%v", flag.CommandLine.Arg(1), err)
		}
		for _, m := range mvs.Movies {
			fmt.Println(movieStr(m, *downloadType, *downloadQuality))
		}
	default:
		usage()
		return
	}
}

func usage() {
	fmt.Printf(`Usage:
ytsgo movie [id]
ytsgo list "search term"
`)
}

func movieStr(m *ytsgoclient.Movie, downloadType, downloadQuality string) string {
	ret := fmt.Sprintf("%q (%v)\n", m.Title, m.Year)
	var trts []string
	sort.Sort(sort.Reverse(ytsgoclient.TorrentsBySize(m.Torrents)))
	for _, t := range m.Torrents {
		if downloadType != "" && downloadType != t.Type {
			continue
		}
		if downloadQuality != "" && downloadQuality != t.Quality {
			continue
		}

		trts = append(trts, fmt.Sprintf("\tQuality: %v, Type: %v, Seeds: %v Peers: %v Size: %v",
			t.Quality, t.Type, t.Seeds, t.Peers, t.Size),
		)
	}

	if len(trts) > 0 {
		return ret + fmt.Sprintf("\tNo Torrents Matched")
	}

	return ret + strings.Join(trts, "\n")
}
