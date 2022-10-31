package entity

import (
	"fmt"
	"reflect"
	"time"

	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts/model"
)

var IdentityEntityType = reflect.TypeOf(identityEntity{})

// noinspection GoUnusedExportedFunction
func InitializeIdentityEntity() {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
}

type identityEntity struct {
	entity
	UsernameField             string `json:"username" previous:"id" storage:"VARCHAR(32)" default:"''"`
	PhoneNumberField          string `json:"phone_number" previous:"username" storage:"VARCHAR(12)" default:"''"`
	PhoneNumberConfirmedField bool   `json:"phone_number_confirmed" previous:"phone_number" storage:"BIT(1)" default:"FALSE"`
	FirstNameField            string `json:"first_name" previous:"phone_number_confirmed" storage:"VARCHAR(128)" default:"''"`
	LastNameField             string `json:"last_name" previous:"first_name" storage:"VARCHAR(128)" default:"''"`
	DisplayNameField          string `json:"display_name" previous:"last_name" storage:"VARCHAR(128)" default:"''"`
	EmailField                string `json:"email" previous:"display_name" storage:"VARCHAR(128)" default:"''"`
	EmailConfirmedField       bool   `json:"email_confirmed" previous:"email" storage:"BIT(1)" default:"FALSE"`
	AvatarField               string `json:"avatar" previous:"email_confirmed" storage:"VARCHAR(512)" default:"''"`
	BannerField               string `json:"banner" previous:"avatar" storage:"VARCHAR(512)" default:"''"`
	SummaryField              string `json:"summary" previous:"banner" storage:"VARCHAR(512)" default:"''"`
	TokenField                string `json:"token" previous:"summary" storage:"VARCHAR(256)" default:"''"`
	MultiFactorField          bool   `json:"multi_factor" previous:"token" storage:"BIT(1)" default:"FALSE"`
	HashField                 string `json:"hash" previous:"multi_factor" storage:"VARCHAR(256)" default:"''"`
	SaltField                 string `json:"salt" previous:"hash" storage:"VARCHAR(64)" default:"''"`
	PublicKeyField            string `json:"public_key" previous:"salt" storage:"VARCHAR(4096)" default:"''"`
	PrivateKeyField           string `json:"private_key" previous:"public_key" storage:"VARCHAR(4096)" default:"''"`
	PermissionField           uint64 `json:"permission" previous:"private_key" storage:"BIGINT UNSIGNED" default:"0"`
	RestrictionField          uint32 `json:"restriction" previous:"permission" storage:"INT UNSIGNED" default:"0"`
	LastLoginField            int64  `json:"last_login" previous:"restriction" storage:"BIGINT" default:"0"`
	LoginCountField           uint32 `json:"login_count" previous:"last_login" storage:"INT UNSIGNED" default:"0"`
}

func NewIdentityEntity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) IIdentityEntity {
	return &identityEntity{
		entity:                    entity{IdField: id},
		UsernameField:             username,
		PhoneNumberField:          phoneNumber,
		PhoneNumberConfirmedField: phoneNumberConfirmed,
		FirstNameField:            firstName,
		LastNameField:             lastName,
		DisplayNameField:          displayName,
		EmailField:                email,
		EmailConfirmedField:       emailConfirmed,
		AvatarField:               avatar,
		BannerField:               banner,
		SummaryField:              summary,
		TokenField:                token,
		MultiFactorField:          multiFactor,
		HashField:                 hash,
		SaltField:                 salt,
		PublicKeyField:            publicKey,
		PrivateKeyField:           privateKey,
		PermissionField:           permission,
		RestrictionField:          restriction,
		LastLoginField:            lastLogin,
		LoginCountField:           loginCount,
	}
}

type identityPipeEntity struct {
	identityEntity
	pipeEntity
}

func NewIdentityPipeEntity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, source string, editor int64, payload string) IIdentityPipeEntity {
	return &identityPipeEntity{
		identityEntity: identityEntity{
			entity:                    entity{IdField: id, PayloadField: payload},
			UsernameField:             username,
			PhoneNumberField:          phoneNumber,
			PhoneNumberConfirmedField: phoneNumberConfirmed,
			FirstNameField:            firstName,
			LastNameField:             lastName,
			DisplayNameField:          displayName,
			EmailField:                email,
			EmailConfirmedField:       emailConfirmed,
			AvatarField:               avatar,
			BannerField:               banner,
			SummaryField:              summary,
			TokenField:                token,
			MultiFactorField:          multiFactor,
			HashField:                 hash,
			SaltField:                 salt,
			PublicKeyField:            publicKey,
			PrivateKeyField:           privateKey,
			PermissionField:           permission,
			RestrictionField:          restriction,
			LastLoginField:            lastLogin,
			LoginCountField:           loginCount,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_IDENTITY,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *identityEntity) Username() string {
	return entity.UsernameField
}

func (entity *identityEntity) PhoneNumber() string {
	return entity.PhoneNumberField
}

func (entity *identityEntity) PhoneNumberConfirmed() bool {
	return entity.PhoneNumberConfirmedField
}

func (entity *identityEntity) FirstName() string {
	return entity.FirstNameField
}

func (entity *identityEntity) LastName() string {
	return entity.LastNameField
}

func (entity *identityEntity) DisplayName() string {
	return entity.DisplayNameField
}

func (entity *identityEntity) Email() string {
	return entity.EmailField
}

func (entity *identityEntity) EmailConfirmed() bool {
	return entity.EmailConfirmedField
}

func (entity *identityEntity) Avatar() string {
	return entity.AvatarField
}

func (entity *identityEntity) Banner() string {
	return entity.BannerField
}

func (entity *identityEntity) Summary() string {
	return entity.SummaryField
}

func (entity *identityEntity) Token() string {
	return entity.TokenField
}

func (entity *identityEntity) MultiFactor() bool {
	return entity.MultiFactorField
}

func (entity *identityEntity) Hash() string {
	return entity.HashField
}

func (entity *identityEntity) Salt() string {
	return entity.SaltField
}

func (entity *identityEntity) PublicKey() string {
	return entity.PublicKeyField
}

func (entity *identityEntity) PrivateKey() string {
	return entity.PrivateKeyField
}

func (entity *identityEntity) Permission() uint64 {
	return entity.PermissionField
}

func (entity *identityEntity) Restriction() uint32 {
	return entity.RestrictionField
}

func (entity *identityEntity) LastLogin() int64 {
	return entity.LastLoginField
}

func (entity *identityEntity) LoginCount() uint32 {
	return entity.LoginCountField
}

func (entity *identityEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *identityEntity) String() string {
	return fmt.Sprintf("Identity (Id: %d, Username: %v, PhoneNumber: %v, PhoneNumberConfirmed: %v, FirstName: %v, LastName: %v, DisplayName: %v, Email: %v, EmailConfirmed: %v, Avatar: %v, Banner: %v, Summary: %v, Token: %v, MultiFactor: %v, Hash: %v, Salt: %v, PublicKey: %v, PrivateKey: %v, Permission: %v, Restriction: %v, LastLogin: %v, LoginCount: %v)", entity.Id(), entity.Username(), entity.PhoneNumber(), entity.PhoneNumberConfirmed(), entity.FirstName(), entity.LastName(), entity.DisplayName(), entity.Email(), entity.EmailConfirmed(), entity.Avatar(), entity.Banner(), entity.Summary(), entity.Token(), entity.MultiFactor(), entity.Hash(), entity.Salt(), entity.PublicKey(), entity.PrivateKey(), entity.Permission(), entity.Restriction(), entity.LastLogin(), entity.LoginCount())
}
