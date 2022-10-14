/* Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:

zaabcbd
Sample Output:

zcd */

package main

import (
	"fmt"
	"strings"
)

func main() {
	//считываем входные данные
	var input string
	fmt.Scanln(&input)

	//определяем повторяющиеся символы
	var b strings.Builder     //Builder используется для создания строки с использованием методов Write
	for _, i := range input { //i предоставляет кодировку Unicode для каждого символа
		if strings.Count(input, string(i)) == 1 { //задаем условие, чтобы кол-во подстрок в строке равнялось 1
			b.WriteRune(i) //WriteRune переводит Unicode i в символы к буферу b
		}
	}
	//выводим результат на печать
	fmt.Println(b.String())

}
