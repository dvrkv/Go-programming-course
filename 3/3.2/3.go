/* Для решения данной задачи вам понадобится пакет strconv, возможно использовать пакеты strings или encoding/csv, или даже bufio - вы не ограничены в выборе способа решения задачи. В решениях мы поделимся своими способами решения этой задачи, предлагаем вам сделать то же самое.

В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в виде пробела, кроме того в России целая часть от дробной отделяется запятой. Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя используется символ ";".

На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления первого числа на второе с точностью до четырех знаков после "запятой" (на самом деле после точки, результат не требуется приводить к исходному формату).

P.S. небольшое отступление, связанное с чтением из стандартного ввода. Кто-то захочет использовать для этого пакет bufio.Reader. Это вполне допустимый вариант, но если вы резонно обрабатываете ошибку метода ReadString('\n'), то получаете ошибку EOF, а точнее (io.EOF - End Of File). На самом деле это не ошибка, а состояние, означающее, что файл (а os.Stdin является файлом) прочитан до конца. Чтобы ошибка была обработана правильно, вы можете поступить так:

if err != nil && err != io.EOF {
	...
}
Sample Input:

1 149,6088607594936;1 179,0666666666666
Sample Output:

0.9750 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cleaner(a string) (string, string) {

	p := strings.Split(strings.Replace(strings.Replace(a, ",", ".", -1), " ", "", -1), ";")
	numb1 := p[0]
	numb2 := p[1]
	trans(numb1, numb2)
	return numb1, numb2
}

func trans(c, d string) float64 {
	x, _ := strconv.ParseFloat(c, 64)
	y, _ := strconv.ParseFloat(d, 64)
	result := x / y
	fmt.Printf("%.4f", result)
	return x / y
}
func main() {

	//считываем всю строку с консоли
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)

	cleaner(input)
}
