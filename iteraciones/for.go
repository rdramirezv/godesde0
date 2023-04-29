package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 10; i > 1; i-- {
		fmt.Println(i)
	}

	for i := 10; i > 1; i-- {
		if i == 6 {
			break
		}

		fmt.Println(i)
	}

	for i := 10; i > 1; i-- {
		if i == 6 {
			continue
		}

		fmt.Println(i)
	}

	for i := 0; i < 100; i += 5 {
		fmt.Println(i)
	}

	for i := 100; i > 10; i -= 5 {
		fmt.Println(i)
	}
}
