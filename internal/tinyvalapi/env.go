package tinyvalapi

type Config struct {
	DataDir string `env:"TINYVAL_DATA_DIR,required"`
	Port    uint16 `env:"TINYVAL_PORT" envDefault:"8080"`
}
