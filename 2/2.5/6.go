/* Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:

fdsghdfgjsdDD1
Sample Output:

Ok */

package main

import (
	"fmt"
	"regexp"
	"unicode"
)

//функция проверяет соблюдение требований к паролю
func ValidPassword(a string) bool {
	rs := []rune(a)
	if len(rs) < 5 {
		return false
	}

	for _, i := range rs {
		x := unicode.Is(unicode.Latin, i)
		y := unicode.Is(unicode.Digit, i)
		if x == true && y == true {
			return true
		}
	}
	return true
}

func main() {
	//считываем входные данные
	var input string
	fmt.Scanln(&input)
	k, _ := regexp.MatchString(`^[a-zA-Z0-9]{5,}$`, input)

	if ValidPassword(input) == true && k == true {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
