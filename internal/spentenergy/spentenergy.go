package spentenergy

import (
	"errors"
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
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("numbers of the steps<=0")
	}
	if weight <= 0 {
		return 0, errors.New("weight<=0")
	}
	if height <= 0 {
		return 0, errors.New("height<=0")
	}
	if duration <= 0 {
		return 0, errors.New("time spent<=0")
	}
	return MeanSpeed(steps, height, duration) * duration.Minutes() * weight / minInH * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("numbers of the steps<=0")
	}
	if weight <= 0 {
		return 0, errors.New("weight<=0")
	}
	if height <= 0 {
		return 0, errors.New("height<=0")
	}
	if duration <= 0 {
		return 0, errors.New("time spent<=0")
	}
	return MeanSpeed(steps, height, duration) * duration.Minutes() * weight / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return float64(steps) * height * stepLengthCoefficient / 1000
}
