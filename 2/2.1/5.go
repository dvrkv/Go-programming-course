/* Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1

Sample Input:

Sample Output: */

func sumInt(a ...int) (int, int) {
	var n int = len(a) //кол-во
	var sum int        //сумма

	for _, elem := range a {
		sum += elem
	}
	return n, sum

}