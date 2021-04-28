package extjson

func SetDefaultOption(entityOption ExtJSONEntityOption) {
	defEntity = &ExtJSONEntity{option: entityOption}
}

type ExtJSONEntityOption struct {
	NamedStyle       int
	SliceNotNull     bool
	StructPtrNotNull bool
}

func (exOption *ExtJSONEntityOption) checkInvalid() {
	if exOption.NamedStyle == 0 {
		exOption.NamedStyle = NamedStyleUpperCamelCase
	}
}

type ExtJSONEntity struct {
	option ExtJSONEntityOption
}

var defEntity *ExtJSONEntity

func init() {
	SetDefaultOption(ExtJSONEntityOption{})
}
