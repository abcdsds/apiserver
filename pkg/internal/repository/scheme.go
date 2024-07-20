package repository

type XLog struct {
	ip        string
	useragent string
	objHash   string
	objName   string
	objType   string
	txid      string
	elapsed   string
	sql       string
	startTime string
	endTime   string
}

type Profile struct {
	step   string
	values []string
}

type ActiveService struct {
	objName string
	mode    string
	ipaddr  string
	elapsed string
}
