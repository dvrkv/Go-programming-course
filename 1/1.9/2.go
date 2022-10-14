/* По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:

237
Sample Output 1:

YES
Sample Input 2:

117
Sample Output 2:

NO */

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	var b, c, d int
	b = a / 100        //определяем первую цифру
	c = (a % 100) / 10 //определяем вторую цифру
	d = a % 10         //определяем третью цифру
	if b != c && c != d && d != b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
