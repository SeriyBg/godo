package storage

var config = readConfig()

func NewRepository() Repository {
	return &inFileRepository{
		fileName: config.FilePath,
	}
}
