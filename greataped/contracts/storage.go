package contracts

type (
	StorageType int

	IStorage interface {
		// Connect initiate the database connection
		Connect(path string)
		// Migrate migrates all the database tables
		Migrate(...interface{}) error
		Prepare(string) IQuery
	}

	IQuery interface {
		Param(string) IResult
		Params(...string) IResult
	}

	IResult interface {
		Get(string) any
		Set(string, string)
		Length() int
	}
)
