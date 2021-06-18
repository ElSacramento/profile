package configuration

type API struct {
	Listen string
}

func NewAPI() API {
	return API{
		Listen: "localhost:8080",
	}
}
