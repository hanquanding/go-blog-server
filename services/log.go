/**
 * @author: hqd
 * @description: log
 * @file: log
 * @date: 2021-02-03 10:30
 */

package services

import "github.com/beego/beego/v2/core/logs"

func InitLogs() error {
	logsConf := `{"filename":"logs/apps.log", "level":7, "maxlines":0, "maxsize":0, "daily":true, "maxdays":10, "color":true}`
	err := logs.SetLogger(logs.AdapterFile, logsConf)
	if err != nil {
		return err
	}
	logs.Info("init log success")
	return nil
}
