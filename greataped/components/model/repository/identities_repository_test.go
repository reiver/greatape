package repository_test

import (
	"testing"

	. "rail.town/infrastructure/components/model/entity"
	. "rail.town/infrastructure/components/model/repository"
)

func TestIdentitiesRepository_Add(test *testing.T) {
	type arguments struct {
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
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:                   0,
				username:             "username",
				phoneNumber:          "phone_number",
				phoneNumberConfirmed: true,
				firstName:            "first_name",
				lastName:             "last_name",
				displayName:          "display_name",
				email:                "email",
				emailConfirmed:       true,
				avatar:               "avatar",
				banner:               "banner",
				summary:              "summary",
				token:                "token",
				multiFactor:          true,
				hash:                 "hash",
				salt:                 "salt",
				publicKey:            "public_key",
				privateKey:           "private_key",
				permission:           0,
				restriction:          0,
				lastLogin:            0,
				loginCount:           0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:                   0,
				username:             "username",
				phoneNumber:          "phone_number",
				phoneNumberConfirmed: true,
				firstName:            "first_name",
				lastName:             "last_name",
				displayName:          "display_name",
				email:                "email",
				emailConfirmed:       true,
				avatar:               "avatar",
				banner:               "banner",
				summary:              "summary",
				token:                "token",
				multiFactor:          true,
				hash:                 "hash",
				salt:                 "salt",
				publicKey:            "public_key",
				privateKey:           "private_key",
				permission:           0,
				restriction:          0,
				lastLogin:            0,
				loginCount:           0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:                   0,
				username:             "username",
				phoneNumber:          "phone_number",
				phoneNumberConfirmed: true,
				firstName:            "first_name",
				lastName:             "last_name",
				displayName:          "display_name",
				email:                "email",
				emailConfirmed:       true,
				avatar:               "avatar",
				banner:               "banner",
				summary:              "summary",
				token:                "token",
				multiFactor:          true,
				hash:                 "hash",
				salt:                 "salt",
				publicKey:            "public_key",
				privateKey:           "private_key",
				permission:           0,
				restriction:          0,
				lastLogin:            0,
				loginCount:           0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewIdentityEntity(testCase.arguments.id, testCase.arguments.username, testCase.arguments.phoneNumber, testCase.arguments.phoneNumberConfirmed, testCase.arguments.firstName, testCase.arguments.lastName, testCase.arguments.displayName, testCase.arguments.email, testCase.arguments.emailConfirmed, testCase.arguments.avatar, testCase.arguments.banner, testCase.arguments.summary, testCase.arguments.token, testCase.arguments.multiFactor, testCase.arguments.hash, testCase.arguments.salt, testCase.arguments.publicKey, testCase.arguments.privateKey, testCase.arguments.permission, testCase.arguments.restriction, testCase.arguments.lastLogin, testCase.arguments.loginCount)
			if result := Identities.Add(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.Add() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_FetchById(test *testing.T) {
	type arguments struct {
		id int64
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity, err := Identities.FetchById(testCase.arguments.id)
			if result := err == nil; result != testCase.expectation {
				test.Errorf("Identities.FetchById() = %v, expected %v", result, testCase.expectation)
			}

			_ = entity
		})
	}
}

func TestIdentitiesRepository_Update(test *testing.T) {
	type arguments struct {
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
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:                   0,
				username:             "username",
				phoneNumber:          "phone_number",
				phoneNumberConfirmed: true,
				firstName:            "first_name",
				lastName:             "last_name",
				displayName:          "display_name",
				email:                "email",
				emailConfirmed:       true,
				avatar:               "avatar",
				banner:               "banner",
				summary:              "summary",
				token:                "token",
				multiFactor:          true,
				hash:                 "hash",
				salt:                 "salt",
				publicKey:            "public_key",
				privateKey:           "private_key",
				permission:           0,
				restriction:          0,
				lastLogin:            0,
				loginCount:           0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:                   0,
				username:             "username",
				phoneNumber:          "phone_number",
				phoneNumberConfirmed: true,
				firstName:            "first_name",
				lastName:             "last_name",
				displayName:          "display_name",
				email:                "email",
				emailConfirmed:       true,
				avatar:               "avatar",
				banner:               "banner",
				summary:              "summary",
				token:                "token",
				multiFactor:          true,
				hash:                 "hash",
				salt:                 "salt",
				publicKey:            "public_key",
				privateKey:           "private_key",
				permission:           0,
				restriction:          0,
				lastLogin:            0,
				loginCount:           0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:                   0,
				username:             "username",
				phoneNumber:          "phone_number",
				phoneNumberConfirmed: true,
				firstName:            "first_name",
				lastName:             "last_name",
				displayName:          "display_name",
				email:                "email",
				emailConfirmed:       true,
				avatar:               "avatar",
				banner:               "banner",
				summary:              "summary",
				token:                "token",
				multiFactor:          true,
				hash:                 "hash",
				salt:                 "salt",
				publicKey:            "public_key",
				privateKey:           "private_key",
				permission:           0,
				restriction:          0,
				lastLogin:            0,
				loginCount:           0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewIdentityEntity(testCase.arguments.id, testCase.arguments.username, testCase.arguments.phoneNumber, testCase.arguments.phoneNumberConfirmed, testCase.arguments.firstName, testCase.arguments.lastName, testCase.arguments.displayName, testCase.arguments.email, testCase.arguments.emailConfirmed, testCase.arguments.avatar, testCase.arguments.banner, testCase.arguments.summary, testCase.arguments.token, testCase.arguments.multiFactor, testCase.arguments.hash, testCase.arguments.salt, testCase.arguments.publicKey, testCase.arguments.privateKey, testCase.arguments.permission, testCase.arguments.restriction, testCase.arguments.lastLogin, testCase.arguments.loginCount)
			if result := Identities.Update(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.Update() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_Remove(test *testing.T) {
	type arguments struct {
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
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			entity := NewIdentityEntity(testCase.arguments.id, testCase.arguments.username, testCase.arguments.phoneNumber, testCase.arguments.phoneNumberConfirmed, testCase.arguments.firstName, testCase.arguments.lastName, testCase.arguments.displayName, testCase.arguments.email, testCase.arguments.emailConfirmed, testCase.arguments.avatar, testCase.arguments.banner, testCase.arguments.summary, testCase.arguments.token, testCase.arguments.multiFactor, testCase.arguments.hash, testCase.arguments.salt, testCase.arguments.publicKey, testCase.arguments.privateKey, testCase.arguments.permission, testCase.arguments.restriction, testCase.arguments.lastLogin, testCase.arguments.loginCount)
			if result := Identities.Remove(entity, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.Remove() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_FetchAll(test *testing.T) {
	entities, err := Identities.FetchAll()
	if err != nil {
		test.Fatal(err)
	}

	_ = entities
}

func TestIdentitiesRepository_UpdateUsername(test *testing.T) {
	type arguments struct {
		id       int64
		username string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:       0,
				username: "username",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:       0,
				username: "username",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:       0,
				username: "username",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateUsername(testCase.arguments.id, testCase.arguments.username, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateUsername() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdatePhoneNumber(test *testing.T) {
	type arguments struct {
		id          int64
		phoneNumber string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:          0,
				phoneNumber: "phone_number",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:          0,
				phoneNumber: "phone_number",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:          0,
				phoneNumber: "phone_number",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdatePhoneNumber(testCase.arguments.id, testCase.arguments.phoneNumber, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdatePhoneNumber() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdatePhoneNumberConfirmed(test *testing.T) {
	type arguments struct {
		id                   int64
		phoneNumberConfirmed bool
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:                   0,
				phoneNumberConfirmed: true,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:                   0,
				phoneNumberConfirmed: true,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:                   0,
				phoneNumberConfirmed: true,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdatePhoneNumberConfirmed(testCase.arguments.id, testCase.arguments.phoneNumberConfirmed, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdatePhoneNumberConfirmed() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateFirstName(test *testing.T) {
	type arguments struct {
		id        int64
		firstName string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:        0,
				firstName: "first_name",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:        0,
				firstName: "first_name",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:        0,
				firstName: "first_name",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateFirstName(testCase.arguments.id, testCase.arguments.firstName, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateFirstName() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateLastName(test *testing.T) {
	type arguments struct {
		id       int64
		lastName string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:       0,
				lastName: "last_name",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:       0,
				lastName: "last_name",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:       0,
				lastName: "last_name",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateLastName(testCase.arguments.id, testCase.arguments.lastName, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateLastName() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateDisplayName(test *testing.T) {
	type arguments struct {
		id          int64
		displayName string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:          0,
				displayName: "display_name",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:          0,
				displayName: "display_name",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:          0,
				displayName: "display_name",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateDisplayName(testCase.arguments.id, testCase.arguments.displayName, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateDisplayName() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateEmail(test *testing.T) {
	type arguments struct {
		id    int64
		email string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:    0,
				email: "email",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:    0,
				email: "email",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:    0,
				email: "email",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateEmail(testCase.arguments.id, testCase.arguments.email, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateEmail() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateEmailConfirmed(test *testing.T) {
	type arguments struct {
		id             int64
		emailConfirmed bool
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:             0,
				emailConfirmed: true,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:             0,
				emailConfirmed: true,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:             0,
				emailConfirmed: true,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateEmailConfirmed(testCase.arguments.id, testCase.arguments.emailConfirmed, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateEmailConfirmed() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateAvatar(test *testing.T) {
	type arguments struct {
		id     int64
		avatar string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:     0,
				avatar: "avatar",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:     0,
				avatar: "avatar",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:     0,
				avatar: "avatar",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateAvatar(testCase.arguments.id, testCase.arguments.avatar, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateAvatar() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateBanner(test *testing.T) {
	type arguments struct {
		id     int64
		banner string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:     0,
				banner: "banner",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:     0,
				banner: "banner",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:     0,
				banner: "banner",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateBanner(testCase.arguments.id, testCase.arguments.banner, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateBanner() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateSummary(test *testing.T) {
	type arguments struct {
		id      int64
		summary string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:      0,
				summary: "summary",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:      0,
				summary: "summary",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:      0,
				summary: "summary",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateSummary(testCase.arguments.id, testCase.arguments.summary, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateSummary() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateToken(test *testing.T) {
	type arguments struct {
		id    int64
		token string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:    0,
				token: "token",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:    0,
				token: "token",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:    0,
				token: "token",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateToken(testCase.arguments.id, testCase.arguments.token, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateToken() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateMultiFactor(test *testing.T) {
	type arguments struct {
		id          int64
		multiFactor bool
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:          0,
				multiFactor: true,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:          0,
				multiFactor: true,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:          0,
				multiFactor: true,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateMultiFactor(testCase.arguments.id, testCase.arguments.multiFactor, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateMultiFactor() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateHash(test *testing.T) {
	type arguments struct {
		id   int64
		hash string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:   0,
				hash: "hash",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:   0,
				hash: "hash",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:   0,
				hash: "hash",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateHash(testCase.arguments.id, testCase.arguments.hash, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateHash() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateSalt(test *testing.T) {
	type arguments struct {
		id   int64
		salt string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:   0,
				salt: "salt",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:   0,
				salt: "salt",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:   0,
				salt: "salt",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateSalt(testCase.arguments.id, testCase.arguments.salt, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateSalt() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdatePublicKey(test *testing.T) {
	type arguments struct {
		id        int64
		publicKey string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:        0,
				publicKey: "public_key",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:        0,
				publicKey: "public_key",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:        0,
				publicKey: "public_key",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdatePublicKey(testCase.arguments.id, testCase.arguments.publicKey, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdatePublicKey() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdatePrivateKey(test *testing.T) {
	type arguments struct {
		id         int64
		privateKey string
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:         0,
				privateKey: "private_key",
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:         0,
				privateKey: "private_key",
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:         0,
				privateKey: "private_key",
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdatePrivateKey(testCase.arguments.id, testCase.arguments.privateKey, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdatePrivateKey() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdatePermission(test *testing.T) {
	type arguments struct {
		id         int64
		permission uint64
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:         0,
				permission: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:         0,
				permission: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:         0,
				permission: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdatePermission(testCase.arguments.id, testCase.arguments.permission, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdatePermission() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateRestriction(test *testing.T) {
	type arguments struct {
		id          int64
		restriction uint32
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:          0,
				restriction: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:          0,
				restriction: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:          0,
				restriction: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateRestriction(testCase.arguments.id, testCase.arguments.restriction, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateRestriction() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateLastLogin(test *testing.T) {
	type arguments struct {
		id        int64
		lastLogin int64
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:        0,
				lastLogin: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:        0,
				lastLogin: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:        0,
				lastLogin: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateLastLogin(testCase.arguments.id, testCase.arguments.lastLogin, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateLastLogin() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}

func TestIdentitiesRepository_UpdateLoginCount(test *testing.T) {
	type arguments struct {
		id         int64
		loginCount uint32
	}

	testCases := []struct {
		name        string
		expectation bool
		arguments   arguments
	}{
		{
			name:        "Case1",
			expectation: false,
			arguments: arguments{
				id:         0,
				loginCount: 0,
			},
		},
		{
			name:        "Case2",
			expectation: false,
			arguments: arguments{
				id:         0,
				loginCount: 0,
			},
		},
		{
			name:        "Case3",
			expectation: false,
			arguments: arguments{
				id:         0,
				loginCount: 0,
			},
		},
	}

	for _, testCase := range testCases {
		test.Run(testCase.name, func(test *testing.T) {
			if result := Identities.UpdateLoginCount(testCase.arguments.id, testCase.arguments.loginCount, -1) == nil; result != testCase.expectation {
				test.Errorf("Identities.UpdateLoginCount() = %v, expected %v", result, testCase.expectation)
			}
		})
	}
}
