package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"testing"
)

func Sign(content, prvKey []byte) (sign string, err error) {
	block, _ := pem.Decode(prvKey)
	if block == nil {
		fmt.Println("pem.Decode err")
		return
	}
	var private interface{}
	private, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	privateKey := private.(*rsa.PrivateKey)
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(content))
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey,
		crypto.SHA1, hashed)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(signature)
	return
}

func RSAVerify(origdata, ciphertext string, publicKey []byte) (bool, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return false, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(origdata))
	digest := h.Sum(nil)
	body, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA1, digest, body)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
Not verfied
*/
func TestSign(t *testing.T) {
	pubKey := "-----BEGIN PUBLIC KEY-----\n{MIIDTjCCAjYCCQDJakXiYJwLazANBgkqhkiG9w0BAQsFADBpMQswCQYDVQQGEwJjbjERMA8GA1UECAwIaGFuZ3pob3UxETAPBgNVBAcMCHpoZWppYW5nMQ4wDAYDVQQKDAVzdW5taTEOMAwGA1UECwwFc3VubWkxFDASBgNVBAMMCyouc3VubWkuY29tMB4XDTIyMDQxOTAzMzQyNVoXDTMyMDQxNjAzMzQyNVowaTELMAkGA1UEBhMCY24xETAPBgNVBAgMCGhhbmd6aG91MREwDwYDVQQHDAh6aGVqaWFuZzEOMAwGA1UECgwFc3VubWkxDjAMBgNVBAsMBXN1bm1pMRQwEgYDVQQDDAsqLnN1bm1pLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOCGgBbSagC1R4EAvWGTOIf/DeepNIyJVY8I6z9w66EEZ0YVJzFCLI/G/IDwSZb3qul1qirIJJVPORF8o32zpObJMtlDabdTGqyc50HPJiZDk5BFRIf9iInviWKgxYudF8wiJwCedhQl1cvw3PEmvyKMt4ZlGjN8EEdIAiTFw0tEZcn4xMLjZ7Ceh1nDhxUIoI42dcY1DqLc9sDdE2S3htA6SC8x4m41QbXHxSrrNq6qQLJNPDjD4Q1jrjKU2G+POqF0ySJ003ysJbxMZFBAJ5Z2rF/bJbiecvHuIlTw7b9K+D2we7cMgSIs64QjOJ9q/l8ENnmbqrqYm5ty3puxrhECAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAv5DOgRvdShqVtbtj2/ztHmOAQ51R9CGhFv5r94h5+SYuct5OaiJshMQHuWtrXCeH+OUDdan4qE8IG3jTrxoaL6biegKgdHQ8tpxDTgEB7Jhlta2LFbKcoKUDjMN7N3GhEarQksQU47enTE8WLc262If3dMYs3QzYWixlE3D+5qIbC5kBEYLiBaXQQxdewqcE6K1ujIGrWPghz0oIICZ52JD4U1JRqH8TNwz0CPXuxdPeMGBWJ9ggsMuik3Ez8BBfCCMCYXzogIzMJcX0YGvQFAmxn4GvO5F8F6cURC40iq96CG9KH7Xg6Tgg47ib478QClg0HkCCXJXh7iUQ1gH6OQ==}\n-----END PUBLIC KEY-----"
	prvKey := "-----BEGIN RSA PRIVATE KEY-----\n{MIIEowIBAAKCAQEA4IaAFtJqALVHgQC9YZM4h/8N56k0jIlVjwjrP3DroQRnRhUnMUIsj8b8gPBJlveq6XWqKsgklU85EXyjfbOk5sky2UNpt1MarJznQc8mJkOTkEVEh/2Iie+JYqDFi50XzCInAJ52FCXVy/Dc8Sa/Ioy3hmUaM3wQR0gCJMXDS0RlyfjEwuNnsJ6HWcOHFQigjjZ1xjUOotz2wN0TZLeG0DpILzHibjVBtcfFKus2rqpAsk08OMPhDWOuMpTYb486oXTJInTTfKwlvExkUEAnlnasX9sluJ5y8e4iVPDtv0r4PbB7twyBIizrhCM4n2r+XwQ2eZuqupibm3Lem7GuEQIDAQABAoIBAEdvI4OfUHCHPIezp41K3LqQEGl7MSfhbeJDMS2PDLi/AOiQRFbsuebIpX+Uc6VfiPYcJJbV9KW4feytXgrZRAbVTqNHSnQ0MZFnnkAW2wljiKhnEWW+6VkRaAGEKzW/NloRJ52PzPueCgaHzJPBAyDH6oAM3Kgyua8kHuJ6NSdVtAGH8Ajdh3XfyuKQOporWK9YlfPR/G3beAlKatGgBEbfioPLVLJ4EBzJA0FarDrYDNpmoHoo4GY21tZkgcPdSjI/fwpBnCHGVpE1X6a0ym8pUqSYhpMr3T71j/8TAZ+Z5Cw9o3PXxw3wu5k0rYjgugsQNd6wFNj/JlUxcSS1aFECgYEA9PUFElFl98JWoYisziyR9TWKFM10odQNXfAVp19ktph0ZJYY+td6chzTrfWZTJ+gFNJAxtr9FnCuXqMbCiBpPKxhzaXXtbbob/km2JGxNC9ljWuP4bT6K0vuBw4yXMA8molWimr0u2Rzl+ytiKMsaA5eYMtKFdHJOQwF+MNRnB0CgYEA6qWvFTsxCsVR2HRRYfDmHWJn+ggJu0BUrM5w0IMolGP5fKu/s7GiA+0p9tcNtIKE8E+iPx1A4p0cSEPEmVnztuzEGGdN3BUrj9rLRT7XbHXQs+tUTMWZAQFfYZUQJ5kYwXItbYQWjuaMrbjgPEyPrn3/gtIbwy7EkHyLk00Xb4UCgYANkVy8jQGm3X7K57UanmFfQZ3qVQ72v7YV0+x/HsuHSZ54y8+KZVEE7Q/UfNwG4HiPbq3j1dFa4tblqwceYnkxwSKRr3PpPr5VJWm/aSJ1j4KCeMi5abrJlyUSAvlLJeK1dJH0jMQNdRzp91QNU2xsPw9/MQNjfgE1RbM2+iqVCQKBgE0YpR8ntKRiUtL43Oh+O016UMmBLJleuLOSnNSV05Z2BrokwDbtbVs26GvXGwStQbqnn5p3JSOQFYPU6FquiHoY7xFJl/Zw6kA41kLpM+TKDQmgj7Et12jSJ6GrVYR9M/oTZsOt+692JtDJhrupOChP88zq9f46dpE2qrF6SfH1AoGBALGgfV2a1oPU78q/GD1jn45/Fb6xRw0rtLHmKqH4tGoKUmYQFY0+TZBa/wc1EB7816kBXimCP0KD8uC7U9D3bCUik/DKGy3obKqe9ataf0AlFuGauiXKkrOxLZOoa/V2I7zutohrl145STdDizUvbcqmPo7btlGQ7GkWImG4ULMX}\n-----END RSA PRIVATE KEY-----"
	content := "模型训练需要花多长时间"

	sign, err := Sign([]byte(content), []byte(prvKey))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("sign签名结果：", sign)

	res, err := RSAVerify(content, sign, []byte(pubKey))
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("验签结果：", res)
}
