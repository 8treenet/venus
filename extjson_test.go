package extjson_test

import (
	"testing"
	"time"

	"github.com/8treenet/extjson"
)

type Style struct {
	StyleText    string
	StyleNumber  int
	StyleBoolean bool
	StyleTag     string `json:"tagtagtag"`
}

func TestLowerCamelCase(t *testing.T) {
	extjson.SetDefaultOption(extjson.ExtJSONEntityOption{NamedStyle: extjson.NamedStyleLowerCamelCase})
	styleBytes, _ := extjson.Marshal(Style{StyleText: "extjson", StyleNumber: 100, StyleBoolean: true, StyleTag: "Tag"})
	t.Log(string(styleBytes))
	//out: {"styleText":"extjson","styleNumber":100,"styleBoolean":true,"tagtagtag":"Tag"}

	var read Style
	extjson.Unmarshal([]byte(`{"styleText":"extjson","styleNumber":100,"styleBoolean":true,"tagtagtag":"Tag"}`), &read)
	t.Log(read)
	//out: {extjson 100 true Tag}
}

func TestUnderScoreCase(t *testing.T) {
	extjson.SetDefaultOption(extjson.ExtJSONEntityOption{NamedStyle: extjson.NamedStyleUnderScoreCase})
	styleBytes, _ := extjson.Marshal(Style{StyleText: "extjson", StyleNumber: 100, StyleBoolean: true, StyleTag: "Tag"})
	t.Log(string(styleBytes))
	//out: {"style_text":"extjson","style_number":100,"style_boolean":true,"tagtagtag":"Tag"}

	var read Style
	extjson.Unmarshal([]byte(`{"style_text":"extjson","style_number":100,"style_boolean":true,"tagtagtag":"Tag"}`), &read)
	t.Log(read)
	//out: {extjson 100 true Tag}
}

func TestNull(t *testing.T) {
	extjson.SetDefaultOption(extjson.ExtJSONEntityOption{
		NamedStyle:       extjson.NamedStyleLowerCamelCase,
		SliceNotNull:     true,
		StructPtrNotNull: true,
	})
	var out struct {
		Slice  []*Style
		Slice2 []Style
		Struct *Style
	}
	outBytes, _ := extjson.Marshal(out)
	t.Log(string(outBytes))
	// {"slice":[],"slice2":[],"struct":{}}
}

type Timeformat struct {
	Time1 time.Time `json:",timeformat=ms"`
	Time2 time.Time `json:",timeformat=sec"`
	Time3 time.Time `json:",timeformat=2006-01-02 15:04:05"`
	Time4 time.Time
}

func TestTimeformat(t *testing.T) {
	out := Timeformat{
		Time1: time.Now(),
		Time2: time.Now(),
		Time3: time.Now(),
		Time4: time.Now(),
	}
	outBytes, _ := extjson.Marshal(out)
	t.Log(string(outBytes))
	//out : {"Time1":"1578799200468","Time2":"1578799200","Time3":"2020-01-12 11:20:00","Time4":"2020-01-12T11:20:00.468543+08:00"}

	var in Timeformat
	extjson.Unmarshal([]byte(`{"Time1":"1578799200468","Time2":"1578799200","Time3":"2020-01-12 11:20:00","Time4":"2020-01-12T11:20:00.468543+08:00"}`), &in)
	t.Log(in)
	//out : {2020-01-12 11:20:00.468 +0800 CST 2020-01-12 11:20:00 +0800 CST 2020-01-12 11:20:00 +0000 UTC 2020-01-12 11:20:00.468543 +0800 CST}
}
