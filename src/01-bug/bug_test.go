package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemote(t *testing.T) {
	render("file.html")
	t.Error("Error occured")
}

func TestRecorder(t *testing.T) {
	// Manipulating fake request and respnse
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		http.Error(rw, "something failed", http.StatusInternalServerError)
	})

	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err.Error())
	}
	w := httptest.NewRecorder()
	handler(w, req)

	fmt.Printf("%d & %s", w.Code, w.Result().Body)

}
func TestBug(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello World !!!")
	}))
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Errorf(err.Error())
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("%s", greeting)

}

func TestTemplate(t *testing.T) {
	templater()
}


func TestDB(t *testing.T){
	usingGORM();
}