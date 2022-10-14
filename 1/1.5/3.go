/*
По данному целому числу, найдите его квадрат.

Формат входных данных
На вход дается одно целое число.

Формат выходных данных
Программа должна вывести квадрат данного числа.

Sample Input:

3
Sample Output:

9
*/

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	a *= a
	fmt.Println(a)
}
