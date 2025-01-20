package factory

type Factory[T any] interface {
	// GetOrCreate 获取或创建实例
	GetOrCreate() (string, T, error)
}
