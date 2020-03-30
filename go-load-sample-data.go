package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Privet")

	content, err := ioutil.ReadFile("sample-data.json")

	if err != nil {
		log.Fatal(err)
	}

	split := strings.Split(string(content), "\n")

	for i, s := range split {
		if len(s) < 10 {
			continue
		}

		post, err := http.Post("http://localhost:9200/vensi/_doc/"+strconv.Itoa(i), "application/json", strings.NewReader(s))

		if err != nil {
			log.Fatal(err)
		}

		defer post.Body.Close()
		all, err := ioutil.ReadAll(post.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(all))

		log.Printf("%d %d\n", i, len(s))
	}

}
