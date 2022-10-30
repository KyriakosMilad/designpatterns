package main

func main() {
	p := NewPost()

	NewUser(p, "pola")
	NewUser(p, "henawy")
	NewUser(p, "jimmy")
	NewUser(p, "geowergy")
	NewUser(p, "kyriakos")

	p.UpdateLikes(1)
}
