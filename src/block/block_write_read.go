package block

import (
	"encoding/json"
	"helper"
	"log"
	"strconv"
)

/*
  将block数据写入临时文件
*/
func writeToTmpBlockFile(block Block) bool {
	return writeToFile(block, TMP_BLOCK_FILE, 0)
}

/*
  将block数据写入block文件
*/
func writeToBlockFile(block Block) bool {
	i := 1
	for {
		file := BLOCK_FILE + "-" + strconv.Itoa(i)
		is_succ := writeToFile(block, file, BLOCK_FILE_SIZE)
		if is_succ {
			return true
		}
		i++
	}
}

/**
* 将block数据写入文件
 */
func writeToFile(block Block, file string, size int64) bool {
	is_exist, info := helper.Sys_file_is_exist(file)
	if is_exist == false || info.Size() <= size {
		str := blockToJson(block)
		_, err := helper.Sys_write_in_file(str+"\n", file)
		if err != nil {
			log.Fatal("Write block in " + file + " error!")
			log.Fatal(err)
		}
		return true
	}
	return false
}

/*
  block数据结构转json
*/
func blockToJson(block Block) string {
	buf, err := json.Marshal(block)
	if err != nil {
		log.Fatal("block encoding to json error!")
	}
	return string(buf)
}

/*
  json数据转block结构
*/
func jsonToBlock(block string) Block {
	b := Block{}
	if err := json.Unmarshal(block, &b); err != nil {
		log.Fatal("json decoding to block error!")
	}
	return b
}
