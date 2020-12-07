package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type myGetResp struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Company  struct {
		Name string `json:"name"`
	} `json:"company"`
}

type putBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// GET Method
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(string(b))
	mygresp := &myGetResp{}
	if err := json.Unmarshal(b, mygresp); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Header["Cache-Control"][0])
	fmt.Println(mygresp.Name, mygresp.Company.Name)

	// POST Method

	postRes, err := http.Post(
		"http://dummy.restapiexample.com/api/v1/create",
		"application/json; charset=UTF-8",
		resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	defer postRes.Body.Close()
	log.Println(string(postRes.Status))
	if postRes.StatusCode == 200 {
		b, err := ioutil.ReadAll(postRes.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(b))
	}

	// PUT/DELETE Method
	// These are not available directly

	user := putBody{
		Name:  "Parth",
		Email: "wparth777@gmail.com",
	}

	data, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(http.MethodPut, "http://dummy.restapiexample.com/api/v1/update/21", bytes.NewBuffer(data))
	if err != nil {
		log.Println(req)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	putResp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer putResp.Body.Close()

	log.Println(putResp.Status)

}
