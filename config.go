package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
"math/rand"
	"net/http"
	"math"
	"log"
	"os"
)

//Config is a singleton configuration holder
var Config Configuration

//Configuration holds config data
type Configuration struct {
	Level int64
}

var PageStatus []int

//LoadConfig parses file to Configuration struct
func LoadConfig() {
//	source, err := ioutil.ReadFile("/home/bajra/go-yml/src/github.com/degendra/go-web/config.yml")
	source, err := ioutil.ReadFile("/root/config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &Config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", Config.Level)
}



var Status = [45]int{
100,
101,
200,
201,
202,
203,
204,
205,
206,
300,
301,
302,
303,
304,
305,
307,
400,
401,
402,
403,
404,
405,
406,
407,
408,
409,
410,
411,
412,
413,
414,
415,
416,
417,
418,
500,
501,
502,
503,
504,
505,
428,
429,
431,
511,
}

//var PageStatus := [45]int

//GenerateRandomPageState
func GenerateRandomPageState(){
	dest := make([]int, len(Status))
	perm := rand.Perm(len(Status))
	for i, v := range perm {
		dest[v] = Status[i]
	}
	PageStatus = dest
	savePageStatus()
}

func savePageStatus(){
	for k, v:= range PageStatus{
		if int(math.Mod(float64(k), float64(Config.Level)))==0{
			log.Println()
			log.Println("Level: ", int(k/int(Config.Level))+1)
			log.Println("================================")
		}
		log.Println("Link status ", v, http.StatusText(v))

		if k>=int(Config.Level*Config.Level-1){
			break
		}

	}
}


func SetupLogging(){
	f, err := os.OpenFile("go-web-log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	//	defer f.Close()
	log.SetOutput(f)

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[go-web] ")
}