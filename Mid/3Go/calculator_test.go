package main

import "testing"

func TestSum(t *testing.T) {
	expected := 55
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	if sum != expected {
		t.Errorf("Ошибка: ожидалось %d, получено %d", expected, sum)
	}
}
