package conversionHelpers

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

// Convert String Value Into Int
func StrToInt(data string) (int, error) {
	newData, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}

	return newData, nil
}

// Convert String Value Into Int64
func StrToInt64(data string) (int64, error) {
	newData, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}

	return newData, nil
}

// Convert Phone Number Into +62 Format
func ConvertPhoneNumber(data string) string {
	if data[0:1] == "0" {
		data = "+62" + data[1:]
	}
	if data[:2] == "62" {
		data = "+" + data
	}

	return data
}

// Convert Time Into TZ
func ConvertTimeToTZ(data time.Time, timezone string) time.Time {
	loc, _ := time.LoadLocation(timezone)

	newData := data.In(loc)

	return newData
}

// Hide First Some Digit Data
func HideFirstSomeDigitData(data string, digitToShow int) string {
	hideAccountNumber := ""
	for i, v := range data {
		if i < len(data)-digitToShow {
			hideAccountNumber += "*"
			continue
		}

		hideAccountNumber += fmt.Sprintf("%c", v)
	}

	return hideAccountNumber
}

// Convert To RP Currency
func ConvertToRpCurrency(data int64) string {
	humanizeValue := humanize.CommafWithDigits(float64(data), 0)
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)

	return stringValue
}

// Split string use "-"
func SplitStringDash(data string) []string {
	newArray := strings.Split(data, " - ")

	return newArray
}

func ConvertToPositiveInt(data int) int {
	if data > 0 {
		return data
	}

	return data * -1
}

func ConvertToPositiveFloat64(data float64) float64 {
	if data > 0 {
		return data
	}

	return data * -1
}

func CheckString(a string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9\s]+$`).MatchString(a)
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
