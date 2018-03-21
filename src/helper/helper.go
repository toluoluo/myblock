package helper
import (
	"time"
	"crypto/sha256"
	"encoding/hex"
	"os"
	// "log"
)

/**
* 格式化时间戳
 （格式：“2006-01-02 15:03:04”）
*/
func Sys_time() string{
	return time.Now().Format("2006-01-02 15:03:04")
	// return time.Now().String()
}

/*
* sha256哈希
*/
func Sys_sha256(str string) string{
	obj := sha256.New()
	obj.Write([]byte(str))
	return hex.EncodeToString(obj.Sum(nil))
}

func Sys_write_in_file(data, file_path string) (bool, error){
	f, err := os.OpenFile(file_path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil{
		// log.Fatal(err)
		return false, err
	}
	f.WriteString(data)
	f.Close()
	return true, nil
}

func Sys_file_is_exist(file_path string) (bool, os.FileInfo){
	 info, err := os.Stat(file_path)
	if err == nil{
		return true, info
	}
	return false, nil
}

