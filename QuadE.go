package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if y <= 0 || x <= 0 {
		return
	}

	for a := 1; a <= y; a++ {
		for b := 1; b <= x; b++ {
			if a == 1 && b == 1 {
				z01.PrintRune('A')
			} else if a == y && b == 1 {
				z01.PrintRune('C')
			} else if a == 1 && b == x {
				z01.PrintRune('C')
			} else if a == y && b == x {
				z01.PrintRune('A')
			} else if b == 1 && (a != y || a != 1) {
				z01.PrintRune('B')
			} else if b == x && (a != 1 || a != y) {
				z01.PrintRune('B')
			} else if a == 1 && (b != 1 || b != x) {
				z01.PrintRune('B')
			} else if a == y && (b != 1 || b != x) {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
