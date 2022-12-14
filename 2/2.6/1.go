/* Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления, но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции. Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error). Пакет main уже объявлен, а нужные пакеты уже импортированы!

Не забудьте считать два числа которые необходимо поделить (передать в функцию)!

Sample Input:

10 5
Sample Output:

2 */

func main() {
	//считываем входные данные
	var aa, bb int
	fmt.Scan(&aa, &bb)

	/*возрвращаем из функции divide результат
	и информацию об ошибке*/
	g, error := divide(aa, bb)

	if error == nil {
		fmt.Println(g)
	} else {
		fmt.Println("ошибка")
	}
}