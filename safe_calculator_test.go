package main

import "testing"

func TestDivideByZero(t *testing.T) {
	_, err := calculate(10, 0, "/")
	if err == nil {
		t.Errorf("ожидалась ошибка при делении на ноль")
	}
}

func TestUnknownOp(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ожидалась паника для неизвестной операции")
		}
	}()
	calculate(1, 2, "^")
}
