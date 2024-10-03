package minios3

type Config struct {
	Endpoint        string
	SecretAccessKey string
	AccessKeyID     string
	UserSSL         bool
}

type MinioConf struct {
	config Config
}

func New(cfg Config) MinioConf {
	return MinioConf{config: cfg}
}
