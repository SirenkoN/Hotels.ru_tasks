/*Задача 1. Разработайте функцию, которая принимает целое число в качестве аргумента и возвращает строку, 
содержащую это число и слово "компьютер" в нужном склонении по падежам в зависимости от числа. Например, 
при вводе числа 25 функция должна возвращать "25 компьютеров", для числа 41 — "41 компьютер", а для числа 1048 — "1048 компьютеров".
*/

func outputComputers(number int) string {
    numberString := strconv.Itoa(number)

    for {
        if number == 0 {
            return numberString + " Компьютеров"
            break
        } else if number == 1 {
            return numberString + " Компьютер"
            break
        } else if number >= 2 && number <= 4 {
            return numberString + " Компьютера"
            break
        } else if number >= 5 && number <= 20 {
            return numberString + " Компьютеров"
            break
        } else if number >= 21 && number <= 99 {
            number = number % 10
        } else if number > 99 {
            number = number % 100
        }
    }
    return numberString
}
