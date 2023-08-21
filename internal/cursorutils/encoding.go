package cursorutils

import "encoding/base64"

func EncodeCursor(cursor string) string {
	return base64.StdEncoding.EncodeToString([]byte(cursor))
}

func DecodeCursor(cursor string) (string, error) {
	decodedCursor, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return "", err
	}
	return string(decodedCursor), nil
}
