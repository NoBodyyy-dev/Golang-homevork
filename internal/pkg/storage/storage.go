package storage

import (
	"go.uber.org/zap"
	"reflect"
)

type Storage struct {
	inner  map[string]interface{}
	Logger *zap.Logger
}

func NewStructure() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()
	logger.Info("NEW STORAGE")

	return Storage{
		inner:  make(map[string]interface{}),
		Logger: logger,
	}, nil
}

func (s Storage) Set(key string, value interface{}) {
	s.inner[key] = value
	s.Logger.Info("key set", zap.Any("key", key), zap.Any("value", value))
	s.Logger.Sync()
}

func (s Storage) Get(key string) *interface{} {
	res, ok := s.inner[key]
	if !ok {
		return nil
	}
	return &res
}

func (s Storage) GetKind(key string) *string {
	res, ok := s.inner[key]
	if !ok {
		return nil
	}

	switch reflect.TypeOf(res).Kind() {
	case reflect.String:
		str := "S"
		return &str
	case reflect.Int, reflect.Float64:
		str := "D"
		return &str
	default:
		return nil
	}
}
