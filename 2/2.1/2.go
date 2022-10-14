/* Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:

4 5 6 7
Sample Output:

4 */

func minimumFromFour() int {
	var a [4]int                  //массив
	var b int                     //элементы массива
	var minimumFromFour int = 100 //наименьший элемент массива

	for i := range a { //возвращаем копию элементов массива в бесконечном массиве
		fmt.Scan(&a[i]) //считываем элементы массива
	}

	for _, b = range a {
		if b < minimumFromFour {
			minimumFromFour = b
		}
	}

	return minimumFromFour
}