package string551

import (
	"errors"
	"strings"
	"unsafe"
)

var byteList []byte = StringToBytes("0123456789abcdef")

func StringToBytes(src string) []byte {
	return *(*[]byte)(unsafe.Pointer(&src))
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func HexBytesToString(hexBytes []byte) string {
	ret := make([]byte, 0, len(hexBytes)*2)
	for i := 0; i < len(hexBytes); i++ {
		ret = append(ret, byteList[hexBytes[i]/16])
		ret = append(ret, byteList[hexBytes[i]%16])
	}

	return BytesToString(ret)
}

func Join(src ...string) string {
	length := 0
	point := 0
	for i := 0; i < len(src); i++ {
		length += len(StringToBytes(src[i]))
	}
	ret := make([]byte, length)
	for i := 0; i < len(src); i++ {
		byt := StringToBytes(src[i])
		for j := 0; j < len(byt); j++ {
			ret[point] = byt[j]
			point++
		}
	}
	return BytesToString(ret)
}

func CamelCase(src string) string {
	bytes := StringToBytes(src)
	ret := make([]byte, 0, len(bytes))
	lowLine := true
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == 0x5F {
			lowLine = true
			continue
		}
		if lowLine {
			ret = append(ret, bytes[i]-0x20)
			lowLine = false
		} else {
			ret = append(ret, bytes[i])
		}
	}

	return BytesToString(ret)
}

func SnakeCase(src string) string {
	bytes := StringToBytes(src)

	ret := make([]byte, 0, len(bytes)+10)
	for i := 0; i < len(bytes); i++ {
		if 0x41 <= bytes[i] && bytes[i] <= 0x5A {
			// 0x41 - 0x5A => "A" - "Z"
			if i != 0 {
				ret = append(ret, 0x5F) // 0x5f => _
			}
			ret = append(ret, bytes[i]+0x20) // "A" => "a", "B" => "b"
		} else {
			ret = append(ret, bytes[i])
		}
	}

	return BytesToString(ret)
}

func Replace(src, old, new string) string {
	if src == old {
		return new
	}
	if new == "" {
		src = strings.Trim(src, old)
	}

	srcBytes := StringToBytes(src)
	oldBytes := StringToBytes(old)
	newBytes := StringToBytes(new)

	if len(srcBytes) == 0 {
		return src
	}
	if len(srcBytes) < len(oldBytes) {
		return src
	}

	retByte := make([]byte, 0, len(srcBytes))

	i := 0
	for i = 0; i <= len(srcBytes)-len(oldBytes); i++ {
		if srcBytes[i] == oldBytes[0] {
			cut := len(oldBytes)
			if BytesToString(oldBytes) == BytesToString(((srcBytes[i:])[:cut])) {
				retByte = append(retByte, newBytes...)
				i += len(oldBytes) - 1
			} else {
				retByte = append(retByte, srcBytes[i])
			}
		} else {
			retByte = append(retByte, srcBytes[i])
		}
	}
	retByte = append(retByte, srcBytes[i:]...)

	return BytesToString(retByte)

}

func Split(src, separate string) []string {
	return strings.Split(src, separate)
}

func Right(src string, length int) string {
	if len(src) < length {
		panic(errors.New("切り取り指定のバイト数が文字列長を超えました。"))
		return ""
	}

	return src[len(src)-length : len(src)]
}

func Lower(src string) string {
	bytes := StringToBytes(src)
	ret := make([]byte, 0, len(bytes))
	for i := 0; i < len(bytes); i++ {
		if 0x41 <= bytes[i] && bytes[i] <= 0x5A {
			// 0x41 - 0x5A => "A" - "Z"
			ret = append(ret, bytes[i]+0x20) // 0x61 - 0x7A => "a" - "z"
		} else {
			ret = append(ret, bytes[i])
		}
	}
	return BytesToString(ret)
}

func RightRune(src string, length int, suffix ...string) string {

	rSrc := []rune(src)
	if len(rSrc) <= length {
		return src
	}

	ret := ""
	for i := 0; i < length; i++ {
		ret = Join(ret, string(rSrc[i]))
	}

	if len(suffix) == 0 {
		return ret
	} else {
		return ret + suffix[0]
	}
}
