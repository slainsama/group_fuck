package models

type Config struct {
	Server struct {
		Host string `yaml:"host"`
	}
	Secure struct {
		Method    string   `yaml:"method"`
		Key       string   `yaml:"authKey"`
		Whitelist []string `yaml:"whiteList"`
		AesKey    string   `yaml:"aesKey"`
	}
}
