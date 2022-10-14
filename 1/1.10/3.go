/* Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8. Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:

5
38 24 800 8 16
Sample Output:

40 */

package main

import "fmt"

func main() {
	var a int       //кол-во чисел в последовательности
	var b int       //числа в последовательности
	var sum int = 0 //сумма двузначных чисел
	fmt.Scan(&a)    //считываем кол-во чисел

	for i := 0; i < a; i++ {
		fmt.Scan(&b) //считываем числа в последовательности
		if b%8 == 0 && b <= 99 && b >= 10 {
			sum += b
		}
	}
	fmt.Println(sum)
}