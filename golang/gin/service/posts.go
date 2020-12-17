package service

import "example.com/gin/modal"

// PostService collection
type PostService interface {
	Save(modal.Post) modal.Post
	FindAll() []modal.Post
}

type postService struct {
	posts []modal.Post
}

// New func
func New() PostService {
	return &postService{}
}

func (p *postService) Save(post modal.Post) modal.Post {
	p.posts = append(p.posts, post)
	return post

}

func (p *postService) FindAll() []modal.Post {
	return p.posts
}
