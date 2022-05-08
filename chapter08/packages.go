package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"path/filepath"
	"errors"
	"container/list"
	"sync"
	"time"
)

func test_string() {
	// if string contains substring
	fmt.Println(strings.Contains("string", "ring")) // return bool
	// count number of substring in string
	fmt.Println(strings.Count("string", "t")) // return int
	// if string starts with substring
	fmt.Println(strings.HasPrefix("string", "st")) // return bool
	// if string ends with substring
	fmt.Println(strings.HasSuffix("string", "ng")) // return bool
	// index position of substring in string (-1 if not found)
	fmt.Println(strings.Index("string", "t")) // return int
	// join single string together
	fmt.Println(strings.Join([]string{"a", "b"}, "-")) // return string
	// repeat string
	fmt.Println(strings.Repeat("a", 5)) // return string
	// replace old with new (also take number of times, -1 for as many times as possible)
	fmt.Println(strings.Replace("yyeesss", "s", "z", 2)) // return string
	// split a string to list of strings
	fmt.Println(strings.Split("a-b-c-d-e-f-g", "-")) // return []string
	// convert string to lower
	fmt.Println(strings.ToLower("STRING")) // return string
	// convert string to upper
	fmt.Println(strings.ToUpper("string")) // return string
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(arr, str)
}

func open_file() {
	file, err := os.Open("./README.md")
	if err != nil {
		// handle error
		return 
	}
	defer file.Close() // to make sure file close as soon as possible
	// file size
	stat, err := file.Stat()
	// read file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	str := string(bs)
	fmt.Println(str)
}

func read_file() {
	bs, err := ioutil.ReadFile("./README.md")
	if err != nil {
		return
	}
	fmt.Println(string(bs))
}

func create_file() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("Go lang create file test")
}

func list_dir() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.ReadDir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func path_list_dir() {
	// it shows all folders and subfolders 
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	}) 
}

func new_error() {
	err := errors.New("New error message")
	fmt.Println(err)
}

func test_list() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	}
	fmt.Println(x)
}

func test_mutex() {
	m := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}
}

func main(){
	test_mutex()
	var input string
	fmt.Scanln(&input)
}
