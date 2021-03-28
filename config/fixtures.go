package config

import "crypto/tls"

func ParseCertificate() (tls.Certificate, error) {
	var (
		pem = []byte(`-----BEGIN CERTIFICATE-----
MIIDBzCCAnCgAwIBAgIJAIU6bgqkPGvsMA0GCSqGSIb3DQEBBQUAMGExCzAJBgNV
BAYTAlVTMQswCQYDVQQIFAIiIjELMAkGA1UEBxQCIiIxCzAJBgNVBAoUAiIiMQsw
CQYDVQQLFAIiIjELMAkGA1UEAxQCIiIxETAPBgkqhkiG9w0BCQEWAiIiMB4XDTA5
MDEwNDAzMTQzM1oXDTEwMDEwNDAzMTQzM1owYTELMAkGA1UEBhMCVVMxCzAJBgNV
BAgUAiIiMQswCQYDVQQHFAIiIjELMAkGA1UEChQCIiIxCzAJBgNVBAsUAiIiMQsw
CQYDVQQDFAIiIjERMA8GCSqGSIb3DQEJARYCIiIwgZ8wDQYJKoZIhvcNAQEBBQAD
gY0AMIGJAoGBAMXjPy2PmMIq73HqQCFUPwhinHs5Iv3agB8hPo1oz45rcJiVLB5O
eTlF9aPZIFSFeTb1CL6gpgOAYCHWvN747ehzApaEy7T/con0VkH2KPZrnwwd4Jsh
y4YI32vBitajUi/62FoshlINdS32FxGnF63CO9gPz7crLIrEzS2U5BV1AgMBAAGj
gcYwgcMwHQYDVR0OBBYEFABrEqK5EJDk5ej/7FwkRO7twWa3MIGTBgNVHSMEgYsw
gYiAFABrEqK5EJDk5ej/7FwkRO7twWa3oWWkYzBhMQswCQYDVQQGEwJVUzELMAkG
A1UECBQCIiIxCzAJBgNVBAcUAiIiMQswCQYDVQQKFAIiIjELMAkGA1UECxQCIiIx
CzAJBgNVBAMUAiIiMREwDwYJKoZIhvcNAQkBFgIiIoIJAIU6bgqkPGvsMAwGA1Ud
EwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADgYEAM7HQMQQXZ8pUcry3c1qPGyMlfcsj
rhub0pKACV0gJNJzb+dar57Q3VBhlr98LaEKxIj34MbDBDVvrNXR/VWrbJnHZnK4
cCLL04ynGBcuJS8zXFeCZw4p64F006NU+gi6h1AYq8UVac5KczvuEk0cYxGb302h
OA22HfvWuFvCENk=
-----END CERTIFICATE-----
`)
		key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQDF4z8tj5jCKu9x6kAhVD8IYpx7OSL92oAfIT6NaM+Oa3CYlSwe
Tnk5RfWj2SBUhXk29Qi+oKYDgGAh1rze+O3ocwKWhMu0/3KJ9FZB9ij2a58MHeCb
IcuGCN9rwYrWo1Iv+thaLIZSDXUt9hcRpxetwjvYD8+3KyyKxM0tlOQVdQIDAQAB
AoGAWUVcEfSuyCFQZcZ0adS0ntbFmv06oOR6WhDIREjdIXWslNjuzzk9jK3X07O2
1wpjlXxTFpQocHnwZDOYfsozoJc4AekGm1wVPYmjQCpUsXkV8Xz9GMrfU0JsigvB
GHDqfgBkB4Q38hv1KiLp1voDxn+qyKKjZyrT3a42R8FPE+ECQQDsYRG/zYcDpofJ
Lx2AwXNfGed8uWd+SVi/q9g3KSJpeaQGzaxfnroSd/g+0moGtZDk+iOG/0EbEL7k
nSl1fOZJAkEA1lBA/MlJrWlVx6NdUQVbQSvSWnT4FUkG8BpvfbZlF6Bk/3rWmVQN
U5WfbEPeJxvpJBND1doihR2nVaVND15FzQJAUZJN5bqvVPsq8Kppq/0WK0NtNwVk
SZhWIA7VVnPDhFKN4CspyPWlkKoF6OgD3rzZe6s2h2eeuBBXT91MaVbowQJAJwJa
oeidoZPvyjPhM3MvJhCs7EwoL++n9KJLMu21PvSyDZK1ZxlWh6VPbGx6DlJVQHzF
NzLKX8KDB+LbwPVe7QJAG4jzKY1r2zlMppZq12s1hd4cLD8Mjf/1weslPFZjqgPj
ECSHmNRzYkpROwGa2nPyzda74z43sxnZgpEH39CpgA==
-----END RSA PRIVATE KEY-----
`)
	)

	cert, err := tls.X509KeyPair(pem, key)
	if err != nil {
		return tls.Certificate{}, err
	}
	return cert, nil
}
