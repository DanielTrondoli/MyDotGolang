package timeutils

import (
	"fmt"
	"log"
	"time"
)

const LAYOUT_DMY = "02/01/2006"
const LAYOUT_DMYHMS = "02/01/2006 15:04:05"

func GetNowTruncated() (time.Time, error) {

	year, month, day := time.Now().Date()

	now, err := time.Parse(LAYOUT_DMY, fmt.Sprintf("%d/%d/%d", day, month, year))
	if err != nil {
		return time.Time{}, err
	}

	return now, nil
}

func GetNowEndDay() (time.Time, error) {

	year, month, day := time.Now().Date()

	now, err := time.Parse(LAYOUT_DMYHMS, fmt.Sprintf("%d/%d/%d 23:59:59", day, month, year))
	if err != nil {
		return time.Time{}, err
	}

	return now, nil
}

func DifDateFromNow(strDate string) (time.Duration, error) {

	date, err := time.Parse(LAYOUT_DMY, strDate)
	if err != nil {
		return -1, err
	}

	now, err := GetNowTruncated()
	if err != nil {
		log.Default().Print(err.Error())
		return -1, err
	}

	fmt.Println(time.Duration(date.Unix() - now.Unix()))

	return date.Sub(now), nil
}

func GetDayFormatedDMY(strDate string, add int) (string, error) {
	if strDate == "" {
		strDate = time.Now().Format(LAYOUT_DMY)
	}

	date, err := time.Parse(LAYOUT_DMY, strDate)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	retDate := date.Add(time.Duration(add * 24 * int(time.Hour))).Format(LAYOUT_DMY)
	return retDate, nil
}
