/* Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:

8
Sample Output:

6 */

package main

import (
	"fmt"
)

func main() {
	var n int
	var listFib []int

	_, _ = fmt.Scan(&n)

	listFib = append(listFib, 0, 0)

	var a, b int = 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		listFib = append(listFib, a)
	}

	k := check(listFib, n)
	fmt.Println(k)
}

func check(nums []int, n int) int {
	for k, v := range nums {
		if v == n {
			return k
		}
	}
	return -1
}
