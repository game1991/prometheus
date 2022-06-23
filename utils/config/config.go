package config



// ServerConf server config
type ServerConf struct {
	Address string
}

// Config config
type Config struct {
	Server map[string]*ServerConf
	DB     *db.ConnConf
	Redis  *rediser.ConnConf
	Log    log.Options
}