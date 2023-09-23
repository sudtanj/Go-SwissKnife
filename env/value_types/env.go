package value_types

type EnvConst string

const (
	Local EnvConst = "local"
	Dev   EnvConst = "dev"
	Uat   EnvConst = "uat"
	Prod  EnvConst = "prod"
)
