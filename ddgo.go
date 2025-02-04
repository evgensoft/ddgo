package ddgo

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Result holds the returned query data
type Result struct {
	Title string
	Info  string
	URL   string
}

// Requests the query and puts the results into an array
func Query(query string, maxResult int) ([]Result, error) {
	results := []Result{}
	queryUrl := fmt.Sprintf("https://duckduckgo.com/html/?q=%s", url.QueryEscape(query))

	response, err := http.Get(queryUrl)
	if err != nil {
		return results, fmt.Errorf("get %v error: %w", queryUrl, err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return results, fmt.Errorf("status code error: %d %s", response.StatusCode, response.Status)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return results, fmt.Errorf("goquery.NewDocument error: %w", err)
	}

	sel := doc.Find(".web-result")

	for i := range sel.Nodes {
		// Break loop once required amount of results are add
		if maxResult == len(results) {
			break
		}
		node := sel.Eq(i)
		titleNode := node.Find(".result__a")

		info := node.Find(".result__snippet").Text()
		title := titleNode.Text()
		ref := ""

		if len(titleNode.Nodes) > 0 && len(titleNode.Nodes[0].Attr) > 2 {
			ref = getDDGUrl(titleNode.Nodes[0].Attr[2].Val)
		}

		results = append(results[:], Result{title, info, ref})

	}

	return results, nil
}

func getDDGUrl(urlStr string) string {
	trimmed := strings.TrimPrefix(urlStr, "//duckduckgo.com/l/?uddg=")
	if idx := strings.Index(trimmed, "&rut="); idx != -1 {
		decodedStr, err := url.PathUnescape(trimmed[:idx])
		if err != nil {
			return ""
		}

		return decodedStr
	}

	return ""
}
