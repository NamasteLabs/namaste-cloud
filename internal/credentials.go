package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Credential structure to store cloud credentials.
type Credential struct {
	Cloud     string `json:"cloud"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

// GetCredentialFilePath returns the path to the encrypted credentials file.
func GetCredentialFilePath() (string, error) {
	configDir, err := GetUserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "credentials.enc"), nil
}

// GetKeyFilePath returns the path to the encryption key file.
func GetKeyFilePath() (string, error) {
	configDir, err := GetUserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "keyfile"), nil
}

// SaveCredential securely stores a cloud credential.
func SaveCredential(cred Credential) error {
	// Ensure the configuration directory exists.
	if err := EnsureConfigDir(); err != nil {
		return err
	}

	credentialFilePath, err := GetCredentialFilePath()
	if err != nil {
		return err
	}

	// Load existing credentials.
	creds, err := LoadAllCredentials()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to load credentials: %w", err)
	}

	// Initialize the map if it's nil.
	if creds == nil {
		creds = make(map[string]Credential)
	}

	// Add or update the credential for the given cloud.
	creds[cred.Cloud] = cred

	// Serialize credentials to JSON.
	data, err := json.Marshal(creds)
	if err != nil {
		return fmt.Errorf("failed to serialize credentials: %w", err)
	}

	// Encrypt the data.
	encryptedData, err := encrypt(data)
	if err != nil {
		return fmt.Errorf("failed to encrypt credentials: %w", err)
	}

	// Write encrypted data to file.
	return os.WriteFile(credentialFilePath, encryptedData, 0600)
}

// LoadAllCredentials decrypts and loads all stored credentials.
func LoadAllCredentials() (map[string]Credential, error) {
	credentialFilePath, err := GetCredentialFilePath()
	if err != nil {
		return nil, err
	}

	// Read the encrypted file.
	data, err := os.ReadFile(credentialFilePath)
	if err != nil {
		return nil, err
	}

	// Decrypt the data.
	decryptedData, err := decrypt(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt credentials: %w", err)
	}

	// Deserialize JSON to map.
	creds := make(map[string]Credential)
	if err := json.Unmarshal(decryptedData, &creds); err != nil {
		return nil, fmt.Errorf("failed to deserialize credentials: %w", err)
	}

	return creds, nil
}

// GetCredential retrieves a specific cloud's credential.
func GetCredential(cloud string) (Credential, error) {
	// Load all stored credentials.
	creds, err := LoadAllCredentials()
	if err != nil {
		return Credential{}, err
	}

	// Check if the requested cloud's credentials exist.
	cred, exists := creds[cloud]
	if !exists {
		return Credential{}, fmt.Errorf("no credentials found for cloud: %s", cloud)
	}

	return cred, nil
}

// Encrypt data using AES.
func encrypt(data []byte) ([]byte, error) {
	key, err := LoadKey()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

// Decrypt data using AES.
func decrypt(data []byte) ([]byte, error) {
	key, err := LoadKey()
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

// GenerateKey generates a new 32-byte encryption key and saves it to a file.
func GenerateKey() ([]byte, error) {
	// Ensure the configuration directory exists.
	if err := EnsureConfigDir(); err != nil {
		return nil, err
	}

	key := make([]byte, 32) // 32 bytes for AES-256
	if _, err := rand.Read(key); err != nil {
		return nil, fmt.Errorf("failed to generate encryption key: %w", err)
	}

	keyFilePath, err := GetKeyFilePath()
	if err != nil {
		return nil, err
	}

	// Save the key to a file.
	err = os.WriteFile(keyFilePath, key, 0600) // Secure file permissions
	if err != nil {
		return nil, fmt.Errorf("failed to save encryption key: %w", err)
	}

	return key, nil
}

// LoadKey loads the encryption key from the file.
func LoadKey() ([]byte, error) {
	keyFilePath, err := GetKeyFilePath()
	if err != nil {
		return nil, err
	}

	key, err := os.ReadFile(keyFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// If the key file doesn't exist, generate a new key.
			return GenerateKey()
		}
		return nil, fmt.Errorf("failed to read encryption key: %w", err)
	}

	if len(key) != 32 {
		return nil, fmt.Errorf("invalid encryption key size")
	}

	return key, nil
}
