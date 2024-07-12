package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	Creationdate time.Time `json:"creation_date"`
}

/*
	 d9736711---> {
					ID: "d9736711",
					OriginalURL: "https/github.com/rohit6294",
					ShortURL: "d936711",
					CreationDate: time.Now()
	}
*/
var urlDB = make(map[string]URL)

func generateshorturl(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("hasher: ", hasher)
	data := hasher.Sum(nil)
	fmt.Println("slice of data  ", data)
	hash := hex.EncodeToString(data)
	fmt.Println("After encode  ", hash)
	fmt.Println("After encode upto 8  ", hash[:8])
	return hash[:8]
}
func createURL(originalURL string) string{
	shorturl := generateshorturl(originalURL)
	id := shorturl
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shorturl,
		Creationdate: time.Now(),
	}
	return shorturl
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url,nil
}

func handler(w http.ResponseWriter,r *http.Request){
	http.ServeFile(w,r,"index.html")
}


func ShortURLHandler(w http.ResponseWriter,r *http.Request ){
	var data struct{
		URL string `json:"url"`
	}
	err:=json.NewDecoder(r.Body).Decode(&data)
	if err!=nil{
		http.Error(w,"Invalid request body", http.StatusBadRequest)
		return 
	}

	shortURL_:=createURL(data.URL)
	// fmt.Fprintf(w,shortURL)
	response:= struct{
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLhandler(w http.ResponseWriter,r *http.Request){
	id:=r.URL.Path[len("/redirect/"):]

	url,err:=getURL(id)
	if err!=nil{
		http.Error(w,"Invalid Request",http.StatusNotFound)
	}
	http.Redirect(w,r,url.OriginalURL,http.StatusFound)
}
func main() {
	fmt.Println("url shortner")

	// OriginalURL := "https://www.youtube.com/watch?v=dVVJU-3eU1g&list=PLzjZaW71kMwSEVpdbHPr0nPo5zdzbDulm&index=34"
	// generateshorturl(OriginalURL)

	// Register the handler function to handle all req to the root url("/")
	http.HandleFunc("/",handler)
	http.HandleFunc("/shorten",ShortURLHandler)
	http.HandleFunc("/redirect/",redirectURLhandler)

	// start http server on port
	fmt.Println("Starting servere on port 3000.......")
	err:= http.ListenAndServe(":3000",nil)
	if err !=nil{
		fmt.Println("Erro on starting server: ",err)
	}
}
