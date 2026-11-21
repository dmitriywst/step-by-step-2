package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps                 int           // количество шагов
	Duration              time.Duration // длительность прогулки
	personaldata.Personal               // встроенная структура Personal из пакета personaldata, у которой есть метод Print().
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	result := strings.SplitN(datastring, ",", 2)
	if len(result) != 2 {
		return fmt.Errorf("Неверный формат полученных данных")
	}
	ds.Steps, err = strconv.Atoi(result[0])
	if err != nil {
		return fmt.Errorf("Неверный формат полученных данных")
	}

	ds.Duration, err = time.ParseDuration(result[1])
	if err != nil {
		return fmt.Errorf("Неверный формат полученных данных")
	}

	return fmt.Errorf("Неверный формат полученных данных")

}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height) // дистанция

	var err error
	var calories float64

	calories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
