/* На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:

12 мин. 13 сек.
Sample Output:

Fri May 15 19:28:18 UTC 2020 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//Конвертируем в дату
	time_1 := time.Unix(1589570165, 0)

	//считываем всю строку с консоли
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)

	//приводим строку к нужному формату
	cleaner1 := strings.ReplaceAll(input, " мин. ", "m")
	cleaner2 := strings.ReplaceAll(cleaner1, " сек.", "")
	//дробим на отдельные части минуты и секунды
	spl := strings.Split(cleaner2, "m")
	min, _ := strconv.Atoi(spl[0])
	sec, _ := strconv.Atoi(spl[1])

	afterTime_1 := time_1.Add(time.Minute * time.Duration(min))
	afterTime_2 := afterTime_1.Add(time.Second * time.Duration(sec))

	//выводим дату в формате UnixDate
	fmt.Print(afterTime_2.Format(time.UnixDate))

}
