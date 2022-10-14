/* По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:

50
Sample Output:

1 2 4 8 16 32 */

package main

import "fmt"

func main() {
	var a uint
	fmt.Scan(&a)

	for i := uint(1); i <= a; i <<= 1 {

		if i <= a {
			fmt.Printf("%d ", i)
		}
	}
}
