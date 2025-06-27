package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)


type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	dataInput := strings.Split(datastring, ",")
	
	if len(dataInput) != 2 {
		err := errors.New("incorrect data entry")
		return err
	}
	stepsNumbers, err := strconv.Atoi(dataInput[0])
	if err != nil {
		return err
	}
	if stepsNumbers <= 0 {
		err := errors.New("incorrect data entry on the number of steps")
		return err
	}
	ds.Steps = stepsNumbers
	t, err := time.ParseDuration(dataInput[1])
	if err != nil {
		return err
	}
	if t <= 0 {
		err := errors.New("incorrect time data entry")
		return err
	}
	
	ds.Duration = t
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance:=spentenergy.Distance(ds.Steps, ds.Height)//float64(ds.Steps) * float64(ds.Personal.Height) * stepLengthCoefficient / 1000
	walkModeEnergy, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)//distance * ds.Duration.Hours() * float64(ds.Personal.Weight) * walkingCaloriesCoefficient
	if err != nil {
			return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, walkModeEnergy), nil
}
