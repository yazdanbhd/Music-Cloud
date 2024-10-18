package minios3

type Config struct {
	Endpoint        string `koanf:"endpoint"`
	SecretAccessKey string `koanf:"secret_access_key"`
	AccessKeyID     string `koanf:"access_key_id"`
	UserSSL         bool   `koanf:"user_ssl"`
}

type MinioConf struct {
	config Config
}

func New(cfg Config) MinioConf {
	return MinioConf{config: cfg}
}
