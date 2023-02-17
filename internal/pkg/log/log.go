package log

// TODO: 创建日志库

type zaplogger struct {
	// z *zap.Logger
}

type Options struct {
}

var (
// mu sync.Mutex

// globlog 定义了默认的全局 Logger.
// globlog = NewLogger(NewOptions())
)

func NewLogger(opts *Options) *zaplogger {
	return &zaplogger{}
}

func NewOptions() *Options {
	return &Options{}
}
