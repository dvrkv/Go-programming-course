/* На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:

ihgewlqlkot
Sample Output:

hello */

package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string   //инпут
	fmt.Scanln(&a) //считываем инпут

	var b strings.Builder //Builder используется для создания строки с использованием методов Write
	//WriteRune добавляет кодировку UTF-8 кодовой точки Unicode r к буферу b
	for i, r := range a {
		if i%2 != 0 {
			b.WriteRune(r)
		}
	}
	fmt.Println(b.String())
}
