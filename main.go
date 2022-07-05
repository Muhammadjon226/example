// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"
// 	"go.uber.org/zap"
// )

// func Upload(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("888888888888888888:", r.Method)
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	if r.Method == "GET" {
// 		t := template.Must(template.ParseFiles("index.html"))
// 		t.Execute(w, nil)
// 	} else if r.Method == "POST" {

// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Content-type", "multipart/form-data")

// 		// if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
// 		// 	ile", http.StatusBadRequest)
// 		// 	fmt.Println(err)
// 		// 	return
// 		// }

// 		// fmt.Println(body.Name)
// 		fmt.Println("body: ", r.Body)
// 		file, header, err := r.FormFile("file")
// 		if err != nil {
// 			fmt.Println("5555555555555", err)

// 			return
// 		}
// 		defer file.Close()

// 		// uuid, err := uuid.NewRandom()
// 		fmt.Println("form-file", header.Filename)

// 		out, err := os.Create("excel/" + header.Filename)
// 		if err != nil {

// 			return
// 		}
// 		_, err = io.Copy(out, file)
// 		defer out.Close()
// 		if err != nil {

// 			return
// 		}
// 		// return that we have successfully uploaded our file!
// 		fmt.Println("Successfully Uploaded File")
// 	}
// }

// func main() {
// 	log.Println("hello service!")

// 	godotenv.Load()

// 	port := os.Getenv("PORT")

// 	fmt.Println("Listen server :", port)
// 	port = "8000"

// 	r := mux.NewRouter()

// 	r.HandleFunc("/data/upload/", Upload).Methods("GET", "POST")

// 	if err := http.ListenAndServe(":"+port, r); err != nil {
// 		log.Println("failed to connect to server: ", zap.Error(err))
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
  var abbr string
  phrase := readString()
  for _, word := range strings.Fields(phrase) {
	sliceOfRunes := []rune(word)
    if unicode.IsLetter(sliceOfRunes[0]) {
      abbr += string(unicode.ToUpper(sliceOfRunes[0]))
    }
    // 1. Разбейте фразу на слова, используя `strings.Fields()`
    // 2. Возьмите первую букву каждого слова и приведите
    //    ее к верхнему регистру через `unicode.ToUpper()`
    // 3. Если слово начинается не с буквы, игнорируйте его
    //    проверяйте через `unicode.IsLetter()`
    // 4. Составьте слово из получившихся букв и запишите его
    //    в переменную `abbr`
    // ...
  }
  fmt.Println(abbr)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
  rdr := bufio.NewReader(os.Stdin)
  str, _ := rdr.ReadString('\n')
  return str
}