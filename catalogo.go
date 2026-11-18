package main

import "fmt"

func mostrarCatalogo() {

	series := []string{
		"Breaking Bad",
		"Dark",
		"The Witcher",
		"Wednesday",
		"Stranger Things",
		"The Boys",
		"Loki",
		"Money Heist",
		"Peaky Blinders",
		"The Walking Dead",
	}

	fmt.Println("CATALOGO")

	for i, serie := range series {
		fmt.Printf("%d. %s\n", i+1, serie)
	}
}
