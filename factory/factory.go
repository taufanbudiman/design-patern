package factory

import (
	"fmt"
	"time"
)

type Content interface{
	Play()
}

type CloudContent struct {}

func (c *CloudContent) Play() {
	fmt.Println("this is a cloud content about kubernetes")
}

type CodingContent struct {}

func (cc *CodingContent) Play() {
	fmt.Println("this is coding content")
}

type Automotive struct {}

func (g *Automotive) Play() {
	fmt.Println("this is a content about Automotive")
}

type ContentCreator interface {
	Produce(time time.Time) Content
}

type Imre struct {}

func (i *Imre) Produce(time time.Time) Content {
	// lakukan pengecekan terhadap hari
	if time.Weekday().String() == "Tuesday"{
		return &CloudContent{}
	} else if time.Weekday().String() == "Thursday" {
		return &CodingContent{}
	} else {
		return nil
	}
}

type FitraEri struct {}

func (f *FitraEri) Produce(time time.Time) Content {
	return &Automotive{}
}