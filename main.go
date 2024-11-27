package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"pqckpg/pqckpg_api"
)

func main() {

	// We define our output name for our keys
	var OUT_NAME_PRIVATE_KEY string = "private_key.key"
	var OUT_NAME_PUBLIC_KEY string = "public_key.key"

	// We generate the keys, the seed is NIL in this case
	pk, sk := pqckpg_api.GenerateKeys(nil)

	// We encrypt the keys for security and beauty. Who likes stupid strange characters?
	encodedPublicKey := encode(pk)
	encodedPrivateKey := encode(sk)

	// We save our keys for easier usage
	err := writeToFile(OUT_NAME_PUBLIC_KEY, encodedPublicKey)
	if err != nil {
		fmt.Println("Error writing public key:", err)
		return
	}

	err = writeToFile(OUT_NAME_PRIVATE_KEY, encodedPrivateKey)
	if err != nil {
		fmt.Println("Error writing private key:", err)
		return
	}

	// Let's extract them from our files
	publicKey, err := readFromFile(OUT_NAME_PUBLIC_KEY)
	if err != nil {
		fmt.Print("No public key found!")
		return
	}
	privateKey, err := readFromFile(OUT_NAME_PRIVATE_KEY)
	if err != nil {
		fmt.Print("No private key found!")
		return
	}

	// MATCH CHECK... Do the keys fit together?
	fmt.Println(pqckpg_api.Match(decode(string(publicKey)), decode(string(privateKey))))

	// SIGNING

	signMe := []byte("SIGN ME PLEASE!")
	signature := pqckpg_api.Sign(decode(string(privateKey)), signMe)
	isValid := pqckpg_api.Verify(decode(string(publicKey)), signMe, signature)

	if isValid {
		fmt.Println("Signature is valid!")
	} else {
		fmt.Println("Signature is invalid!")
	}

	// ENCRYPTING
	var message string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ 12345678910"

	encryptedMessage := pqckpg_api.Encrypt(decode(string(publicKey)), message)
	decryptedMessage := pqckpg_api.Decrypt(decode(string(privateKey)), encryptedMessage)

	fmt.Printf("Encrypted message: %s\n", encryptedMessage)
	fmt.Printf("Decrypted message: %s\n", decryptedMessage)

	// That's all. A quantum safe combination to encrypt and sign!
}

func encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func decode(data string) []byte {
	decoded, _ := base64.StdEncoding.DecodeString(data)
	return decoded
}

func readFromFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", file, err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", file, err)
	}

	return data, nil
}

func writeToFile(filename string, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", filename, err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %v", filename, err)
	}

	return nil
}
