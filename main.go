package main

import "fmt"

func main() {
	fmt.Printf("done")
}

// curl --unix-socket /var/run/docker.sock -H "Content-Type: application/json" -X GET http://localhost/containers/7977a68779ef93cb6a98a0cfbdf320764ff1a34a2f6836243249a71d518188c5/stats\?one-shot\=true\&stream\=false
