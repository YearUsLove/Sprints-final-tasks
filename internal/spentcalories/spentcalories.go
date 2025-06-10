package spentcalories

import (
	"errors"
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
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("длина слайса не равна 3")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования шагов: %w", err)
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("шаги меньше или равны 0")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка парсинга продолжительности: %w", err)
	}

	if duration <= 0 {
		return 0, "", 0, errors.New("продолжительность должна быть больше 0")
	}
	return steps, parts[1], duration, nil
}
func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient

	distanceM := float64(steps) * stepLength

	distanceKm := distanceM / mInKm

	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)

	hours := duration.Hours()

	var avgSpeed float64

	if hours != 0 {
		avgSpeed = dist / hours
	} else {
		avgSpeed = 0
	}
	return avgSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		return "", fmt.Errorf("ошибка парсинга: %w", err)
	}
	distanceKm := distance(steps, height)

	speed := meanSpeed(steps, height, duration)

	durationHours := duration.Hours()

	var calories float64

	switch activityType {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", fmt.Errorf("ошибка подсчета калорий: %w", err)
		}

	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", fmt.Errorf("ошибка подсчета калорий: %w", err)
		}

	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activityType,
		durationHours,
		distanceKm,
		speed,
		calories,
	)
	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("шагов не должно быть меньше или равно 0")
	}

	if weight <= 0 {
		return 0, errors.New("вес не должен быть меньше или равно 0")
	}

	if height <= 0 {
		return 0, errors.New("высота не должна быть меньше или равно 0")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность не должна быть меньше или равно 0")
	}

	speed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("шагов не должно быть меньше или равно 0")
	}

	if weight <= 0 {
		return 0, errors.New("вес не должен быть меньше или равно 0")
	}

	if height <= 0 {
		return 0, errors.New("высота не должна быть меньше или равно 0")
	}

	if duration <= 0 {
		return 0, errors.New("продолжительность не должна быть меньше или равно 0")
	}

	speed := meanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * speed * durationInMinutes) / minInH

	calories *= walkingCaloriesCoefficient

	return calories, nil
}
