package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// you can interface the values in structs
// because by default Golang uses
// PascalCase naming convention

// use `json:"yourDataField"` in the struct fields
// to override the key of struct
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

type User struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var posts []Post = []Post{}

func AddPost(res http.ResponseWriter, req *http.Request) {
	// get json body
	var newPost Post
	json.NewDecoder(req.Body).Decode(&newPost)
	// append to slice
	posts = append(posts, newPost)
	// set response headers
	res.Header().Set("Content-type", "application/json")
	// return response
	json.NewEncoder(res).Encode(posts)
}

func GetAllPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	// return response
	json.NewEncoder(res).Encode(posts)
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	// get the ID param
	var idParam string = mux.Vars(req)["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		res.WriteHeader(400)
		res.Write([]byte("ID could not be converted into integer"))
		return
	}

	if id >= len(posts) {
		res.WriteHeader(404)
		res.Write([]byte("Not found!"))
		return
	}

	post := posts[id]
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(post)
}

func UpdatePost(res http.ResponseWriter, req *http.Request) {
	// get the ID param
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		//change status code
		res.WriteHeader(400)
		res.Write([]byte("ID could not be converted into integer"))
		return
	}

	if id >= len(posts) {
		//change status code
		res.WriteHeader(404)
		res.Write([]byte("Not found!"))
		return
	}

	// get json body
	var updatedPost Post
	json.NewDecoder(req.Body).Decode(&updatedPost)

	posts[id] = updatedPost
	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(updatedPost)
}

func DeletePost(res http.ResponseWriter, req *http.Request) {
	// get the ID param
	var idParam string = mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		//change status code
		res.WriteHeader(400)
		res.Write([]byte("ID could not be converted into integer"))
		return
	}

	if id >= len(posts) {
		//change status code
		res.WriteHeader(404)
		res.Write([]byte("Not found!"))
		return
	}

	posts = append(posts[:id], posts[id+1:]...)
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(200)
}
