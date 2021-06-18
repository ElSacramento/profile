package configuration

type Cfg struct {
	API API
	DB  DB
}

func New() Cfg {
	return Cfg{
		API: NewAPI(),
		DB:  NewDB(),
	}
}
