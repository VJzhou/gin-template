package configx

type Config interface {
	ReadSection(string, interface{}) error
	//GetConfig() (interface{}, error)
}
