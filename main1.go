package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Структура для хранения информации о студенте
type Student struct {
	FIO     string    // ФИО студента
	Grades  []float64 // Оценки
	Average float64   // Средний балл
}

// Метод для подсчета среднего балла студента
func (s *Student) CalculateAverage() {
	if len(s.Grades) == 0 {
		s.Average = 0
		return
	}
	sum := 0.0
	for _, grade := range s.Grades {
		sum += grade
	}
	s.Average = sum / float64(len(s.Grades))
}

func main() {
	var students []Student
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1 - Добавить студента")
		fmt.Println("2 - Показать всех студентов")
		fmt.Println("3 - Фильтрация по среднему баллу")
		fmt.Println("4 - Выйти")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Некорректный ввод, попробуйте снова")
			continue
		}

		switch choice {
		case 1:
			// Ввод данных о студенте
			var student Student

			fmt.Print("Введите ФИО: ")
			fio, _ := reader.ReadString('\n')
			student.FIO = strings.TrimSpace(fio)

			fmt.Print("Введите оценки через запятую: ")
			gradesStr, _ := reader.ReadString('\n')
			gradesStr = strings.TrimSpace(gradesStr)
			gradesSlice := strings.Split(gradesStr, ",")
			for _, g := range gradesSlice {
				g = strings.TrimSpace(g)
				grade, err := strconv.ParseFloat(g, 64)
				if err != nil {
					fmt.Println("Некорректная оценка:", g)
					continue
				}
				student.Grades = append(student.Grades, grade)
			}
			student.CalculateAverage()
			students = append(students, student)
			fmt.Println("Студент добавлен!")

		case 2:
			// Вывод всех студентов
			if len(students) == 0 {
				fmt.Println("Нет данных о студентах.")
				continue
			}
			for i, s := range students {
				fmt.Printf("%d. %s | Оценки: %v | Средний: %.2f\n", i+1, s.FIO, s.Grades, s.Average)
			}

		case 3:
			// Фильтрация по среднему баллу
			fmt.Print("Введите порог среднего балла: ")
			thresholdStr, _ := reader.ReadString('\n')
			thresholdStr = strings.TrimSpace(thresholdStr)
			threshold, err := strconv.ParseFloat(thresholdStr, 64)
			if err != nil {
				fmt.Println("Некорректный ввод")
				continue
			}

			fmt.Printf("Студенты с средним баллом ниже %.2f:\n", threshold)
			for _, s := range students {
				if s.Average < threshold {
					fmt.Printf("%s | Средний: %.2f\n", s.FIO, s.Average)
				}
			}

		case 4:
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Некорректный выбор, попробуйте снова.")
		}
	}
}
