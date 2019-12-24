package Storage

type Storage interface {
	Add(key, value string) error
	Get(key string) (string, error)
	Set(key, value string) error
	Del(key string) error
}
