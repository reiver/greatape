package core

import (
	"fmt"

	. "github.com/reiver/greatape/components/contracts"
)

type spi struct {
}

func NewSpi() (ISpi, error) {
	instance := &spi{}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (spi *spi) Validate() error {
	return nil
}

func (spi *spi) String() string {
	return fmt.Sprintf("Spi (Id: %d)", 0)
}

//------------------------------------------------------------------------------

type spis struct {
	collection Spis
}

// NewSpis creates an empty collection of 'Spi' which is not thread-safe.
func NewSpis() ISpiCollection {
	return &spis{
		collection: make(Spis, 0),
	}
}

func (spis *spis) Count() int {
	return len(spis.collection)
}

func (spis *spis) IsEmpty() bool {
	return len(spis.collection) == 0
}

func (spis *spis) IsNotEmpty() bool {
	return len(spis.collection) > 0
}

func (spis *spis) HasExactlyOneItem() bool {
	return len(spis.collection) == 1
}

func (spis *spis) HasAtLeastOneItem() bool {
	return len(spis.collection) >= 1
}

func (spis *spis) First() ISpi {
	return spis.collection[0]
}

func (spis *spis) Append(spi ISpi) {
	spis.collection = append(spis.collection, spi)
}

func (spis *spis) Reverse() ISpiCollection {
	slice := spis.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	spis.collection = slice

	return spis
}

func (spis *spis) ForEach(iterator SpiIterator) {
	if iterator == nil {
		return
	}

	for _, value := range spis.collection {
		iterator(value)
	}
}

func (spis *spis) Array() Spis {
	return spis.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) SpiExists(id int64) bool {
	return dispatcher.conductor.SpiManager().Exists(id)
}

func (dispatcher *dispatcher) SpiExistsWhich(condition SpiCondition) bool {
	return dispatcher.conductor.SpiManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListSpis() ISpiCollection {
	return dispatcher.conductor.SpiManager().ListSpis(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachSpi(iterator SpiIterator) {
	dispatcher.conductor.SpiManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterSpis(predicate SpiFilterPredicate) ISpiCollection {
	return dispatcher.conductor.SpiManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapSpis(predicate SpiMapPredicate) ISpiCollection {
	return dispatcher.conductor.SpiManager().Map(predicate)
}

func (dispatcher *dispatcher) GetSpi(id int64) ISpi {
	if spi, err := dispatcher.conductor.SpiManager().GetSpi(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return spi
	}
}

func (dispatcher *dispatcher) AddSpi() ISpi {
	transaction := dispatcher.transaction
	if transaction != nil {
		if spi, err := dispatcher.conductor.SpiManager().AddSpiAtomic(transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	} else {
		if spi, err := dispatcher.conductor.SpiManager().AddSpi(dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	}
}

func (dispatcher *dispatcher) AddSpiWithCustomId(id int64) ISpi {
	transaction := dispatcher.transaction
	if transaction != nil {
		if spi, err := dispatcher.conductor.SpiManager().AddSpiWithCustomIdAtomic(id, transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	} else {
		if spi, err := dispatcher.conductor.SpiManager().AddSpiWithCustomId(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	}
}

func (dispatcher *dispatcher) LogSpi(source string, payload string) {
	dispatcher.conductor.SpiManager().Log(source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateSpi(id int64) ISpi {
	transaction := dispatcher.transaction
	if transaction != nil {
		if spi, err := dispatcher.conductor.SpiManager().UpdateSpiAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	} else {
		if spi, err := dispatcher.conductor.SpiManager().UpdateSpi(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateSpiObject(object IObject, spi ISpi) ISpi {
	transaction := dispatcher.transaction
	if transaction != nil {
		if spi, err := dispatcher.conductor.SpiManager().UpdateSpiAtomic(transaction, object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	} else {
		if spi, err := dispatcher.conductor.SpiManager().UpdateSpi(object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateSpiObject(object IObject, spi ISpi) ISpi {
	transaction := dispatcher.transaction
	if transaction != nil {
		if spi, err := dispatcher.conductor.SpiManager().AddOrUpdateSpiObjectAtomic(transaction, object.Id(), spi, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	} else {
		if spi, err := dispatcher.conductor.SpiManager().AddOrUpdateSpiObject(object.Id(), spi, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	}
}

func (dispatcher *dispatcher) RemoveSpi(id int64) ISpi {
	transaction := dispatcher.transaction
	if transaction != nil {
		if spi, err := dispatcher.conductor.SpiManager().RemoveSpiAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	} else {
		if spi, err := dispatcher.conductor.SpiManager().RemoveSpi(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return spi
		}
	}
}

func (dispatcher *dispatcher) Echo(document IDocument) (IEchoResult, error) {
	return dispatcher.conductor.SpiManager().Echo(document, dispatcher.identity)
}

func (dispatcher *dispatcher) GetServerConfiguration() (IGetServerConfigurationResult, error) {
	return dispatcher.conductor.SpiManager().GetServerConfiguration(dispatcher.identity)
}

func (dispatcher *dispatcher) CheckUsernameAvailability(username string) (ICheckUsernameAvailabilityResult, error) {
	return dispatcher.conductor.SpiManager().CheckUsernameAvailability(username, dispatcher.identity)
}

func (dispatcher *dispatcher) Signup(username string, email string, password string) (ISignupResult, error) {
	return dispatcher.conductor.SpiManager().Signup(username, email, password, dispatcher.identity)
}

func (dispatcher *dispatcher) ResendVerificationCode(email string) (IResendVerificationCodeResult, error) {
	return dispatcher.conductor.SpiManager().ResendVerificationCode(email, dispatcher.identity)
}

func (dispatcher *dispatcher) Verify(email string, token string, code string) (IVerifyResult, error) {
	return dispatcher.conductor.SpiManager().Verify(email, token, code, dispatcher.identity)
}

func (dispatcher *dispatcher) Login(email string, password string) (ILoginResult, error) {
	return dispatcher.conductor.SpiManager().Login(email, password, dispatcher.identity)
}

func (dispatcher *dispatcher) GetProfileByUser() (IGetProfileByUserResult, error) {
	return dispatcher.conductor.SpiManager().GetProfileByUser(dispatcher.identity)
}

func (dispatcher *dispatcher) UpdateProfileByUser(displayName string, avatar string, banner string, summary string, github string) (IUpdateProfileByUserResult, error) {
	return dispatcher.conductor.SpiManager().UpdateProfileByUser(displayName, avatar, banner, summary, github, dispatcher.identity)
}

func (dispatcher *dispatcher) ChangePassword(currentPassword string, newPassword string) (IChangePasswordResult, error) {
	return dispatcher.conductor.SpiManager().ChangePassword(currentPassword, newPassword, dispatcher.identity)
}

func (dispatcher *dispatcher) ResetPassword(usernameOrEmail string) (IResetPasswordResult, error) {
	return dispatcher.conductor.SpiManager().ResetPassword(usernameOrEmail, dispatcher.identity)
}

func (dispatcher *dispatcher) Logout() (ILogoutResult, error) {
	return dispatcher.conductor.SpiManager().Logout(dispatcher.identity)
}

func (dispatcher *dispatcher) Webfinger(resource string) (IWebfingerResult, error) {
	return dispatcher.conductor.SpiManager().Webfinger(resource, dispatcher.identity)
}

func (dispatcher *dispatcher) GetPackages() (IGetPackagesResult, error) {
	return dispatcher.conductor.SpiManager().GetPackages(dispatcher.identity)
}

func (dispatcher *dispatcher) GetActor(username string) (IGetActorResult, error) {
	return dispatcher.conductor.SpiManager().GetActor(username, dispatcher.identity)
}

func (dispatcher *dispatcher) FollowActor(username string, acct string) (IFollowActorResult, error) {
	return dispatcher.conductor.SpiManager().FollowActor(username, acct, dispatcher.identity)
}

func (dispatcher *dispatcher) AuthorizeInteraction(uri string) (IAuthorizeInteractionResult, error) {
	return dispatcher.conductor.SpiManager().AuthorizeInteraction(uri, dispatcher.identity)
}

func (dispatcher *dispatcher) GetFollowers(username string) (IGetFollowersResult, error) {
	return dispatcher.conductor.SpiManager().GetFollowers(username, dispatcher.identity)
}

func (dispatcher *dispatcher) GetFollowing(username string) (IGetFollowingResult, error) {
	return dispatcher.conductor.SpiManager().GetFollowing(username, dispatcher.identity)
}

func (dispatcher *dispatcher) PostToOutbox(username string, body []byte) (IPostToOutboxResult, error) {
	return dispatcher.conductor.SpiManager().PostToOutbox(username, body, dispatcher.identity)
}

func (dispatcher *dispatcher) GetOutbox(username string) (IGetOutboxResult, error) {
	return dispatcher.conductor.SpiManager().GetOutbox(username, dispatcher.identity)
}

func (dispatcher *dispatcher) PostToInbox(username string, body []byte) (IPostToInboxResult, error) {
	return dispatcher.conductor.SpiManager().PostToInbox(username, body, dispatcher.identity)
}

func (dispatcher *dispatcher) GetInbox(username string) (IGetInboxResult, error) {
	return dispatcher.conductor.SpiManager().GetInbox(username, dispatcher.identity)
}
