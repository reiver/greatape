package repository

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	. "github.com/xeronith/diamante/contracts/database"
	. "github.com/xeronith/diamante/contracts/logging"
)

type identitiesRepository struct {
	baseRepository
}

func newIdentitiesRepository(logger ILogger) IIdentitiesRepository {
	return &identitiesRepository{
		baseRepository: newBaseRepository("identity", "identities", IdentityEntityType, logger, false),
	}
}

func (repository *identitiesRepository) Add(entity IIdentityEntity, editor int64) error {
	if entity.Id() < 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "identities" ("id", "username", "phone_number", "phone_number_confirmed", "first_name", "last_name", "display_name", "email", "email_confirmed", "avatar", "banner", "summary", "token", "multi_factor", "hash", "salt", "public_key", "private_key", "permission", "restriction", "last_login", "login_count", "editor") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23);`
	return repository.database.InsertSingle(query, entity.Id(), entity.Username(), entity.PhoneNumber(), entity.PhoneNumberConfirmed(), entity.FirstName(), entity.LastName(), entity.DisplayName(), entity.Email(), entity.EmailConfirmed(), entity.Avatar(), entity.Banner(), entity.Summary(), entity.Token(), entity.MultiFactor(), entity.Hash(), entity.Salt(), entity.PublicKey(), entity.PrivateKey(), entity.Permission(), entity.Restriction(), entity.LastLogin(), entity.LoginCount(), editor)
}

func (repository *identitiesRepository) AddAtomic(transaction IRepositoryTransaction, entity IIdentityEntity, editor int64) error {
	if entity.Id() < 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `INSERT INTO "identities" ("id", "username", "phone_number", "phone_number_confirmed", "first_name", "last_name", "display_name", "email", "email_confirmed", "avatar", "banner", "summary", "token", "multi_factor", "hash", "salt", "public_key", "private_key", "permission", "restriction", "last_login", "login_count", "editor") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23);`
	return repository.database.InsertSingleAtomic(transaction, query, entity.Id(), entity.Username(), entity.PhoneNumber(), entity.PhoneNumberConfirmed(), entity.FirstName(), entity.LastName(), entity.DisplayName(), entity.Email(), entity.EmailConfirmed(), entity.Avatar(), entity.Banner(), entity.Summary(), entity.Token(), entity.MultiFactor(), entity.Hash(), entity.Salt(), entity.PublicKey(), entity.PrivateKey(), entity.Permission(), entity.Restriction(), entity.LastLogin(), entity.LoginCount(), editor)
}

func (repository *identitiesRepository) FetchById(id int64) (IIdentityEntity, error) {
	if id <= 0 {
		return nil, ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `SELECT "id", "username", "phone_number", "phone_number_confirmed" = TRUE, "first_name", "last_name", "display_name", "email", "email_confirmed" = TRUE, "avatar", "banner", "summary", "token", "multi_factor" = TRUE, "hash", "salt", "public_key", "private_key", "permission", "restriction", "last_login", "login_count" FROM "identities" WHERE "id" = $1 AND "status" = 0;`

	var identityEntity IIdentityEntity
	if err := repository.database.QuerySingle(func(cursor ICursor) error {
		var (
			id                   int64
			username             string
			phoneNumber          string
			phoneNumberConfirmed bool
			firstName            string
			lastName             string
			displayName          string
			email                string
			emailConfirmed       bool
			avatar               string
			banner               string
			summary              string
			token                string
			multiFactor          bool
			hash                 string
			salt                 string
			publicKey            string
			privateKey           string
			permission           uint64
			restriction          uint32
			lastLogin            int64
			loginCount           uint32
		)

		if err := cursor.Scan(&id, &username, &phoneNumber, &phoneNumberConfirmed, &firstName, &lastName, &displayName, &email, &emailConfirmed, &avatar, &banner, &summary, &token, &multiFactor, &hash, &salt, &publicKey, &privateKey, &permission, &restriction, &lastLogin, &loginCount); err != nil {
			return err
		}

		identityEntity = NewIdentityEntity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount)
		return nil
	}, query, id); err != nil {
		return nil, err
	}

	return identityEntity, nil
}

func (repository *identitiesRepository) Update(entity IIdentityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "username" = $1, "phone_number" = $2, "phone_number_confirmed" = $3, "first_name" = $4, "last_name" = $5, "display_name" = $6, "email" = $7, "email_confirmed" = $8, "avatar" = $9, "banner" = $10, "summary" = $11, "token" = $12, "multi_factor" = $13, "hash" = $14, "salt" = $15, "public_key" = $16, "private_key" = $17, "permission" = $18, "restriction" = $19, "last_login" = $20, "login_count" = $21, "editor" = $22 WHERE "id" = $23;`
	return repository.database.UpdateSingle(query, entity.Username(), entity.PhoneNumber(), entity.PhoneNumberConfirmed(), entity.FirstName(), entity.LastName(), entity.DisplayName(), entity.Email(), entity.EmailConfirmed(), entity.Avatar(), entity.Banner(), entity.Summary(), entity.Token(), entity.MultiFactor(), entity.Hash(), entity.Salt(), entity.PublicKey(), entity.PrivateKey(), entity.Permission(), entity.Restriction(), entity.LastLogin(), entity.LoginCount(), editor, entity.Id())
}

func (repository *identitiesRepository) UpdateAtomic(transaction IRepositoryTransaction, entity IIdentityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "username" = $1, "phone_number" = $2, "phone_number_confirmed" = $3, "first_name" = $4, "last_name" = $5, "display_name" = $6, "email" = $7, "email_confirmed" = $8, "avatar" = $9, "banner" = $10, "summary" = $11, "token" = $12, "multi_factor" = $13, "hash" = $14, "salt" = $15, "public_key" = $16, "private_key" = $17, "permission" = $18, "restriction" = $19, "last_login" = $20, "login_count" = $21, "editor" = $22 WHERE "id" = $23;`
	return repository.database.UpdateSingleAtomic(transaction, query, entity.Username(), entity.PhoneNumber(), entity.PhoneNumberConfirmed(), entity.FirstName(), entity.LastName(), entity.DisplayName(), entity.Email(), entity.EmailConfirmed(), entity.Avatar(), entity.Banner(), entity.Summary(), entity.Token(), entity.MultiFactor(), entity.Hash(), entity.Salt(), entity.PublicKey(), entity.PrivateKey(), entity.Permission(), entity.Restriction(), entity.LastLogin(), entity.LoginCount(), editor, entity.Id())
}

func (repository *identitiesRepository) Remove(entity IIdentityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingle(query, editor, entity.Id())
}

func (repository *identitiesRepository) RemoveAtomic(transaction IRepositoryTransaction, entity IIdentityEntity, editor int64) error {
	if entity.Id() <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "status" = 1, "editor" = $1 WHERE "id" = $2;`
	return repository.database.DeleteSingleAtomic(transaction, query, editor, entity.Id())
}

func (repository *identitiesRepository) FetchAll() (IdentityEntities, error) {
	// language=SQL
	query := `SELECT "id", "username", "phone_number", "phone_number_confirmed" = TRUE, "first_name", "last_name", "display_name", "email", "email_confirmed" = TRUE, "avatar", "banner", "summary", "token", "multi_factor" = TRUE, "hash", "salt", "public_key", "private_key", "permission", "restriction", "last_login", "login_count" FROM "identities" WHERE "id" > 0 AND "status" = 0;`

	var identityEntities IdentityEntities
	if err := repository.database.Query(func(cursor ICursor) error {
		var (
			id                   int64
			username             string
			phoneNumber          string
			phoneNumberConfirmed bool
			firstName            string
			lastName             string
			displayName          string
			email                string
			emailConfirmed       bool
			avatar               string
			banner               string
			summary              string
			token                string
			multiFactor          bool
			hash                 string
			salt                 string
			publicKey            string
			privateKey           string
			permission           uint64
			restriction          uint32
			lastLogin            int64
			loginCount           uint32
		)

		if err := cursor.Scan(&id, &username, &phoneNumber, &phoneNumberConfirmed, &firstName, &lastName, &displayName, &email, &emailConfirmed, &avatar, &banner, &summary, &token, &multiFactor, &hash, &salt, &publicKey, &privateKey, &permission, &restriction, &lastLogin, &loginCount); err != nil {
			return err
		}

		identityEntities = append(identityEntities, NewIdentityEntity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount))
		return nil
	}, query); err != nil {
		return nil, err
	}

	return identityEntities, nil
}

func (repository *identitiesRepository) UpdateUsername(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "username" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateUsernameAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "username" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePhoneNumber(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "phone_number" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePhoneNumberAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "phone_number" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePhoneNumberConfirmed(id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "phone_number_confirmed" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePhoneNumberConfirmedAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "phone_number_confirmed" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateFirstName(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "first_name" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateFirstNameAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "first_name" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateLastName(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "last_name" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateLastNameAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "last_name" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateDisplayName(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "display_name" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateDisplayNameAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "display_name" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateEmail(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "email" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateEmailAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "email" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateEmailConfirmed(id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "email_confirmed" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateEmailConfirmedAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "email_confirmed" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateAvatar(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "avatar" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateAvatarAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "avatar" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateBanner(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "banner" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateBannerAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "banner" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateSummary(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "summary" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateSummaryAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "summary" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateToken(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "token" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateTokenAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "token" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateMultiFactor(id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "multi_factor" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateMultiFactorAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "multi_factor" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateHash(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "hash" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateHashAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "hash" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateSalt(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "salt" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateSaltAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "salt" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePublicKey(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "public_key" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePublicKeyAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "public_key" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePrivateKey(id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "private_key" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePrivateKeyAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "private_key" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePermission(id int64, value uint64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "permission" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdatePermissionAtomic(transaction IRepositoryTransaction, id int64, value uint64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "permission" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateRestriction(id int64, value uint32, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "restriction" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateRestrictionAtomic(transaction IRepositoryTransaction, id int64, value uint32, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "restriction" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateLastLogin(id int64, value int64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "last_login" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateLastLoginAtomic(transaction IRepositoryTransaction, id int64, value int64, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "last_login" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}

func (repository *identitiesRepository) UpdateLoginCount(id int64, value uint32, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "login_count" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingle(query, value, editor, id)
}

func (repository *identitiesRepository) UpdateLoginCountAtomic(transaction IRepositoryTransaction, id int64, value uint32, editor int64) error {
	if id <= 0 {
		return ERROR_INVALID_PARAMETERS
	}

	// language=SQL
	query := `UPDATE "identities" SET "login_count" = $1, "editor" = $2 WHERE "id" = $3;`
	return repository.database.UpdateSingleAtomic(transaction, query, value, editor, id)
}
