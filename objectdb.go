package personstore

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type ObjectDB interface {
	Store(ctx context.Context, object Object) error
	GetObjectByID(ctx context.Context, id string) (Object, error)
	GetObjectByName(ctx context.Context, name string) (Object, error)
	ListObjects(ctx context.Context, kind string) ([]Object, error)
	DeleteObject(ctx context.Context, id string) error
}

type RedisDB struct {
	client *redis.Client
}

func NewRedisDB(opt *redis.Options) (*RedisDB, error) {
	client := redis.NewClient(opt)
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &RedisDB{client: client}, nil
}

func (r *RedisDB) Store(ctx context.Context, object Object) error {
	object.SetID(uuid.New().String())
	return r.client.Set(object.GetId(), object.GetName(), 0).Err()
}

func (r *RedisDB) GetObjectByID(ctx context.Context, id string) (Object, error) {
	//todo: convert string to object
	name, err := r.client.Get(id).Result()
	if err != nil {
		return nil, err
	}

	return &Person{Name: name}, nil
}

func (r *RedisDB) GetObjectByName(ctx context.Context, name string) (Object, error) {
	return nil, nil //r.client.Get(name).Result()
}

func (r *RedisDB) GetObjects(ctx context.Context, kind string) ([]Object, error) {
	// return r.client.Get(kind).Result()
	return nil, nil
}

func (r *RedisDB) DeleteObject(ctx context.Context, id string) error {
	// return r.client.Del(id).Result()
	return nil
}
