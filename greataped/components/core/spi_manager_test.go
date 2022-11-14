package core_test

import (
	"testing"

	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/core"
)

func TestSpiManager_GetName(test *testing.T) {
	manager := Conductor.SpiManager()

	if manager.Name() != SPI_MANAGER {
		test.Fail()
	}
}

func TestSpiManager_ResolveDependencies(test *testing.T) {
	manager := Conductor.SpiManager()

	if err := manager.ResolveDependencies(); err != nil {
		test.Fatal(err)
	}
}

func TestSpiManager_Load(test *testing.T) {
	manager := Conductor.SpiManager()

	if err := manager.Load(); err != nil {
		test.Fatal(err)
	}
}

func TestSpiManager_Reload(test *testing.T) {
	manager := Conductor.SpiManager()

	if err := manager.Reload(); err != nil && err != ERROR_OPERATION_NOT_SUPPORTED {
		test.Fatal(err)
	}
}

func TestSpiManager_Count(test *testing.T) {
	manager := Conductor.SpiManager()

	_ = manager.Count()
}

func TestSpiManager_Exists(test *testing.T) {
	manager := Conductor.SpiManager()

	if manager.Exists(0) {
		test.FailNow()
	}
}

func TestSpiManager_ListSpis(test *testing.T) {
	manager := Conductor.SpiManager()

	_ = manager.ListSpis(0, 0, "", nil)
}

func TestSpiManager_GetSpi(test *testing.T) {
	manager := Conductor.SpiManager()

	if spi, err := manager.GetSpi(0, nil); err == nil {
		_ = spi
		test.FailNow()
	}
}

func TestSpiManager_AddSpi(test *testing.T) {
	manager := Conductor.SpiManager()

	spi, err := manager.AddSpi(nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = spi
}

func TestSpiManager_UpdateSpi(test *testing.T) {
	manager := Conductor.SpiManager()

	spi, err := manager.UpdateSpi(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = spi
}

func TestSpiManager_RemoveSpi(test *testing.T) {
	manager := Conductor.SpiManager()

	spi, err := manager.RemoveSpi(0, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = spi
}

func TestSpiManager_Find(test *testing.T) {
	manager := Conductor.SpiManager()

	spi := manager.Find(0)
	if spi == nil {
		test.Fail()
	}

	_ = spi
}

func TestSpiManager_ForEach(test *testing.T) {
	manager := Conductor.SpiManager()

	manager.ForEach(func(spi ISpi) {
		_ = spi
	})
}

func TestSpiManager_Filter(test *testing.T) {
	manager := Conductor.SpiManager()

	spis := manager.Filter(func(spi ISpi) bool {
		return false
	})

	if spis.IsNotEmpty() {
		test.Fail()
	}

	_ = spis
}

func TestSpiManager_Map(test *testing.T) {
	manager := Conductor.SpiManager()

	spis := manager.Map(func(spi ISpi) ISpi {
		return spi
	})

	if spis.Count() != manager.Count() {
		test.Fail()
	}

	_ = spis
}

func TestSpiManager_Echo(test *testing.T) {
	manager := Conductor.SpiManager()

	result, err := manager.Echo(nil, nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = result
}

func TestSpiManager_Signup(test *testing.T) {
	manager := Conductor.SpiManager()

	result, err := manager.Signup("username", "email", "password", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = result
}

func TestSpiManager_Verify(test *testing.T) {
	manager := Conductor.SpiManager()

	result, err := manager.Verify("email", "token", "code", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = result
}

func TestSpiManager_Login(test *testing.T) {
	manager := Conductor.SpiManager()

	result, err := manager.Login("email", "password", nil)
	if err != nil {
		test.Fatal(err)
	}

	_ = result
}
