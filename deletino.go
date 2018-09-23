package main

import (
	"bytes"
)

func DeleteComments(code []byte) []byte {

	for i := 0; i < len(code); i++ {
		if code[i] == byte('"') {
			for j := 1; j < len(code)-2; j++ {
				if bytes.Compare(code[i+j:i+j+2], []byte(`""`)) == 0 {
					j++
					continue
				}
				if code[i+j] == byte('"') {
					i += j
					break
				}
			}
			continue
		}
		if i+2 != len(code) {
			if bytes.Compare(code[i:i+2], []byte("//")) == 0 {
				for j := 2; i+j < len(code); j++ {
					if code[i+j] == byte('\n') || i+j == len(code)-1 {
						code = append(code[:i], code[i+j+1:]...)
						i--
						break
					}
				}
			} else {
				if i+2 != len(code) {
					if bytes.Compare(code[i:i+2], []byte("/*")) == 0 {
						for j := 2; j < len(code)-2; j++ {
							if bytes.Compare(code[i+j:i+j+2], []byte("*/")) == 0 {
								code = append(code[:i], code[i+j+2:]...)
								break
							}
						}
					}
				}
			}
		}
	}
	return code
}
