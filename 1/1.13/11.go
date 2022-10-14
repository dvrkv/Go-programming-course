/* По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные

Дано число n (0<n<100).

Выходные данные

Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

Sample Input:

10
Sample Output:

10 korov */

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Printf("%d korova", n)
	} else if n >= 2 && n <= 4 {
		fmt.Printf("%d korovy", n)
	} else if n >= 5 && n <= 20 {
		fmt.Printf("%d korov", n)
	}

	if n >= 21 && n < 100 {
		a := n % 10 / 1
		if a == 1 {
			fmt.Printf("%d korova", n)
		} else if a >= 2 && a <= 4 {
			fmt.Printf("%d korovy", n)
		} else {
			fmt.Printf("%d korov", n)
		}
	}
}
