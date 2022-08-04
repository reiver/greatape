package contracts

type IStorage interface {
	// Connect initiate the database connection
	Connect(path string)
	// Migrate migrates all the database tables
	Migrate(...interface{}) error
	Prepare(string) IQuery
}

type IQuery interface {
	Param(string) IResult
	Params(...string) IResult
}

type IResult interface {
	Get(string) any
	Set(string, string)
	Length() int
}
