package config

type Config struct {
	MySQLURL string `envconfig:"MYSQL_URL" required:"true"`
	APIKey   string `envconfig:"API_KEY" default:"controlnomey-hongminh-229297"`
}
