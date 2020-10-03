package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/colinmarc/hdfs"
)

type Vertex struct {
	X int
	Y int
}

func main() {

	dataDir := "/data"
	filename := "MobyDick.txt"

	// read file as bytes
	file, err := os.Open("..\\assets\\" + filename)
	panicOn(err)
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	panicOn(err)
	fmt.Print(b)

	// read file into string
	// file, err := ioutil.ReadFile("..\\assets\\MobyDick.txt")
	// text := string(file)
	// fmt.Println(text)

	// connect to namenode and init data load directory
	client, err := hdfs.New("localhost:9000")
	panicOn(err)

	// check load data directory
	info, err := client.Stat(dataDir)
	fmt.Println(info)
	// panic(err)

	if info == nil {
		// load data directory does not exist
		fmt.Println("load data directory does not exist")
		client.Mkdir(dataDir, 777)
	}

	fmt.Println(dataDir + "/" + filename)
	fmt.Println("-------")

	writer, err := client.Create(dataDir + "/E.txt")
	// panic(err)

	ret, err := writer.Write(b)
	panic(err)
	fmt.Println("writer return code: ", ret)

}

func panicOn(err error) {

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
