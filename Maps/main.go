package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "google.com",
		"Amazon": "amazon.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

}
