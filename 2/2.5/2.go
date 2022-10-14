/* На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input:

топот
Sample Output:

Палиндром */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int = 0
	var rev string //перевернутое слово

	//считываем всю строку с консоли
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	rs := make([]rune, len(text)) //руиним входные данные
	for _, i := range text {
		rs[n] = i
		n++
	}
	rs = rs[0:n]

	for a := 0; a < n/2; a++ {
		rs[a], rs[n-1-a] = rs[n-1-a], rs[a]
	}
	rev = string(rs)

	if text == rev {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}
