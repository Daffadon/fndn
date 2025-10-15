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
		MQInstanceType    string
		MQCloseConn       string
	}
)
