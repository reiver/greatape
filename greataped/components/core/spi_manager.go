package core

import (
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/system"
	commands "rail.town/infrastructure/app/commands/spi"
	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
)

// noinspection GoSnakeCaseUsage
const SPI_MANAGER = "SpiManager"

type spiManager struct {
	systemComponent
	cache ICache
}

func newSpiManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) ISpiManager {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize

	manager := &spiManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *spiManager) Name() string {
	return SPI_MANAGER
}

func (manager *spiManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *spiManager) Load() error {
	return nil
}

func (manager *spiManager) Reload() error {
	return manager.Load()
}

func (manager *spiManager) OnCacheChanged(callback SpiCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *spiManager) Count() int {
	return manager.cache.Size()
}

func (manager *spiManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *spiManager) ExistsWhich(condition SpiCondition) bool {
	var spis Spis
	manager.ForEach(func(spi ISpi) {
		if condition(spi) {
			spis = append(spis, spi)
		}
	})

	return len(spis) > 0
}

func (manager *spiManager) ListSpis(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) ISpiCollection {
	return manager.Filter(SpiPassThroughFilter)
}

func (manager *spiManager) GetSpi(id int64, _ Identity) (ISpi, error) {
	if spi := manager.Find(id); spi == nil {
		return nil, ERROR_SPI_NOT_FOUND
	} else {
		return spi, nil
	}
}

func (manager *spiManager) AddSpi(editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiWithCustomId(id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiObject(spi ISpi, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiAtomic(transaction ITransaction, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiObjectAtomic(transaction ITransaction, spi ISpi, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) Log(source string, editor Identity, payload string) {
}

func (manager *spiManager) UpdateSpi(id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) UpdateSpiObject(id int64, spi ISpi, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) UpdateSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) UpdateSpiObjectAtomic(transaction ITransaction, id int64, spi ISpi, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddOrUpdateSpiObject(id int64, spi ISpi, editor Identity) (ISpi, error) {
	if manager.Exists(id) {
		return manager.UpdateSpiObject(id, spi, editor)
	} else {
		return manager.AddSpiObject(spi, editor)
	}
}

func (manager *spiManager) AddOrUpdateSpiObjectAtomic(transaction ITransaction, id int64, spi ISpi, editor Identity) (ISpi, error) {
	if manager.Exists(id) {
		return manager.UpdateSpiObjectAtomic(transaction, id, spi, editor)
	} else {
		return manager.AddSpiObjectAtomic(transaction, spi, editor)
	}
}

func (manager *spiManager) RemoveSpi(id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) RemoveSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) Find(id int64) ISpi {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(ISpi)
	}
}

func (manager *spiManager) ForEach(iterator SpiIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(ISpi))
	})
}

func (manager *spiManager) Filter(predicate SpiFilterPredicate) ISpiCollection {
	spis := NewSpis()
	if predicate == nil {
		return spis
	}

	manager.ForEach(func(spi ISpi) {
		if predicate(spi) {
			spis.Append(spi)
		}
	})

	return spis
}

func (manager *spiManager) Map(predicate SpiMapPredicate) ISpiCollection {
	spis := NewSpis()
	if predicate == nil {
		return spis
	}

	manager.ForEach(func(spi ISpi) {
		spis.Append(predicate(spi))
	})

	return spis
}

//region IEchoResult Implementation

type echoResult struct {
	document IDocument
}

func NewEchoResult(document IDocument, _ interface{}) IEchoResult {
	return &echoResult{
		document: document,
	}
}

func (result echoResult) Document() IDocument {
	return result.document
}

//endregion

func (manager *spiManager) Echo(document IDocument, editor Identity) (result IEchoResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(ECHO_REQUEST)
	defer editor.Unlock(ECHO_REQUEST)

	if result, err = commands.Echo(NewDispatcher(Conductor, editor), document); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
