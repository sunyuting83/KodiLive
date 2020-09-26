package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// GetList Get List
func GetList(n string, l string, t bool) (data string) {
	n = strings.Split(n, ".")[0]
	return LiveScrape(n, l, t, true)
}

// LiveScrape get live list
func LiveScrape(name string, local string, t bool, cors bool) (List string) {
	// Request the HTML page.
	var (
		err error
		url string
	)
	List = "#EXTM3U\n"
	url = strings.Join([]string{"https://members.sexcamvideos.net", name}, "/")
	if name == "index" {
		url = "https://members.sexcamvideos.net"
	}
	if t {
		url = strings.Join([]string{"https://members.sexcamvideos.net", "tag", name}, "/")
	}
	if cors {
		url = strings.Join([]string{"https://cors.zme.ink", url}, "/")
	}
	// fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return List
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return List
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return List
	}
	root := doc.Find("body ul#room_list > li.room_list_room")
	imgurl := "https://roomimg.stream.highwebmedia.com/ri"
	root.Each(func(index int, ele *goquery.Selection) {
		title := ele.Find("div.details > div.title > a").Text()
		title = strings.Replace(title, " ", "", -1)
		t := time.Now().Unix()
		tt := strconv.FormatInt(t, 10)
		cc := strings.Join([]string{imgurl, title}, "/")
		cc = strings.Join([]string{cc, "jpg?"}, ".")
		cover := strings.Join([]string{cc, tt}, "")
		first := `#EXTINF:-1 tvg-id="" tvg-name="`
		cv := `" tvg-language="English" tvg-logo="`
		after := `" group-title="livecam",RTVA`
		r := "\n"
		m3u8 := strings.Join([]string{local, "livecam", title, "playlist.m3u8"}, "/")
		m3u8 = strings.Join([]string{m3u8, "\n"}, "")
		var str []string = []string{first, title, cv, cover, after, r, m3u8}
		List += strings.Join(str, "")
	})
	return List
}
