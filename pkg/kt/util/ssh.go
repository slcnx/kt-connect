package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/alibaba/kt-connect/pkg/common"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ssh"
)

// SSHCredential ssh info
type SSHCredential struct {
	RemoteHost     string
	Port           string
	PrivateKeyPath string
}

// SSHGenerator ssh key pair generator
type SSHGenerator struct {
	PrivateKey, PublicKey []byte
	PrivateKeyPath        string
}

// NewSSHGenerator create ssh generator
func NewSSHGenerator(privateKey string, publicKey string, privateKeyPath string) *SSHGenerator {
	return &SSHGenerator{
		PrivateKey:     []byte(privateKey),
		PublicKey:      []byte(publicKey),
		PrivateKeyPath: privateKeyPath,
	}
}

// NewDefaultSSHCredential ...
func NewDefaultSSHCredential() *SSHCredential {
	return &SSHCredential{
		Port:       "2222",
		RemoteHost: "127.0.0.1",
	}
}

// Generate generate SSHGenerator
func Generate(privateKeyPath string) (*SSHGenerator, error) {
	privateKey, err := generatePrivateKey(common.SSHBitSize)
	if err != nil {
		return nil, err
	}

	publicKeyBytes, err := generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}
	privateKeyBytes := encodePrivateKeyToPEM(privateKey)

	sshKey := &SSHGenerator{
		PrivateKey:     privateKeyBytes,
		PrivateKeyPath: privateKeyPath,
		PublicKey:      publicKeyBytes,
	}
	err = WritePrivateKey(sshKey.PrivateKeyPath, sshKey.PrivateKey)
	return sshKey, err
}

// PrivateKeyPath ...
func PrivateKeyPath(component, identifier string) string {
	return fmt.Sprintf("%s/%s/"+common.SSHPrivateKeyName, KtHome, component, identifier)
}

// CleanRsaKeys ...
func CleanRsaKeys() {
	for _, c := range common.AllKtComponents {
		dir := fmt.Sprintf("%s/%s/", KtHome, c)
		files, _ := ioutil.ReadDir(dir)
		for _, f := range files {
			if strings.HasSuffix(f.Name(), common.PostfixRsaKey) {
				rsaKey := fmt.Sprintf("%s/%s", dir, f.Name())
				err := os.Remove(rsaKey)
				if err != nil {
					log.Debug().Msgf("Failed to remove rsa key file: %s", rsaKey)
				}
			}
		}
	}
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	log.Debug().Msg("Private Key generated")
	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)
	log.Debug().Msg("Public key generated")
	return pubKeyBytes, nil
}

// WritePrivateKey write ssh private key to privateKeyPath
func WritePrivateKey(privateKeyPath string, data []byte) error {
	dir := filepath.Dir(privateKeyPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0700); err != nil {
			log.Error().Err(err).Msgf("Can't create dir %s", dir)
			return err
		}
	}
	if err := ioutil.WriteFile(privateKeyPath, data, 0400); err != nil {
		log.Error().Err(err).Msgf("Write ssh private key to %s failed", privateKeyPath)
		return err
	}
	return nil
}
