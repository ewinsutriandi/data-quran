package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/go-shiori/dom"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

var (
	rxDateTime = regexp.MustCompile(`(?i)\w{3}, \d{1,2} \w{3} \d{4} \d{2}:\d{2}:\d{2} [+-]\d{4}`)
)

func tanzilTransCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tanzil-trans <dst-dir>",
		Short: "Download Quran translation from Tanzil.net",
		Args:  cobra.MaximumNArgs(1),
		Run:   tanzilTransCmdHandler,
	}
}

func tanzilTransCmdHandler(cmd *cobra.Command, args []string) {
	// Get arguments
	dstDir := args[0]
	if dstDir == "" {
		dstDir = "."
	}

	// Download translation list
	logrus.Println("fetching translation list")
	translationList, err := getTranslationList()
	if err != nil {
		logrus.Fatalln(err)
	}

	// Download last update time for each translation
	for i, item := range translationList {
		logrus.Println("fetching update date for", item.ID)
		lastUpdate, err := getTranslationUpdateTime(item.ID)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		item.LastUpdate = lastUpdate.Format("2006-01-02 15:04:05")
		translationList[i] = item
	}

	// Download translation and convert it to JSON
	for _, item := range translationList {
		logrus.Println("fetching translation for", item.ID)
		texts, err := getTranlationText(item.ID)
		if err != nil {
			logrus.Errorln(err)
			continue
		}

		transData := struct {
			Translations []string       `json:"translations"`
			Footnotes    map[int]string `json:"footnotes"`
		}{texts, nil}

		dstPath := filepath.Join(dstDir, "tanzil-"+item.ID+".json")
		err = saveToJSON(&transData, dstPath)
		if err != nil {
			logrus.Errorln(err)
			continue
		}
	}

	// Save translation list
	dstPath := filepath.Join(dstDir, "00-tanzil-translation-list.json")
	err = saveToJSON(&translationList, dstPath)
	if err != nil {
		logrus.Fatalln(err)
	}
}

func getTranslationList() ([]TanzilTranslationData, error) {
	// Download translation page
	resp, err := downloadFile("http://tanzil.net/trans")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse page
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	// Find transition list
	table := dom.QuerySelector(doc, "table.transList")
	if table == nil {
		return nil, fmt.Errorf("translation list not found")
	}

	translationList := []TanzilTranslationData{}
	tableRows := dom.QuerySelectorAll(table, "tr")
	for _, tr := range tableRows {
		tds := dom.QuerySelectorAll(tr, "td")
		if len(tds) < 4 {
			continue
		}

		downloadLink := dom.QuerySelector(tds[3], "a.download")
		if downloadLink == nil {
			continue
		}
		downloadURL := dom.GetAttribute(downloadLink, "href")

		id := path.Base(downloadURL)
		language := dom.TextContent(tds[0])
		name := dom.TextContent(tds[1])

		translator := dom.TextContent(tds[2])
		translator = strings.ReplaceAll(translator, "*", "")
		translator = strings.ReplaceAll(translator, "†", "")

		translationList = append(translationList, TanzilTranslationData{
			ID:         id,
			Language:   strings.TrimSpace(language),
			Name:       strings.TrimSpace(name),
			Translator: strings.TrimSpace(translator),
		})
	}

	// Sort translation list
	sort.Slice(translationList, func(i, j int) bool {
		transI := translationList[i]
		transJ := translationList[j]

		if transI.Language != transJ.Language {
			return transI.Language < transJ.Language
		}

		return transI.ID < transJ.ID
	})

	return translationList, nil
}

func getTranslationUpdateTime(id string) (time.Time, error) {
	// Download translation changelog
	url := "http://tanzil.ca/trans/log/" + id
	resp, err := downloadFile(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	// Parse page
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return time.Time{}, err
	}

	// Find last update
	dateNode := dom.QuerySelector(doc, "div.date")
	if dateNode == nil {
		return time.Time{}, fmt.Errorf("last change not found")
	}

	strDate := dom.TextContent(dateNode)
	strDate = rxDateTime.FindString(strDate)
	date, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", strDate)
	if err != nil {
		return time.Time{}, err
	}

	return date.UTC(), nil
}

func getTranlationText(id string) ([]string, error) {
	// Download translation
	url := "http://tanzil.net/trans?type=txt&transID=" + id
	resp, err := downloadFile(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	texts := []string{}
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)

		if text == "" || strings.HasPrefix(text, "#") {
			continue
		}

		texts = append(texts, text)
	}

	if len(texts) != 6236 {
		return nil, fmt.Errorf("translation is not complete")
	}

	return texts, nil
}

func saveToJSON(data interface{}, dstPath string) error {
	bt, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
	return ioutil.WriteFile(dstPath, bt, os.ModePerm)
}

func downloadFile(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
