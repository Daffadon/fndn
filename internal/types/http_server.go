package types

// None deselects an entity. Init* funcs treat it as "skip, generate nothing".
const None = "none"

type (
	HTTPServerParse struct {
		FrameworkImport   string
		FrameworkRouter   string
		RouterHandler     string
		DBInstanceType    string
		DBCloseConnection string
		DBImport          string
		MQImport          string
		MQInstance        string
		MQCloseConn       string
		CacheImport       string
		CacheInstanceType string
		CacheCloseConn    string
		HasDB             bool
		HasMQ             bool
		HasCache          bool
		ModuleName        string
	}
)
