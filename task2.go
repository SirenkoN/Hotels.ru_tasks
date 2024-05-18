/*Задача 2. Написать функцию/метод, которая на вход получает массив положительных целых чисел произвольной длины.
Например [42, 12, 18]. На выходе возвращает массив чисел, которые являются общими делителями для всех указанных числе.
В примере это будет [2, 3, 6].*/

func commonDivisors(numbers []int) []int {
	var cd []int

	for i := 2; i <= 10; i++ {
		tempFlag := 0

		for i2 := 0; i2 < len(numbers); i2++ {
			if numbers[i2]%i == 0 {
				tempFlag = 1
			} else {
				tempFlag = 0
				break
			}
		}

		if tempFlag == 1 {
			cd = append(cd, i)
		}
	}
	return cd
}
