package main


import (
	"crypto/x509"
	"fmt"
	"os"
)


func main(){
	certCerFile, err := os.Open("lvsoso.xyz.cer")
	checkError(err)
	derBytes := make([]byte, 1000)
	count, err := certCerFile.Read(derBytes)
	checkError(err)

	// 按照长度进行切分
	cert, err := x509.ParseCertificate(derBytes[0:count])
	checkError(err)

	fmt.Printf("Name %s\n", cert.Subject.CommonName)
	fmt.Printf("Not before %s\n", cert.NotBefore.String())
	fmt.Printf("Not after %s\n", cert.NotAfter.String())

}



func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
