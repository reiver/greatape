package commands

import . "github.com/reiver/greatape/components/contracts"

func GetServerConfiguration(x IDispatcher) (IGetServerConfigurationResult, error) {
	environment := ""
	if x.IsDevelopmentEnvironment() {
		environment = "development"
	} else if x.IsTestEnvironment() {
		environment = "test"
	} else if x.IsStagingEnvironment() {
		environment = "staging"
	} else if x.IsProductionEnvironment() {
		environment = "production"
	}

	return x.NewGetServerConfigurationResult(
		"GreatApe",
		environment,
		x.FQDN(),
	), nil
}
