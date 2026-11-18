package main

func obtenerPrecio(popularidad string) float64 {

	if popularidad == "Alta" {
		return 12
	}

	if popularidad == "Media" {
		return 8
	}

	return 5
}