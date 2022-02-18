package source

// ConfSource ...
type ConfSource interface {
	ReadConfig() (map[string]interface{}, error)
	ConfigChanged() <-chan map[string]interface{}
}
