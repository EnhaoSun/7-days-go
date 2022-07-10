package goecache

// A Getter loads data for a key.
type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFun implements Getter with a function.
type GetterFun func(key string) ([]byte, error)

func (f GetterFun) Get(key string) ([]byte, error) {
	return f(key)
}
