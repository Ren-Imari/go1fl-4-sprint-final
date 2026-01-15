package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"log"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	list := strings.Split(data, ",")
	if len(list) != 2 {
		log.Println("неверное количество параметров")
		return 0, 0, fmt.Errorf("неверное количество параметров")
	}
	steps, err := strconv.Atoi(list[0])
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	if steps <= 0 {
		log.Println("неверное количество шагов")
		return 0, 0, fmt.Errorf("неверное количество шагов")
	}
	duration, err := time.ParseDuration(list[1])
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	if duration <= 0 {
		log.Println("неверная продолжительность")
		return 0, 0, fmt.Errorf("неверная продолжительность")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distance := float64(steps) * stepLength
	distanceKm := distance / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
}
