package storage

var config = readConfig()

type Repository interface {
	AddNote(name string, description string) (err error)
	GetAll() (notes []Note, err error)
	CompleteById(id string) (err error)
}

func NewRepository() Repository {
	return &inFileRepository{
		fileName: config.FilePath,
	}
}
