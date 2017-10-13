package storage

func NewRepository() Repository {
	config := readConfig()
	return &inFileRepository{
		fileName: config.FilePath,
	}
}
