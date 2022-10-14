/* Требуется вычислить период колебаний (t) математического маятника (мы округлили некоторые значения для удобства проверки), для этого нужно найти циклическую частоту колебания пружинного маятника (w), в формуле w встречается масса которую также нужно найти, все нужные формулы приведены ниже:


ВАЖНО! Считайте, что пакет main уже объявлен, а также функция main() с вызовом ВАШЕЙ будущей функции T() уже есть. Несмотря на то, что тестирование будет через ввод-вывод, вам НЕ требуется вводить и выводить что-либо. Для подсчета используйте УЖЕ ВВЕДЕННЫЕ ГЛОБАЛЬНЫЕ переменные k,p,v ТИПА float64 !!!

Sample Input:

1296 6 6
Sample Output:

1 */

func T() float64 {
	return 6 / W() //6/6 = 1
}

func W() float64 {
	return math.Sqrt(k / M()) //6
}

func M() float64 {
	return p * v //6*6 = 36
}