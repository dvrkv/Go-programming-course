/* Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
Функция должна называться task(). */

func task(c chan int, N int) {
	c <- N + 1
}