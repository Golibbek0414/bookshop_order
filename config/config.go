package config
import (
	"os"
	"github.com/spf13/cast"
)


type Config struct {
	Environment        string
	PostgresHost       string
	PostgresPort       int
	PostgresDatabase   string
	PostgresUser       string
	PostgresPassword   string
	LogLevel           string
	RPCPort            string
	CatalogServiceHost string
	CatalogServicePort int
	ReviewServiceHost  string
	ReviewServicePort  int
	CtxTimeout         int
}