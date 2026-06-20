package objects

type numberValue struct {
	raw     string
	isFloat bool
}

func (v *numberValue) Raw() string {
	return v.raw
}

func (v *numberValue) IsFloat() bool {
	return v.isFloat
}
