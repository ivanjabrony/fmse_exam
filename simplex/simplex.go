package main

import (
	"fmt"
	"math"
)

// Симплекс-таблица
type Table struct {
	Table      [][]float64 // Основная таблица
	Rows, Cols int         // Размеры таблицы (включая целевую строку и правую часть)
}

// NewTable создает таблицу для задачи:
// Maximize cᵀx, при Ax <= b, x >= 0
func NewTable(A [][]float64, b, c []float64) *Table {
	rows := len(A) + 1
	cols := len(A[0]) + len(A) + 1

	// Инициализация таблицы
	table := make([][]float64, rows)
	for i := range table {
		table[i] = make([]float64, cols)
	}

	// Заполняем ограничения
	for i := range len(A) {
		copy(table[i], A[i])
		table[i][len(A[i])+i] = 1 // Добавляем slack-переменные
		table[i][cols-1] = b[i]   // Правая часть (RHS)
	}

	// Целевая функция (последняя строка)
	for j := range len(c) {
		table[rows-1][j] = -c[j] // Инвертируем для максимизации
	}

	return &Table{
		Table: table,
		Rows:  rows,
		Cols:  cols,
	}
}

func (t *Table) FindPivot() (int, int) {
	// Ищем вводящую переменную (колонку с минимальным значением в целевой строке)
	col := -1
	minVal := 0.0
	for j := 0; j < t.Cols-1; j++ {
		if t.Table[t.Rows-1][j] < minVal {
			minVal = t.Table[t.Rows-1][j]
			col = j
		}
	}
	if col == -1 {
		return -1, -1 // Оптимальное решение найдено
	}

	// Ищем выводимую переменную (минимальное отношение RHS / A[i][col])
	row := -1
	minRatio := math.Inf(1)
	for i := 0; i < t.Rows-1; i++ {
		if t.Table[i][col] <= 0 {
			continue
		}
		ratio := t.Table[i][t.Cols-1] / t.Table[i][col]
		if ratio < minRatio {
			minRatio = ratio
			row = i
		}
	}

	return row, col
}

func (t *Table) Pivot(row, col int) {
	// Нормализуем опорную строку
	pivotVal := t.Table[row][col]
	for j := range t.Cols {
		t.Table[row][j] /= pivotVal
	}

	// Обновляем остальные строки
	for i := range t.Rows {
		if i == row {
			continue
		}
		mult := t.Table[i][col]
		for j := 0; j < t.Cols; j++ {
			t.Table[i][j] -= mult * t.Table[row][j]
		}
	}
}

func (t *Table) Solve() (float64, []float64) {
	for {
		row, col := t.FindPivot()
		if col == -1 {
			break // Оптимум достигнут
		}
		if row == -1 {
			return math.Inf(1), nil // Задача неограничена
		}
		t.Pivot(row, col)
	}

	// Извлекаем решение
	solution := make([]float64, t.Cols-1)
	for i := 0; i < t.Rows-1; i++ {
		for j := 0; j < t.Cols-1; j++ {
			if t.Table[i][j] == 1 {
				// Проверяем, что в столбце только одна 1
				isBasic := true
				for k := 0; k < t.Rows; k++ {
					if k != i && t.Table[k][j] != 0 {
						isBasic = false
						break
					}
				}
				if isBasic {
					solution[j] = t.Table[i][t.Cols-1]
				}
			}
		}
	}

	// Значение целевой функции
	optimum := t.Table[t.Rows-1][t.Cols-1]
	return optimum, solution
}

func main() {
	// Пример задачи:
	// Maximize: 3x1 + 2x2
	// При условиях:
	//   x1 + x2 ≤ 4
	//   x1 - x2 ≤ 2
	//   x1, x2 ≥ 0
	A := [][]float64{
		{1, 1},  // x1 + x2 ≤ 4
		{1, -1}, // x1 - x2 ≤ 2
	}
	b := []float64{4, 2}
	c := []float64{3, 2} // Целевая функция: 3x1 + 2x2

	// Инициализация и решение
	Table := NewTable(A, b, c)
	optimum, solution := Table.Solve()

	fmt.Printf("Optimal: %.2f\n", optimum)
	fmt.Printf("Solution: x1 = %.2f, x2 = %.2f\n", solution[0], solution[1])
}
