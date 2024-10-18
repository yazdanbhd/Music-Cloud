package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/labstack/gommon/log"
	"github.com/yazdanbhd/Music-Cloud/repository/mysqldb"
	"github.com/yazdanbhd/Music-Cloud/repository/s3/minios3"
	"strings"
)

const (
	Prefix = "MusicCloud"
)

type Config struct {
	MinioS3  minios3.Config `koanf:"minio"`
	DataBase mysqldb.Config `koanf:"db"`
}

func New(path string) Config {
	var instance Config

	k := koanf.New(".")

	// load configuration from file
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	// load environment variables
	if err := k.Load(env.Provider(Prefix, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, Prefix)), "_", ".")
	}), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	if err := k.Unmarshal("config", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	log.Printf("following configuration is loaded:\n%+v", instance)

	return instance
}
