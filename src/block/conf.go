package block

const (
	BLOCK_FILE_SIZE       = 100 * 1024 * 1024         // 运行时必须固定，固定后不能再修改
	BLOCK_FILE            = "blockdata/my.block"      // block前缀文件，存block数据
	CHECK_BLOCK_FILE      = "block/ckecklog.block"    // 校验文件日志，记录校验不正确的所在文件的block
	TMP_BLOCK_FILE        = "block/tmp.block"         // 待写入链的block数据	
	BLOCK_HASH_FILE       = "block/hash.block"        // 各个block文件的hash，用来校验分发
) 

/*********************
my.block
作用：存储block，当block写入判断文件大小超过100M时，另起一个block文件存储
文件名规则：my.block-1
		   my.block-2 ...

checklog.block
作用：存储校验失败的记录，全链校验正确文件不存在或为空
校验不正确
  1、记录校验所在的文件名、block数据，不在允许写入本地block、本地tmp_block
  2、不允许传递节点（即关闭节点，除非清空所有block，重新下载校验）

tmp.block
作用：本地新block，待其他节点确认有效后，写入本地节点，在之前不能再产生新block(待商榷)

*/
