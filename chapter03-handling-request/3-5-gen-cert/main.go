package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publication Co."},
		OrganizationalUnit: []string{"book"},
		CommonName:         "go-web-programming",
	}
	
	template := x509.Certificate{
		SerialNumber:                serialNumber,
		Subject:                     subject,
		NotBefore:                   time.Now(),
		NotAfter:                    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:                    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		IPAddresses:                 []net.IP{net.ParseIP("127.0.0.1")},
	}
	
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	certOut, _ := os.Create("./cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()
	
	keyOut, _ := os.Create("./key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	keyOut.Close()
}
