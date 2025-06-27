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
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	dataInput := strings.Split(datastring, ",")

	if len(dataInput) != 3 {
		err := errors.New("incorrect data entry")
		//log.Println(err)
		return err
	}
		
	stepsNumbers, err := strconv.Atoi(dataInput[0])
	if err != nil {
		//	log.Println(err)
		return err
	}
	if stepsNumbers <= 0 {
		err := errors.New("incorrect data entry on the number of steps")
		//log.Println(err)
		return err
	}
	t.Steps = stepsNumbers
	t.TrainingType = dataInput[1]
	duration, err := time.ParseDuration(dataInput[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		err := errors.New("incorrect time data entry")
		return err
	}
	t.Duration = duration
	return nil
}
func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanspeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	activityMode := t.TrainingType

	switch activityMode {
	case "Бег":
		runModeEnergy, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityMode, float64(t.Duration.Hours()), distance, meanspeed, runModeEnergy), nil
	case "Ходьба":
		walkModeEnergy, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
     	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityMode, float64(t.Duration.Hours()), distance, meanspeed, walkModeEnergy), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}
