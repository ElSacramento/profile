package configuration

type DB struct {
	URL string
}

func NewDB() DB {
	return DB{
		URL: "postgres://postgres:@localhost/test_db?sslmode=disable&application_name=test",
	}
}
