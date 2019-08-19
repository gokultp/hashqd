package config

type Config struct {
	Service *Service `toml:"service"`
	AdminUI *AdminUI `toml:"admin_ui"`
}

type Service struct {
	Port     *string `toml:"port"`
	Logs     *string `toml:"logs"`
	DataPath *string `toml:"data_path"`
}

type AdminUI struct {
	Enable bool
	Port   *int
}
