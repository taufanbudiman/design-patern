package main

import (
	"design-pattern/factory"
	"time"
)

func main() {
	var contentCreator factory.ContentCreator
	contentCreator = &factory.Imre{}

	content := contentCreator.Produce(time.Now())
	content.Play()
}
