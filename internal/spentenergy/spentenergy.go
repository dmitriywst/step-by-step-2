package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Функция принимает:
	// steps int — количество шагов.
	// weight, height float64 — вес(кг.) и рост(м.) пользователя.
	// duration time.Duration — продолжительность ходьбы.
	// И возвращает два значения:
	// float64 — количество калорий, потраченных при ходьбе.
	// error — ошибку, если входные параметры некорректны.
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Некорректные данные")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient

	return calories, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Функция принимает: steps int — количество шагов.
	// weight, height float64 — вес(кг.) и рост(м.) пользователя.
	// duration time.Duration — продолжительность бега.
	// И возвращает два значения:
	// float64 — количество калорий, потраченных при беге.
	// error — ошибку, если входные параметры некорректны.

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Некорректные данные")
	}
	meanSpeed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()

	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// Функция принимает количество шагов steps,
	// рост пользователя height и продолжительность активности duration
	// и возвращает среднюю скорость.

	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	averageSpeed := distance / duration.Hours()

	return averageSpeed

}

func Distance(steps int, height float64) float64 {
	// Функция принимает количество шагов и рост пользователя
	//  в метрах, а возвращает дистанцию в километрах.

	distance := (float64(steps) * (height * stepLengthCoefficient)) / float64(mInKm)
	return distance
}
