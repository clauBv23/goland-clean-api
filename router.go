package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func addPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	res.WriteHeader(http.StatusOK)

	result, _ := json.Marshal(post)
	res.Write(result)
}

func addUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	//userUseCaseInputAdapter, err := user_input_adapter.NewUserInputAdapter(user_input_adapter.UserInputAdapterOptions{UserUseCase: &user_use_case.DefaultUserUseCase{}})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	userUseCaseInputAdapter.CreateUser(user)
	res.WriteHeader(http.StatusOK)

	result, _ := json.Marshal(post)
	res.Write(result)
}
