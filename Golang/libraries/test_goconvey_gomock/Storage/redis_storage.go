package Storage

type RedisStorage struct {

}

func (repo *RedisStorage) Add(key, value string) error {
	return nil
}

func (repo *RedisStorage) Get(key string) (string, error) {
	return "", nil
}

func (repo *RedisStorage) Set(key, value string) error {
	return nil
}

func (repo *RedisStorage) Del(key string) error {
	return nil
}

