package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"math"
)
//Index func handles default page
//It shows repeated contents defined by level in config.yml file
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, DynamicPage(int64(1), int64(1)))
}

//DynamicPage function is used for generation dynamic content in website.
func DynamicPage(level, count int64)string{
	levelLimit :=Config.Level
	strLevel:= strconv.FormatInt(level, 10)
	body := `<html><head><title></title></head><body><h1>Level `
	body+=strLevel+` of `
	body+=strconv.FormatInt(levelLimit, 10)+`</h1>`
	if level<levelLimit{
		for i:=int64(0); i<levelLimit; i++{
			body+=`<p>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.
			It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus	PageMaker including versions of Lorem Ipsum.<a href="http://tournepal.org:8080/pages/`
			body+=strconv.FormatInt(level+1, 10)
			body+=`/`+strconv.FormatInt(i+1, 10)
			body+=`">more...</a></p>`
			fmt.Printf("count:", i)
		}
	}else{
		body+="Termination page"
	}
	body+=`</body></html>`
	return body
}

func Page(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars==nil{
		w.WriteHeader(http.StatusNoContent)
	}
	level:=vars["level"]
	link:=vars["link"]
	intLevel, err := strconv.ParseInt(level, 10, 64)
	if err != nil {
		panic(err)
	}
	intLink, err := strconv.ParseInt(link, 10, 64)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(PageStatus[HttpStatus(intLevel, intLink)])
	fmt.Printf("Level: %d Link: %d", intLevel, intLink)
	w.Write([]byte(DynamicPage(intLevel, intLink)))
}

func HttpStatus(intLevel, intLink int64) int {
	res:=int(math.Mod(float64(int((intLevel-1)*Config.Level+(intLink-1))), 45.0))
	fmt.Println(res)
	return res
}

/*

func OpenEmail(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Println(err)
		return
	}
	go handleOpenEmail(body)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode("OK"); err != nil {
		log.Println(err)
	}
}
*/
