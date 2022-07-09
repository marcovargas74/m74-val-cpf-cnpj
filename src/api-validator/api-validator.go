package m74validatorapi

const (
	versionPackage = "2022-07-04"
)

//IsMongoDBI VAR used to choise BD
var IsMongoDBI bool

//GetVersion Get the number version of packet
func GetVersion() string {
	return versionPackage
}
