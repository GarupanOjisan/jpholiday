package main

import (
	"bufio"
	"encoding/csv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
)

const (
	url     = "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
	outPath = "syukujitsu.csv"

	holidayTemplate = `package japan_holiday

var holidays = map[string]string{
	{{- range $date, $name := . }}
		"{{ $date }}": "{{ $name }}",
	{{- end }}
}
`
)

// 祝日データをダウンロードして、祝日のmapを生成する
func main() {
	// Download the CSV csvFile
	err := downloadCSVFile(url, outPath)
	if err != nil {
		panic(err)
	}

	csvFile, err := os.Open(outPath)
	if err != nil {
		panic(err)
	}
	defer func() {
		os.Remove(outPath)
		csvFile.Close()
	}()

	reader := csv.NewReader(transform.NewReader(bufio.NewReader(csvFile), japanese.ShiftJIS.NewDecoder()))

	// Create a map to store holidays
	var holidays = make(map[string]string)

	// Read each record from CSV
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// Convert the date format from 2021/01/01 to 2021-01-01
		date := strings.ReplaceAll(record[0], "/", "-")

		// Assuming the CSV format is date,holiday_name
		holidayName := record[1]
		holidays[date] = holidayName
	}

	tmpl, err := template.New("holidays").Parse(holidayTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(os.Stdout, holidays); err != nil {
		panic(err)
	}
}

func downloadCSVFile(url string, outPath string) error {
	// HTTP GETリクエストを発行してファイルをダウンロード
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 出力ファイルを開く（存在しない場合は作成）
	file, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// ダウンロードした内容をファイルに書き込む
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	return nil
}
