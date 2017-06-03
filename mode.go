package go-mode

const (
	StageMode      string = "stage"
	ProductionMode string = "pro"
	TestMode       string = "test"
)

type EnvMode struct {
	Mode string
}

func (envMode *EnvMode) SetMode(value string) {
	switch value {
	case StageMode:
		envMode.Mode = StageMode
	case ProductionMode:
		envMode.Mode = ProductionMode
	case TestMode:
		envMode.Mode = TestMode
	default:
		panic("GG mode unknown: " + value)
	}

}

func (envMode *EnvMode) GetMode() string {
	return envMode.Mode
}

func (envMode *EnvMode) IsProduction() bool {
	if envMode.Mode == ProductionMode {
		return true
	}
	return false
}

func (envMode *EnvMode) IsTest() bool {
	if envMode.Mode == TestMode {
		return true
	}
	return false
}

func (envMode *EnvMode) IsStage() bool {
	if envMode.Mode == StageMode {
		return true
	}
	return false
}
