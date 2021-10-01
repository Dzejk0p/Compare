package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Dzejk0P/Compare/diff"
)

type daneZamowienia struct {
	StatusTransakcjiID PoleInt `diff:"statusTransakcjiId"`
	UzytkownikID       PoleInt `diff:"uzytkownikId"`
	// UzytkownikIDCC     int       `diff:"uzytkownikIdCC"`
	// DataZamowienia     time.Time `diff:"dataZamowienia"`
}

type daneKlientow struct {
	Klienci []klient `diff:"klienci"`
}

type klient struct {
	A PoleInt `diff:"a"`
}

type PoleInt struct {
	Bylo int `json:"bylo"`
	Jest int `json:"jest" diff:""`
}

func main() {
	// a := daneZamowienia{"3", 2, 1, time.Now()}
	// b := daneZamowienia{"2", 3, 1, time.Now().Add(1 * time.Hour)}

	// a := daneZamowienia{
	// 	StatusTransakcjiID: PoleInt{0, 1},
	// 	UzytkownikID:       PoleInt{0, 1},
	// }

	// b := daneZamowienia{
	// 	StatusTransakcjiID: PoleInt{0, 2},
	// 	UzytkownikID:       PoleInt{0, 2},
	// }

	a := daneKlientow{
		Klienci: []klient{
			{PoleInt{0, 2}},
		},
	}

	b := daneKlientow{
		Klienci: []klient{
			{PoleInt{0, 1}},
		},
	}

	z, _, err := diff.Diff(a, b)
	if err != nil {
		log.Fatal(err)
	}

	file, _ := json.MarshalIndent(z, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)

	fmt.Printf("%+v", z)
}
