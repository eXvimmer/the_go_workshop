package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	serverCert, serverKey, err := generate()
	if err != nil {
		fmt.Printf("error generating server certificate: %v", err)
		os.Exit(1)
	}
	fmt.Println("Server Certificate:")
	fmt.Printf("%s\n", serverCert)
	fmt.Println("Server (Private) Key:")
	fmt.Printf("%s\n", serverKey)
}

// generate returns the generated certificate, the private key and any
// potential error
func generate() ([]byte, []byte, error) {
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		return []byte{}, []byte{}, err
	}
	notBefore := time.Now()
	// create certificate template
	ca := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"example.com"},
		},
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(time.Hour * 24 * 365),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}
	// create the private key which will be used to sign the certificate
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return []byte{}, []byte{}, err
	}
	// create a self signed DER (binary encrypted) certificate
	DER, err := x509.CreateCertificate(rand.Reader, ca, ca, &rsaKey.PublicKey, rsaKey)
	if err != nil {
		return []byte{}, []byte{}, err
	}
	// convert the binary encoded DER certificate to an ASCII encoded PEM
	// (Privacy Enhanced Mail) cert.
	b := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: DER,
	}
	cert := pem.EncodeToMemory(&b)
	privateKey := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(rsaKey),
		})
	return cert, privateKey, nil
}
