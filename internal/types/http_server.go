package types

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
	}
)
