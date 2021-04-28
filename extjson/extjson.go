package extjson

// SetDefaultOption .
func SetDefaultOption(entityOption ExtOption) {
	defEntity = &ExtJSON{option: entityOption}
}

// ExtOption .
type ExtOption struct {
	NamedStyle       int
	SliceNotNull     bool
	StructPtrNotNull bool
}

func (exOption *ExtOption) checkInvalid() {
	if exOption.NamedStyle == 0 {
		exOption.NamedStyle = NamedStyleUpperCamelCase
	}
}

// ExtJSON .
type ExtJSON struct {
	option ExtOption
}

var defEntity *ExtJSON

func init() {
	SetDefaultOption(ExtOption{})
}
