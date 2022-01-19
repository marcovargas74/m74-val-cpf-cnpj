package m74bankapi

const (
	versionPackage = "2022-01-15"
	exemplebool    = true
)

//usado para dizer se está em modo producao ou desenvolvimento
var isProd = false

//GetVersion Get the number version of packet
func GetVersion() string {
	return versionPackage
}

//SetLocalVar Set variable test
//default is false
func SetIsProduction(isInProduction bool) {
	isProd = isInProduction
}

//GetIsProduction Set variable Pruduction
func GetIsProduction() bool {
	return isProd
}
