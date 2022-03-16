package m74bankapi

const (
	versionPackage = "2022-03-16"
)

var isProd = false

//GetVersion Get the number version of packet
func GetVersion() string {
	return versionPackage
}

//SetIsProduction Set variable test
//default is false
func SetIsProduction(isInProduction bool) {
	isProd = isInProduction
}

//GetIsProduction Set variable Pruduction
func GetIsProduction() bool {
	return isProd
}
