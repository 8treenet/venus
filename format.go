package extjson

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func namedStyleCaseFormat(field string, exEntity *ExtJSONEntity) string {
	switch exEntity.option.NamedStyle {
	case NamedStyleUnderScoreCase:
		return underScoreCase(field)
	case NamedStyleLowerCamelCase:
		return lowerCamelCase(field)
	default:
		return field
	}
}

func lowerCamelCase(field string) string {
	var lowerStr string
	vv := []rune(field)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			lowerStr += strings.ToLower(string(vv[i]))
		} else {
			lowerStr += string(vv[i])
		}
	}
	return lowerStr
}
func underScoreCase(field string) string {
	var result string
	if field == "" {
		return field
	}

	list := []rune(field)
	result += string(list[0])
	for index := 1; index < len(list); index++ {
		if list[index] > 64 && list[index] < 91 && list[index-1] > 96 && list[index-1] < 123 {
			result += "_"
		}
		result += string(list[index])
	}

	return strings.ToLower(result)
}

var timeType = reflect.TypeOf(time.Time{})

func timeTo(value reflect.Value, format string) string {
	ti := value.Interface()
	timeValue, ok := ti.(time.Time)
	if !ok {
		return ""
	}
	if format == "" {
		b := make([]byte, 0, len(time.RFC3339Nano)+2)
		b = timeValue.AppendFormat(b, time.RFC3339Nano)
		return string(b)
	}

	if format == "sec" {
		return fmt.Sprint(timeValue.Unix())
	}

	if format == "ms" {
		ms := time.Now().UnixNano() / 1e6
		return fmt.Sprint(ms)
	}

	return timeValue.Format(format)
}

func toTime(value string, format string) (t time.Time, e error) {
	if format == "" {
		e = t.UnmarshalJSON([]byte(value))
		return
	}
	if len(value) < 2 {
		return
	}
	value = value[1 : len(value)-1]
	if format == "sec" {
		sec, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			e = err
			return
		}
		t = time.Unix(sec, 0)
		return
	}

	if format == "ms" {
		ms, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			e = err
			return
		}
		t = time.Unix(0, ms*1e6)
		return
	}

	t, e = time.Parse(format, value)
	return
}
