/* Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.

Sample Input:

1 5
Sample Output:

15 */

package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	var sum int = 0
	for i := A; i <= B && A <= 100 && B <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
