package helper

type Helper struct {
	LoggerHelper LoggerHelper
	CacheHelper  CacheHelper
}

func NewHelper() Helper {
	logger := NewLoggerHelper()
	return Helper{
		LoggerHelper: logger,
		CacheHelper:  NewCacheHelper(logger),
	}
}
