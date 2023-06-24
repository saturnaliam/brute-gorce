package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	loops := flag.Int("l", 1, "Number of loops.")
	show_elapsed := flag.Bool("t", false, "Show elapsed time.")
	flag.Parse()

	start := time.Now()
	generate(*loops)
	duration := time.Since(start)

	if *show_elapsed {
		fmt.Printf("elapsed time: %v", duration)
	}
}

func generate(loops int) {
	loops = loops - 1

	sub_gen := sub_gen(loops)

	for _, str := range sub_gen {
		fmt.Println(str)
	}
}

func sub_gen(loops int) []string {
	symbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-="
	var gen []string

	for i := 0; i < len(symbols); i++ {
		if loops > 0 {
			sub_gens := sub_gen(loops - 1)

			for _, char := range sub_gens {
				gen = append(gen, string(symbols[i]))
				gen[len(gen)-1] = fmt.Sprintf("%s%s", gen[len(gen)-1], string(char))
			}
		} else {
			gen = append(gen, string(symbols[i]))
		}
	}

	return gen
}
