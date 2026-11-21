package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps                 int           // количество шагов, проделанных за тренировку.
	TrainingType          string        // тип тренировки(бег или ходьба).
	Duration              time.Duration // длительность тренировки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print().

}

func (t *Training) Parse(datastring string) (err error) {
	result := strings.SplitN(datastring, ",", 3)
	if len(result) != 3 {
		return fmt.Errorf("Неверный формат полученных данных")
	}
	t.Steps, err = strconv.Atoi(result[0])
	if err != nil {
		return fmt.Errorf("Неверный формат полученных данных: %w", err)
	}

	t.TrainingType = result[1]

	t.Duration, err = time.ParseDuration(result[2])
	if err != nil {
		return fmt.Errorf("Неверный формат полученных данных: %w", err)
	}

	return fmt.Errorf("Неверный формат полученных данных: %w", err)
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)               // дистанция
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration) // скорость

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}
