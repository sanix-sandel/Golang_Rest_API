package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type Book struct {
	ID     string `json:"id, omitempty"`
	Name   string `json:"name, omitempty"`
	Author string `json:"author, omitempty"`
}

type JSONServer struct {
}

//Rpc implementation
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	//Read Json file and load data
	absPath, _ := filepath.Abs("rpcServer/books.json")
	raw, readerr := ioutil.ReadFile(absPath)
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}

	//Unmarshall Json raw data into books array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}

	//Iterate over each book to find the given book
	for _, book := range books {
		if book.ID == args.ID {
			//if Book found, fill reply with it
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	//Create new Rpc server
	s := rpc.NewServer()
	//Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	//Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
