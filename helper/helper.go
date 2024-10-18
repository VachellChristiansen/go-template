package helper

type Helper struct {
	LoggerHelper LoggerHelper
}

func NewHelper() Helper {
	return Helper{
		LoggerHelper: NewLoggerHelper(),
	}
}
