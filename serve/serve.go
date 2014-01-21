package serve

import (
	//    "fmt"
	"encoding/csv"
	"html/template"
	"io"
	"net/http"
	"os"
  "strings"
)

func readLang(file string) map[string]string {
	inFile, err := os.Open("lang/"+file+".csv")
	defer inFile.Close()
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(inFile)
	reader.Comma = ':'
	m := make(map[string]string)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		m[record[0]] = record[1]
	}
	return m
}
func handler() func(http.ResponseWriter, *http.Request) {
	handler := http.FileServer(http.Dir("public"))
	templates := template.Must(template.ParseFiles("public/index.html"))
	langEn := readLang("en")
  langRu := readLang("ru")
  langUk := readLang("uk")
	return func(w http.ResponseWriter, r *http.Request) {
    var lang map[string]string
    if r.URL.Path == "/" {
      // todo: detect lang
      lang = langEn
    }
		if r.URL.Path == "/en" {
      lang = langEn			
		}
    if r.URL.Path == "/ru" {
      lang = langRu     
    }
    if r.URL.Path == "/uk" {
      lang = langUk     
    }
    if lang!=nil {
      templates.ExecuteTemplate(w, "index.html", lang)
      return
    }
    if strings.HasSuffix(r.URL.Path,".svg") {
      w.Header().Set("Content-Type", "image/svg+xml")  
    }
		handler.ServeHTTP(w, r)
	}
}

func init() {
	http.HandleFunc("/", handler())
}
