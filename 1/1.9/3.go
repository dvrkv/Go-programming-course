/* Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:

1234
Sample Output:

1 */

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a == 10000 {
		fmt.Println((a % 100000) / 10000)
	} else if a >= 1000 {
		fmt.Println((a % 10000) / 1000)
	} else if a >= 100 {
		fmt.Println((a % 1000) / 100)
	} else if a >= 10 {
		fmt.Println((a % 100) / 10)
	} else if a >= 1 {
		fmt.Println((a % 10) / 1)
	}
}
