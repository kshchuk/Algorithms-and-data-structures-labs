package main

import (
	"fmt"
	"math/big"
	"strings"
)

/*
Функція main є вхідним пунктом в програму і містить код для
створення та взаємодії з хеш-таблицею.
*/
func main() {

	// Звернення до функції NewHashTable("numbers.txt") створює хеш-таблицю із заданого файлу "numbers.txt".
	ht, err := NewHashTable("numbers.txt")

	// Якщо сталася помилка при створенні хеш-таблиці, то виводиться повідомлення про помилку, та програма завершується.
	if err != nil {
		fmt.Print(err)
		return
	}

	// Виклик методу Print() для виведення всіх елементів, що зберігаються в хеш-таблиці.
	ht.Print()

	// Виконання циклу, доки користувач не вирішить вийти з програми.
	for {
		fmt.Println("1 - Print")
		fmt.Println("2 - Check")
		fmt.Println("0 - Exit")

		var ans int
		fmt.Scan(&ans) // Очікування вводу користувачем вибраної опції.

		switch ans {
		case 1: // Якщо вибрано опцію 1, викликається метод Print() для виведення всіх елементів, що зберігаються в хеш-таблиці.
			fmt.Println()
			ht.Print()
			break
		case 2: // Якщо вибрано опцію 2, користувач повинен ввести дріб у вигляді "a/b". Цей дріб переводиться в об'єкт типу big.Rat,
			//а потім викликається метод Check() для перевірки наявності дробу в хеш-таблиці. Результат перевірки виводиться на екран.
			fmt.Print(`Write Fraction ("a/b"): `)
			var frac string
			fmt.Scan(&frac)

			parts := strings.Split(frac, "/")
			if len(parts) != 2 {
				fmt.Errorf("invalid format: %s", frac)
			}
			num := new(big.Int)
			den := new(big.Int)
			if _, ok := num.SetString(parts[0], 10); !ok {
				fmt.Errorf("invalid numerator: %s", parts[0])
			}
			if _, ok := den.SetString(parts[1], 10); !ok {
				fmt.Errorf("invalid denominator: %s", parts[1])
			}
			rat := new(big.Rat)
			rat.SetFrac(num, den)

			fmt.Println(ht.Check(rat))
		case 0: // Якщо вибрано опцію 0, програма завершується.
			return
		default:
			continue
		}
	}
}
