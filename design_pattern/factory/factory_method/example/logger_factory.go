package example

//-------------------日志记录器工厂接口-------------------

type LoggerFactory interface {
	CreateLogger() Logger
}

//-------------------数据库日志工厂-----------------------
type DbLoggerFactory struct{}

func (this *DbLoggerFactory) CreateLogger() Logger {
	return &DbLogger{}
}

//-------------------文件日志工厂-------------------------
type FileLoggerFactory struct{}

func (this *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}
