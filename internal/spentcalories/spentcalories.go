package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	Training := strings.Split(data, ",")
	if len(Training) != 3 {
		return 0, "", 0, fmt.Errorf("данные должны содеражить три значения, разделенные запятой")
	}
	steps, err := strconv.Atoi(Training[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("не удалось преобразовать количество шагов")
	}
	duration, err := time.ParseDuration(Training[1])
	if err != nil {
		return 0, "", 0, fmt.Errorf("некорректная продолжительность")
	}
	return steps, Training[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)
	speed := distanceKm / (duration.Seconds() / 3600)
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		// Если произошла ошибка, возвращаем её
		return "", err
	}

	// Здесь можно добавить дополнительную логику для обработки полученных данных
	// Например, вычислить дистанцию или скорость

	var calories float64
	var errCalories error

	switch activity {
	case "бег":
		calories, errCalories = RunningSpentCalories(steps, weight, height, duration)

	case "ходьба":
		calories, errCalories = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный вид активности: %s", activity)

	}
	if errCalories != nil {
		return "", errCalories
	}
	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	return fmt.Sprintf("Тип тренировки: %s\n"+"Длительность: %.2f ч.\n"+"Дистанция: %.2f км.\n"+"Средняя скорость: %.2f км/ч\n"+"Калории: %.2f", activity, duration.Hours(), dist, speed, calories), err
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("все параметры должныбыть больше нуля")
	}

	RunningMeanSpeed := meanSpeed(steps, weight, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * RunningMeanSpeed * durationInMinutes) / minInH
	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("все параметры должныбыть больше нуля")
	}

	WalkingMeanSpeed := meanSpeed(steps, weight, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * WalkingMeanSpeed * durationInMinutes) / minInH
	adjustedCalories := spentCalories * walkingCaloriesCoefficient
	return adjustedCalories, nil
}
