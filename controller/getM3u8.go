package controller

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetM3u8 get m3u8 fun
func GetM3u8(n string, proxy bool) (m string) {
	return getM3u8Scrape(n, proxy)
}

// LiveScrape get live list
func getM3u8Scrape(t string, cors bool) (live string) {
	// Request the HTML page.
	var (
		err error
		url string
	)
	live = "#EXTM3U\n#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=800000,RESOLUTION=1080x608\n"
	if len(t) <= 0 {
		return live
	}
	url = strings.Join([]string{"https://members.sexcamvideos.net", t}, "/")
	if cors {
		url = strings.Join([]string{"https://cors.zme.ink", url}, "/")
	}
	// fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return live
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return live
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return live
	}
	doc.Find("body script").Each(func(index int, ele *goquery.Selection) {
		if strings.Contains(ele.Text(), "window.initialRoomDossier") {
			s := ele.Text()
			x := strings.Index(s, `hls_source\u0022: \u0022`) + 24
			y := strings.Index(s, `\u0022, \u0022allow_show_recordings`)
			s = s[x:y]
			d := strings.Replace(s, `\u002D`, `-`, -1)
			// fmt.Println(d)
			live = strings.Join([]string{live, d}, "")
		}
	})

	return live
}
