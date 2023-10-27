package test

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"testing"
)

/* // 分片的大小
const chunkSize = 100 * 1024 * 1024 //100M */
// 分片的大小
const chunkSize = 10 * 1024 * 1024 //10M
// 测试文件的分片
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("img/test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	myFile, err := os.OpenFile("img/test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		//定位文件的起始位置
		myFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)
		f, err := os.OpenFile("./img/"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 分片文件的合并
func TestMergeChunk(t *testing.T) {
	myFile, err := os.OpenFile("img/test2.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	fileInfo, err := os.Stat("img/test.mp4")
	if err != nil {
		t.Fatal(err)
	}
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./img/"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := io.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myFile.Write(b)
		f.Close()
	}
	myFile.Close()

}

// 文件的一致性校验
func TestCheckHash(t *testing.T) {
	//获取第一个文件的信息
	file1, err := os.OpenFile("img/test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer file1.Close()
	b1, err := io.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}
	file2, err := os.OpenFile("img/test2.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer file2.Close()
	b2, err := io.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	str1 := fmt.Sprintf("%x", md5.Sum(b1))
	str2 := fmt.Sprintf("%x", md5.Sum(b2))
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str1 == str2)

}
