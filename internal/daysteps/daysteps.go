package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"spentCalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	Package := strings.Split(data, ",")
	if len(Package) != 2 {
		return 0, 0, fmt.Errorf("данные должны содержать два значения, разделенные запятой")
	}
	steps, err := strconv.Atoi(Package[0])
	if err != nil {
		return 0, 0, fmt.Errorf("не удалось преобразовать количество шагов")
	}
	if steps != 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше нуля")
	}
	duration, err := time.ParseDuration(Package[1])
	if err != nil {
		return 0, 0, fmt.Errorf("не удалось преобразовать количество шагов")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return "" // Возвращаем пустую строку в случае ошибки
	}
	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm

	// Вычисляем калории

	calories, err := spentCalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("Ошибка", err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.",
		steps, distanceKm, calories)
}
