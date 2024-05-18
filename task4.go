/*Задача 4. Написать метод, который в консоль выводит таблицу умножения. На вход метод получает число,
до которого выводит таблицу умножения. В консоли должна появиться таблица. Например, если на вход пришло число 5, то получим:*/

type Number struct {
	n int
}

func (n Number) printMultiplicationTable(number int) {
	if number > 0 {
		for i := 1; i <= number; i++ {
			for i2 := 1; i2 <= number; i2++ {
				fmt.Printf("%d \t", i2*i)
			}

			fmt.Println()
		}
	}
}