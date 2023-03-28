package main

import "net/http"

type Person struct {
	fName string
}

func (p *Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("First Name: " + p.fName))
}
func main() {
	person1 := &Person {
		fName: "Jim",
	}
	http.ListenAndServe(":8080", person1)
}