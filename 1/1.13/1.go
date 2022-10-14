/* Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:

745
Sample Output:

16 */

package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)

	var a int = i % 1000 / 100
	var b int = i % 100 / 10
	var c int = i % 10 / 1
	fmt.Println(a + b + c)
}
