package main

import "fmt"

type User struct {
	post *Post
	name string
}

func NewUser(post *Post, name string) *User {
	u := &User{
		post: post,
		name: name,
	}

	u.post.AddObserver(u)

	return u
}

func (u *User) Update(likes uint) {
	fmt.Printf("user %v received notification that post likes updated to %v\n", u.name, likes)
}

type Post struct {
	likes     uint
	observers []*User
}

func NewPost() *Post {
	return &Post{
		likes:     0,
		observers: []*User{},
	}
}

func (p *Post) AddObserver(u *User) {
	p.observers = append(p.observers, u)
}

func (p *Post) UpdateLikes(likes uint) {
	p.likes = likes
	p.notify()
}

func (p *Post) GetLikes() uint {
	return p.likes
}

func (p *Post) notify() {
	for _, observer := range p.observers {
		observer.Update(p.likes)
	}
}
