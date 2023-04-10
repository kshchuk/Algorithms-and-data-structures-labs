package main

import (
	"fmt"
	"math/big"
)

type HashTableInterface interface {
	Check(rat *big.Rat) *big.Rat
	Print()
}

// Структура хеш таблиці
type HashTable struct {
	function  *HashFunction // Хеш-функція цієї таблиці
	size      int           // Розмір таблиці
	rats      []*big.Rat    // Значення, які потрібно захешувти
	table     []*big.Rat    // Сама таблиця
	subTables []*HashTable  // Підтаблиці (Поки немає колізій в основній таблиці мають значення nil)
}

// Функція створює нову хеш-таблицю на основі даних, збережених у файлі з вказаним ім'ям.
func NewHashTable(filename string) (*HashTable, error) {

	// Зчитуємо раціональні числа з файлу
	rats, err := ReadRats(filename)

	if err != nil {
		return nil, err
	}

	hashTable := &HashTable{}

	hashTable.rats = RemoveDuplicates(rats) // Видаляємо повторення (випадки 2/2 та 5/5 вважаються одним числом)
	hashTable.size = len(hashTable.rats)
	hashTable.function = NewHashFunction(hashTable.size)
	hashTable.table = make([]*big.Rat, hashTable.size)
	hashTable.subTables = make([]*HashTable, hashTable.size)

	for _, rat := range hashTable.rats {
		k := hashTable.function.hash((rat))

		// Якщо комірка основної таблиці порожня
		if hashTable.table[k] == nil {
			// Якщо комірка основної таблиці порожня та відсутня її підтаблиця
			if hashTable.subTables[k] == nil {
				hashTable.table[k] = rat
			} else { // Підтаблиця вже існує
				hashTable.subTables[k] = hashTable.NewSubTable(rat, hashTable.subTables[k])
			}
		} else { // Колізія у основній таблиці, створюємо підтаблицю з двома елементами та чотирма комірками
			hashTable.subTables[k] = hashTable.NewSubTable(hashTable.table[k], nil)
			hashTable.subTables[k] = hashTable.NewSubTable(rat, hashTable.subTables[k])
			hashTable.table[k] = nil
		}
	}
	return hashTable, nil
}

// Функція, яка генерує підтаблицю на основі попередньої, додаючи до неї новий елемент
func (*HashTable) NewSubTable(toAdd *big.Rat, oldSubTable *HashTable) *HashTable {

	subTable := &HashTable{}

	// Якщо попердньої підтаблиці нема
	if oldSubTable == nil {
		subTable.rats = append(subTable.rats, toAdd)
		subTable.size = 1
		subTable.table = append(subTable.table, toAdd)
	} else { // існує попередня підтаблиця
		subTable.rats = append(oldSubTable.rats, toAdd)
		subTable.size = len(subTable.rats) * len(subTable.rats)
		subTable.function = NewHashFunction(subTable.size)
		subTable.table = make([]*big.Rat, subTable.size)

		for _, rat := range subTable.rats {
			k := subTable.function.hash(rat)
			if subTable.table[k] == nil {
				subTable.table[k] = rat
			} else { // Присутня колізія у підтаблиці => перехешовуємо підтаблицю
				return oldSubTable.NewSubTable(toAdd, oldSubTable)
			}
		}
	}
	return subTable
}

// Функція, яка видалаяє однакові раціональні числа
func RemoveDuplicates(rats []*big.Rat) []*big.Rat {
	result := []*big.Rat{}
	encountered := false

	for _, rat := range rats {
		encountered = false
		for _, res := range result {
			if res == rat {
				encountered = true
				break
			}
		}

		if !encountered {
			result = append(result, rat)
		}
	}
	return result
}

// Функція, яка перевіряє чи є раціональне число у таблиці.
func (table *HashTable) Check(rat *big.Rat) *big.Rat {

	k := table.function.hash(rat) // Дізнаємось індекс в основній таблиці

	if table.table[k] == nil { // Якщо комірка порожня
		if table.subTables[k] == nil { // Якщо підтаблиця порожня
			return nil
		} else { // Якщо існує підтаблиця
			k1 := table.subTables[k].function.hash(rat) // Дізнаємось індекс числа у підтаблиці
			return table.subTables[k].table[k1]
		}
	} else {
		return table.table[k]
	}
}

// Вивід усієї хеш-таблиці
func (ht *HashTable) Print() {

	for i, rat := range ht.table {
		fmt.Print("Element ", i, ": ")

		// Якщо комірка не порожня
		if rat != nil {
			fmt.Print(rat, "\n")
		} else if ht.subTables[i] != nil {
			// Якщо є підтаблиця
			fmt.Print(" Subtable: ")
			for _, rat_1 := range ht.subTables[i].table {
				fmt.Print(rat_1, ", ")
			}
			fmt.Print("\n")
		} else {
			// Комірка і підтаблиця порожні
			fmt.Print(rat, "\n")
		}
	}
}
