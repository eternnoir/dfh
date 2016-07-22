package main

import (
	"errors"
	"time"
)

func ParseBaseTime(td string, template string) (time.Time, error) {
	if td == "" {
		return time.Now(), nil
	}
	return time.Parse(template, td)
}

func ParseStartTime(td string, starttime string, baseTime time.Time) (time.Time, error) {
	if td != "" && starttime != "" {
		return time.Time{}, errors.New("Time duration and Start Time only one can be set.")
	}

	if td != "" {
		dur, err := time.ParseDuration(td)
		if err != nil {
			return time.Time{}, err
		}
		return CalStartDate(dur, baseTime), nil
	}

	if td == "" && starttime == "" {
		return time.Time{}, nil
	}

	// TODO
	return time.Time{}, errors.New("Not Impl.")
}

func CalStartDate(duration time.Duration, baseTime time.Time) time.Time {
	return baseTime.Add(-1 * duration)

}
