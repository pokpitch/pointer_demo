package main

func main() {
	num := 2
	double(&num)
}

func double(p *int) {
	*p = *p * 2
}
