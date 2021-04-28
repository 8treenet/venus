package extjson

// SetDefaultOption .
func SetDefaultOption(entityOption ExtEntityOption) {
	defEntity = &ExtJSON{option: entityOption}
}

// ExtEntityOption .
type ExtEntityOption struct {
	NamedStyle       int
	SliceNotNull     bool
	StructPtrNotNull bool
}

func (exOption *ExtEntityOption) checkInvalid() {
	if exOption.NamedStyle == 0 {
		exOption.NamedStyle = NamedStyleUpperCamelCase
	}
}

// ExtJSON .
type ExtJSON struct {
	option ExtEntityOption
}

var defEntity *ExtJSON

func init() {
	SetDefaultOption(ExtEntityOption{})
}
