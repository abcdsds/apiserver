package repository

import "fmt"

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

type XLogs []XLog

func (receiver XLogs) ConvertedToString() string {
	var converted string
	for _, xLog := range receiver {
		converted += fmt.Sprintf("ip:%s, elapsed:%s, objName:%s", xLog.ip, xLog.elapsed, xLog.objName)
	}

	return converted
}
