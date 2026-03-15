package configs

func InitializeConfigs() error {
	Log = newLogger("gestao-financeira")

	initializeEnvVars()

	return nil
}
