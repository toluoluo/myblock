package block

import (
	"helper"
)

type Block struct {
	Index      int64     // 区块当前序列
	CTime      string    // 区块生成时间戳（格式化为 1900-01-01 00:00：00）
	PreHash    string    // 上一个block的hash
	Hash       string    // 当前区块的hash
	Merkle     string    // Merkle根
	DegreeDiff int64     // 难度系数
	Data       BlockData // 当前区块数据
}

func (b Block) CreateFirstBlock() Block {
	block := b.New(Block{}, BlockData{})
	block.PreHash = block.Hash
	return block
}

func (b Block) New(pre_block Block, data BlockData) Block {
	block := Block{
		Index:   pre_block.Index + 1,
		CTime:   helper.Sys_time(),
		PreHash: pre_block.Hash,
		Data:    data,
	}
	block.Hash = b.hashBlock(block)
	writeToBlockFile(block)
	return block
}

func (b Block) IsValidBlock(pre_block, curr_block Block) bool {
	if pre_block.Index+1 != curr_block.Index {
		return false
	}
	if pre_block.Hash != curr_block.PreHash {
		return false
	}
	if b.hashBlock(curr_block) != curr_block.Hash {
		return false
	}
	return true
}

func (b Block) hashBlock(block Block) string {
	str := string(block.Index) +
		block.CTime +
		block.PreHash +
		block.Data.DataHash
	return helper.Sys_sha256(str)
}
