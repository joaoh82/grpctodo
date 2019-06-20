package repository

type DB interface {
	Insert(collection string, items map[string]interface{}) error
	ListTasks(collection string) ([]interface{}, error)
}

// TODO - Package to be Implemented
