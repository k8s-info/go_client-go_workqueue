package main

// https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html

import "fmt"

func main() {
	for {
		desired := getDesiredState()
		current := getCurrentState()
		makechanges(desired, current)
	}
}
