package main

import (
    "fmt"
    "log"
    "net/http"
	"os"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(responseWriter, "Hi there, I love %s!", request.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	   p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
       p1.save()

	   p2 := &Page{Title: "Daniel", Body: []byte("This is a Daniel's Page.")}
       p2.save()
    //http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	fmt.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}