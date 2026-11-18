package main

func calcularPlan(cantidad int) float64 {

	total := float64(cantidad * 8)

	if total > 200 {
		total = 200
	}

	return total
}
