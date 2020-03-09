package config

type Config struct {
	MySQLURL string `envconfig:"MYSQL_URL" default:"root:1@tcp(localhost:3306)/cm?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"`
	APIKey   string `envconfig:"API_KEY" default:"controlnomey-hongminh-229297"`
	SecretKet string `envconfig:"SECRET_KEY" default:"controlnomey-hongminh-229297"`
}
