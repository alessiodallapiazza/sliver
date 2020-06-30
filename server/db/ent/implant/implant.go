// Code generated by entc, DO NOT EDIT.

package implant

const (
	// Label holds the string label denoting the implant type in the database.
	Label = "implant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGOOS holds the string denoting the goos field in the database.
	FieldGOOS = "goos"
	// FieldGOARCH holds the string denoting the goarch field in the database.
	FieldGOARCH = "goarch"
	// FieldECCClientCert holds the string denoting the ecc_clientcert field in the database.
	FieldECCClientCert = "ecc_client_cert"
	// FieldECCClientKey holds the string denoting the ecc_clientkey field in the database.
	FieldECCClientKey = "ecc_client_key"
	// FieldRSACert holds the string denoting the rsa_cert field in the database.
	FieldRSACert = "rsa_cert"
	// FieldDebug holds the string denoting the debug field in the database.
	FieldDebug = "debug"
	// FieldObfuscateSymbols holds the string denoting the obfuscatesymbols field in the database.
	FieldObfuscateSymbols = "obfuscate_symbols"
	// FieldReconnectInterval holds the string denoting the reconnectinterval field in the database.
	FieldReconnectInterval = "reconnect_interval"
	// FieldMaxConnectionErrors holds the string denoting the maxconnectionerrors field in the database.
	FieldMaxConnectionErrors = "max_connection_errors"
	// FieldLimitDomainJoined holds the string denoting the limitdomainjoined field in the database.
	FieldLimitDomainJoined = "limit_domain_joined"
	// FieldLimitDatetime holds the string denoting the limitdatetime field in the database.
	FieldLimitDatetime = "limit_datetime"
	// FieldLimitHostname holds the string denoting the limithostname field in the database.
	FieldLimitHostname = "limit_hostname"
	// FieldLimitUsername holds the string denoting the limitusername field in the database.
	FieldLimitUsername = "limit_username"
	// FieldOutputFormat holds the string denoting the outputformat field in the database.
	FieldOutputFormat = "output_format"

	// Table holds the table name of the implant in the database.
	Table = "implants"
)

// Columns holds all SQL columns for implant fields.
var Columns = []string{
	FieldID,
	FieldGOOS,
	FieldGOARCH,
	FieldECCClientCert,
	FieldECCClientKey,
	FieldRSACert,
	FieldDebug,
	FieldObfuscateSymbols,
	FieldReconnectInterval,
	FieldMaxConnectionErrors,
	FieldLimitDomainJoined,
	FieldLimitDatetime,
	FieldLimitHostname,
	FieldLimitUsername,
	FieldOutputFormat,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Implant type.
var ForeignKeys = []string{
	"build_task_implant",
}
