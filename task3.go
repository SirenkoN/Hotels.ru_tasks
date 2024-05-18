/*Задача 3. Написать функцию/метод, которая возвращает массив простых чисел в диапазоне
(2 числа - минимальное и максимальное) заданных чисел.
Например, на вход переданы 2 числа: от 11 до 20  (диапазон считается включая граничные значения).
На выходе программа должна выдать [11, 13 , 17, 19]
*/

func primeNumbers(min int, max int) []int {
	var primeNumbersSlice []int

	for i := min; i <= max; i++ {
		tempFlag := 0

		for i2 := 2; i2 <= i/2; i2++ {
			if i%i2 == 0 {
				tempFlag = 1

				break
			}
		}

		if tempFlag == 0 {
			primeNumbersSlice = append(primeNumbersSlice, i)
		}
	}
	return primeNumbersSlice
}
