package didcomm

import "encoding/json"

// Encrypt encrypts payload with server JWK, return cipher data
func Encrypt(plain []byte) (cipher []byte, err error) {
	// TODO @fangjian read server master key from env and encrypt plain data
	return plain, nil
}

func EncryptJSON(obj any) (cipher []byte, err error) {
	plain, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return Encrypt(plain)
}

// DecryptByClientID decrypts cipher data to plain data by client JWK
func DecryptByClientID(clientID string, cipher []byte) (plain []byte, err error) {
	// TODO @fangjian request did-vc service to decrypt cipher data by client JWK
	return cipher, nil
}
