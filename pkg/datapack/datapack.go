package datapack

import (
	"bytes"
	"encoding/binary"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	//DataLen uint32(4g个字节)
	return 4
}

// 封包
// 格式： datalen + Data
func (dp *DataPack) Pack(data []byte) ([]byte, error) {
	//创建bytes缓冲
	dataBuf := bytes.NewBuffer([]byte{})

	//写入dataLen
	if err := binary.Write(dataBuf, binary.LittleEndian, uint32(len(data))); err != nil {
		return nil, err
	}

	//写入data数据
	if err := binary.Write(dataBuf, binary.LittleEndian, data); err != nil {
		return nil, err
	}
	return dataBuf.Bytes(), nil
}

//解包
func (dp *DataPack) Unpack(data []byte) (left []byte, result []byte, err error) {
	//1.先获取前4个字节，得到data的长度
	head := data[:dp.GetHeadLen()]
	headBuf := bytes.NewReader(head)
	var dataLen uint32
	//读dataLen
	if err = binary.Read(headBuf, binary.LittleEndian, &dataLen); err != nil {
		return nil, nil, err
	}
	//2.根据data的长度再获取data内容
	return data[dp.GetHeadLen()+dataLen:], data[dp.GetHeadLen() : dp.GetHeadLen()+dataLen], nil
}
