package block

import (
	"helper"
)

type BlockData struct {
	DataHash string // 数据签名
	// Body     string //
	DCTime string // 数据生成时间
}

/*
* 数据体签名
 */
func (bd BlockData) New(body string) BlockData {
	blockData := BlockData{
		// Body : body,
		DCTime:   helper.Sys_time(),
		DataHash: helper.Sys_sha256(body),
	}
	return blockData
}
