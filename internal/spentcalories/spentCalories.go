package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	s := strings.Split(data, ",")
	if len(s) != 3 {
		return 0, "", 0, fmt.Errorf("Input incorrect")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return 0, "", 0, err
	}
	duration, err := time.ParseDuration(s[2])
	if err != nil {
		return 0, "", 0, err
	}
	activity := s[1]
	if s[1] == "" {
		return steps, "", duration, nil
	}
	return steps, activity, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	dist := (float64(steps) * lenStep) / mInKm
	return dist
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration < 0 {
		return 0
	}
	dist := distance(steps)
	mean := dist / duration.Hours()
	return mean
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, _, duration, _ := parseTraining(data)

	_, activity, _, _ := parseTraining(data)

	if activity != "Бег" && activity != "Ходьба" {
		fmt.Println("неизвестный тип тренировки")
	}
	dist := distance(steps)

	dur := duration.Hours()

	speed := meanSpeed(steps, duration)

	var calories float64

	switch activity {

	case "Бег":
		calories = RunningSpentCalories(steps, weight, duration)

	case "Ходьба":
		calories = WalkingSpentCalories(steps, weight, height, duration)

		//default:
		//fmt.Println("неизвестный тип тренировки")

	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f. ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", activity, dur, dist, speed, calories)
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	mean := meanSpeed(steps, duration)
	calories := (runningCaloriesMeanSpeedMultiplier * mean) - runningCaloriesMeanSpeedShift*weight
	return calories

}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	mean := meanSpeed(steps, duration)
	calories := ((walkingCaloriesWeightMultiplier * weight) + (mean*mean/height)*walkingSpeedHeightMultiplier) * float64(duration) * minInH
	return calories
}
