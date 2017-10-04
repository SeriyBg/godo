package storage

const fileName = "file_storage"

type Repository interface {
	AddNote(name string, description string) (err error)
	ShowAll() (notes []Note, err error)
	CompleteById(id string) (err error)
}

func GetRepository() Repository {
	return &inFileRepository{
		fileName: fileName,
	}
}
