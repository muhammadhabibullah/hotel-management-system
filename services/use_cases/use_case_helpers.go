package usecases

import "time"

func parseDate(str string) (time.Time, error) {
	date, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return time.Time{}, err
	}
	return roundDate(date), nil
}

func roundDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}
