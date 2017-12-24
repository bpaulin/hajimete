package main

func main() {
	a := App{}
	a.Initialize(
		"couchbase://localhost",
		"couchbase",
		"couchbase",
		"hajimete")
	a.Run(":8080")
}
