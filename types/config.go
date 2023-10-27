package types

type Configuration struct {
	ConnectionString string
}

type ResponseApi struct {
	Status  int
	Message string
	Data    interface{}
	Error   error
	Title   string
}
