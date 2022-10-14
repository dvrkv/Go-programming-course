/* На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

Sample Input:

13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:

24h0m0s */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//считываем всю строку с консоли
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)

	//разбиваем строку на две части для преобразования в дату
	a := strings.Split(input, ",")[0]
	b := strings.Split(input, ",")[1]

	//конвертируем строки в даты
	date_a, _ := time.Parse("02.01.2006 15:04:05", a)
	date_b, _ := time.Parse("02.01.2006 15:04:05", b)

	//высчитываем разницу во времени
	if date_a.After(date_b) {
		fmt.Println(date_a.Sub(date_b))
	} else {
		fmt.Println(date_b.Sub(date_a))
	}

}
