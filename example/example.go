package main

import (
	"log"
	"os"

	"github.com/majiru/aqtk"
)

func main() {
	s := "こにちわ"
	v := aqtk.NewVoice(aqtk.F1)
	v.SetSpeed(120)
	out, err := v.Talk(s)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("./out.wav")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	f.Write(out)
}
