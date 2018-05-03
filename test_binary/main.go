package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"errors"
)

type Package struct {
	Begin byte
	Len uint32
	Seq uint32
	Data []byte
	End byte
}

func (pkg *Package) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	err := buf.WriteByte(pkg.Begin)

	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, pkg.Len)

	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, pkg.Seq)

	if err != nil {
		return nil, err
	}

	_, err = buf.Write(pkg.Data)

	if err != nil {
		return nil, err
	}

	err = buf.WriteByte(pkg.End)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (pkg *Package) Decode(data []byte) error {
	if len(data) < 10 {
		return errors.New("Package format invalid!")
	}

	pkg.Begin = data[0]
	pkg.Len = binary.BigEndian.Uint32(data[1:5])
	pkg.Seq = binary.BigEndian.Uint32(data[5:9])

	if pkg.Len != uint32(len(data) - 10) {
		return errors.New("Data length invalid!")
	}

	pkg.Data = data[9: len(data) - 1]
	pkg.End = data[len(data) - 1]

	return nil
}

func main() {
	data := `{"text":"hahaha"}`

	pkg := &Package{
		Begin: 0x4,
		Len: uint32(len(data)),
		Seq: 12345,
		Data: []byte(data),
		End: 0x5,
	}

	buf, err := pkg.Encode()

	if err != nil {
		fmt.Printf("Encode failed! error:%v\r\n", err.Error())
		return
	}

	fmt.Printf("Encode success! buf size:%v, buf:%v\r\n", len(buf), buf)

	pkgNew := &Package{}

	err = pkgNew.Decode(buf)

	if err != nil {
		fmt.Printf("Decode failed! error:%v\r\n", err.Error())
		return
	}

	fmt.Printf("Decode success! package:%v\r\n", pkgNew)
}
