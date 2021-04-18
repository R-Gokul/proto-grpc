package main

import (
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello Gokul")

	gokul := &Person{
		Name: "Gokul",
		Age:  31,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1500,
		},
	}
	data, err := proto.Marshal(gokul)
	if err != nil {
		log.Fatal("Marshal error")
	}
	fmt.Println(data)

	newGokul := &Person{}

	err = proto.Unmarshal(data, newGokul)
	if err != nil {
		log.Fatal("Unmarshal Error", err)
	}

	fmt.Println(newGokul.GetName())
	fmt.Println(newGokul.GetAge())
	fmt.Println(newGokul.SocialFollowers.GetYoutube())
	fmt.Println(newGokul.SocialFollowers.GetTwitter())

}
