/* Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:

843
Sample Output:

348 */

package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	var a int = i % 10 / 1
	var b int = i % 100 / 10
	var c int = i % 1000 / 100
	fmt.Printf("%d%d%d", a, b, c)
}
