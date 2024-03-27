package core

import "errors"

func readSimpleString(data []byte) (string, int, error) {
	pos := 1
	for ; data[pos] != '\r'; pos++ {
	}
	return string(data[1:pos]), pos + 2, nil
}

func Decode(data []byte) (interface{}, error) {
	if len(data) == 0 {
		return nil, errors.New("empty data provided")
	}
	res, _, err := readSimpleString(data)
	if err != nil {
		return nil, errors.New("could not decode")
	}
	return res, nil
}
