package contracts

import (
	"activitypub"
	"app/models/domain"
	"mime/multipart"
)

type (
	IContext interface {
		GUID() string
		Config() IConfig
		Storage() IStorage
		Cache() ICache
		Request() IRequest
		Response() IResponse
		StringUtil() IStringUtil
		Nothing() error
		ParseJson(interface{}) IResult
		ParseBodyAndValidate(body interface{}) error
		SaveFile(file *multipart.FileHeader, path string) error
		GetUser() uint
		Redirect(location string, status ...int) error
		Render(name string, bind interface{}, layouts ...string) error

		String(interface{}) error
		Json(interface{}) error
		Activity(interface{}) error
		File(string) error

		GetActivityStream(url string, data []byte, output interface{}) error
		PostActivityStream(url string, data []byte, output interface{}) error
		GetActivityStreamSigned(url, keyId, privateKey string, data []byte, output interface{}) error
		PostActivityStreamSigned(url, keyId, privateKey string, data []byte, output interface{}) error

		GetWebFinger(username domain.Username) (*activitypub.Webfinger, error)
		GetActor(*activitypub.Webfinger) (*activitypub.Actor, error)
		GetOrderedCollection(url string) (*activitypub.OrderedCollection, error)

		BadRequest(interface{}, ...any) IServerError
		NotFound(interface{}, ...any) IServerError
		InternalServerError(interface{}, ...any) IServerError
		Unauthorized(interface{}, ...any) IServerError
		Conflict(interface{}, ...any) IServerError
	}
)
