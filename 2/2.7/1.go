/* На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

Sample Input:

6 8
Sample Output:

10 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)

	c := math.Sqrt(a*a + b*b)
	fmt.Println(c)
}
