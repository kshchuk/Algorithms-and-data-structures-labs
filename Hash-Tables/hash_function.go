package main

import (
	"math/big"
	"math/rand"
	"time"
)

// Структура хеш-функції
type HashFunction struct {
	A int // Випадкове велике просте число
	B int // Випадкове велике просте число
	M int // Розмір таблиці, де застосовуєтсья ця хеш фунція
	P int // Перше просте число, яке більше за M
}

// Функція створює новий об'єкт типу HashFunction, який представляє
// Хеш-функцію для таблиці з заданою кількістю слотів M.
func NewHashFunction(M int) *HashFunction {
	rand.Seed(time.Now().UnixNano())

	P := nextPrime(M)
	A := nextPrime(rand.Intn(1000000))
	B := nextPrime(rand.Intn(1000000))
	return &HashFunction{A, B, M, P}
}

// Повертає хеш для заданого раціонального числа
func (hf *HashFunction) hash(rat *big.Rat) int {

	num := rat.Num().Uint64()
	denom := rat.Denom().Uint64()

	// Робимо так, аби хеші протилежних чисел відрізнялись
	var hash uint64
	if rat.Num().Sign() < 0 {
		hash = ^num
	} else {
		hash = num
	}

	// хешуємо
	hash = (hash * uint64(hf.B)) ^ (denom << 32)
	hash = (hash * uint64(hf.A)) ^ (hash >> 32)

	return int(hash % uint64(hf.P) % uint64(hf.M))
}

// Функція пошуку наступного простого числа
func nextPrime(n int) int {
	if n < 2 {
		return 2
	}

	for i := n + 1; ; i++ {
		if i%2 == 0 {
			continue
		}

		if isPrime(i) {
			return i
		}
	}
}

// Функція перевірки числа на простоту
func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
