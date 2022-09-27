package main

func main() {

	//Create Gin Router with Handlers attached
	r := httprouter.New()

	//Run Gin server
	r.Run()
}
