package mocks

import entity "github.com/kharisma-wardhana/spe-academy-learn-golang/final-project/entity"

type MockLogger struct {
	LogErrorFunc func(processName string, funcName string, err error, logFields entity.CaptureFields, message string)
}

func (m *MockLogger) LogError(processName string, funcName string, err error, logFields entity.CaptureFields, message string) {
	if m.LogErrorFunc != nil {
		m.LogErrorFunc(processName, funcName, err, logFields, message)
	}
}
