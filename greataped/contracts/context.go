package contracts

type (
	IContext interface {
		GUID() string
		Config() IConfig
		Storage() IStorage
		Request() IRequest
		Response() IResponse
		StringUtil() IStringUtil
		Nothing() error
		ParseJson(interface{}) IResult
		ParseBodyAndValidate(body interface{}) error
		GetUser() uint
		Redirect(location string, status ...int) error
		Render(name string, bind interface{}, layouts ...string) error

		String(interface{}) error
		Json(interface{}) error
		Activity(interface{}) error

		GetActivityStream(url, keyId, privateKey string, data []byte, output interface{}) error
		PostActivityStream(url, keyId, privateKey string, data []byte, output interface{}) error

		// Error(int, string, ...any) IServerError
		BadRequest(string, ...any) IServerError
		NotFound(string, ...any) IServerError
		InternalServerError(string, ...any) IServerError
		Unauthorized(string, ...any) IServerError
		Conflict(string, ...any) IServerError
	}
)
