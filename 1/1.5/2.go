/*
Петя торопился в школу и неправильно написал программу, которая сначала находит квадраты двух чисел, а затем их суммирует. Исправьте его программу.

Sample Input:

2
2
Sample Output:

8
*/

package main

import "fmt"

func main() {
	var a, b int
	var s int = 0
	fmt.Scan(&a, &b)

	a *= a
	b *= b
	s = a + b

	fmt.Println(s)
}
