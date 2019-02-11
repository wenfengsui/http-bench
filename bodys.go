package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func ParseSiteSkc(filename string) map[string][]string {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	ret := make(map[string][]string, 4)
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		temp := ret[line[0]]
		temp = append(temp, line[1])
		ret[line[0]] = temp
	}
	return ret
}

func ContructBody(src map[string][]string) []byte {
	size := 200
	sites := [4]string{"shein", "romwe"}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	rnd := r.Intn(len(sites))
	site := sites[rnd]
	skcs := src[site]
	// fmt.Println(len(skcs))
	slice := len(skcs) - size
	if slice > 0 {
		skcs, src[site] = skcs[slice:], skcs[:slice]
	} else {
		skcs = []string{}
	}
	ret := make(map[string]interface{})
	ret["site"] = site
	ret["skc_list"] = skcs
	jsonString, _ := json.Marshal(ret)
	return jsonString
}

func _main() {
	csv := ParseSiteSkc("c:/Users/wenfe/site_skc.csv")
	fmt.Println(len(csv["shein"]))
	fmt.Println(len(csv["romwe"]))
	ContructBody(csv)
	ContructBody(csv)
	ContructBody(csv)
	ContructBody(csv)
	ContructBody(csv)
	ContructBody(csv)
	ContructBody(csv)
	ContructBody(csv)
	ContructBody(csv)
	ContructBody(csv)
	ContructBody(csv)
}
