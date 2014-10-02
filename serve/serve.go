package serve

import (
	"encoding/csv"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)
type translations map[string]template.HTML;

func readLang(folder,file string) translations {
	inFile, err := os.Open("data/"+folder+"/lang/" + file + ".csv")
	defer inFile.Close()
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(inFile)
	reader.Comma = ':'
	m := make(translations)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		m[record[0]] = template.HTML(record[1])
	}
	return m
}
func loadTranslations(folder string) (translations,translations,translations) {
	langEn := readLang(folder,"en")
	langRu := readLang(folder,"ru")
	langUa := readLang(folder,"ua")
	return langEn,langRu,langUa
}

func handler(folder string) func(http.ResponseWriter, *http.Request) {
	handler := http.FileServer(http.Dir("public"))
	templates := template.Must(template.ParseFiles("data/"+folder+"/index.html"))
	/*
	langEn := readLang(folder,"en")
	langRu := readLang(folder,"ru")
	langUa := readLang(folder,"ua")
	*/
	langEn, langRu, langUa := loadTranslations(folder)

	return func(w http.ResponseWriter, r *http.Request) {
		var lang map[string]template.HTML
//
		langEn, langRu, langUa = loadTranslations(folder)
		if strings.HasSuffix(r.URL.Path,"/") {
			// todo: detect lang
			lang = langEn
		}
		if strings.HasSuffix(r.URL.Path,"/en") {
			lang = langEn
		}
		if strings.HasSuffix(r.URL.Path,"/ru") {
			lang = langRu
		}
		if strings.HasSuffix(r.URL.Path,"/ua") {
			lang = langUa
		}
		if lang != nil {
			templates.ExecuteTemplate(w, "index.html", lang)
			return
		}
		if strings.HasSuffix(r.URL.Path, ".svg") {
			w.Header().Set("Content-Type", "image/svg+xml")
		}
		handler.ServeHTTP(w, r)
	}
}

func init() {	
	http.HandleFunc("/1/", handler("1"))
	http.HandleFunc("/", handler("default"))

}
