package xmlenc

import "testing"

func TestDecrypt(t *testing.T) {
	key := []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDU8wdiaFmPfTyRYuFlVPi866WrH/2JubkHzp89bBQopDaLXYxi
3PTu3O6Q/KaKxMOFBqrInwqpv/omOGZ4ycQ51O9I+Yc7ybVlW94lTo2gpGf+Y/8E
PsVbnZaFutRctJ4dVIp9aQ2TpLiGT0xX1OzBO/JEgq9GzDRf+B+eqSuglwIDAQAB
AoGBAMuy1eN6cgFiCOgBsB3gVDdTKpww87Qk5ivjqEt28SmXO13A1KNVPS6oQ8SJ
CT5Azc6X/BIAoJCURVL+LHdqebogKljhH/3yIel1kH19vr4E2kTM/tYH+qj8afUS
JEmArUzsmmK8ccuNqBcllqdwCZjxL4CHDUmyRudFcHVX9oyhAkEA/OV1OkjM3CLU
N3sqELdMmHq5QZCUihBmk3/N5OvGdqAFGBlEeewlepEVxkh7JnaNXAXrKHRVu/f/
fbCQxH+qrwJBANeQERF97b9Sibp9xgolb749UWNlAdqmEpmlvmS202TdcaaT1msU
4rRLiQN3X9O9mq4LZMSVethrQAdX1whawpkCQQDk1yGf7xZpMJ8F4U5sN+F4rLyM
Rq8Sy8p2OBTwzCUXXK+fYeXjybsUUMr6VMYTRP2fQr/LKJIX+E5ZxvcIyFmDAkEA
yfjNVUNVaIbQTzEbRlRvT6MqR+PTCefC072NF9aJWR93JimspGZMR7viY6IM4lrr
vBkm0F5yXKaYtoiiDMzlOQJADqmEwXl0D72ZG/2KDg8b4QZEmC9i5gidpQwJXUc6
hU+IVQoLxRq0fBib/36K9tcrrO5Ba4iEvDcNY+D8yGbUtA==
-----END RSA PRIVATE KEY-----
`)

	// This is an actual encrypted SAML response we got from testshib.org
	docStr := []byte(`<?xml version="1.0" encoding="UTF-8"?><saml2p:Response xmlns:saml2p="urn:oasis:names:tc:SAML:2.0:protocol" Destination="https://15661444.ngrok.io/saml2/acs" ID="_59c2431b4916af2018984213940ee675" InResponseTo="id-3d21faf29a101222d740735fa512f161" IssueInstant="2015-11-29T21:29:09.991Z" Version="2.0"><saml2:Issuer xmlns:saml2="urn:oasis:names:tc:SAML:2.0:assertion" Format="urn:oasis:names:tc:SAML:2.0:nameid-format:entity">https://idp.testshib.org/idp/shibboleth</saml2:Issuer><saml2p:Status><saml2p:StatusCode Value="urn:oasis:names:tc:SAML:2.0:status:Success"/></saml2p:Status><saml2:EncryptedAssertion xmlns:saml2="urn:oasis:names:tc:SAML:2.0:assertion"><xenc:EncryptedData xmlns:xenc="http://www.w3.org/2001/04/xmlenc#" Id="_62d64ca491a287a346272669fbe93191" Type="http://www.w3.org/2001/04/xmlenc#Element"><xenc:EncryptionMethod Algorithm="http://www.w3.org/2001/04/xmlenc#aes128-cbc" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"/><ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#"><xenc:EncryptedKey Id="_6d45c5b6b68d786f8322c0741f43a085" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"><xenc:EncryptionMethod Algorithm="http://www.w3.org/2001/04/xmlenc#rsa-oaep-mgf1p" xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"><ds:DigestMethod Algorithm="http://www.w3.org/2000/09/xmldsig#sha1" xmlns:ds="http://www.w3.org/2000/09/xmldsig#"/></xenc:EncryptionMethod><ds:KeyInfo><ds:X509Data><ds:X509Certificate>MIIB7zCCAVgCCQDFzbKIp7b3MTANBgkqhkiG9w0BAQUFADA8MQswCQYDVQQGEwJVUzELMAkGA1UE
CAwCR0ExDDAKBgNVBAoMA2ZvbzESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTEzMTAwMjAwMDg1MVoX
DTE0MTAwMjAwMDg1MVowPDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkdBMQwwCgYDVQQKDANmb28x
EjAQBgNVBAMMCWxvY2FsaG9zdDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1PMHYmhZj308
kWLhZVT4vOulqx/9ibm5B86fPWwUKKQ2i12MYtz07tzukPymisTDhQaqyJ8Kqb/6JjhmeMnEOdTv
SPmHO8m1ZVveJU6NoKRn/mP/BD7FW52WhbrUXLSeHVSKfWkNk6S4hk9MV9TswTvyRIKvRsw0X/gf
nqkroJcCAwEAATANBgkqhkiG9w0BAQUFAAOBgQCMMlIO+GNcGekevKgkakpMdAqJfs24maGb90Dv
TLbRZRD7Xvn1MnVBBS9hzlXiFLYOInXACMW5gcoRFfeTQLSouMM8o57h0uKjfTmuoWHLQLi6hnF+
cvCsEFiJZ4AbF+DgmO6TarJ8O05t8zvnOwJlNCASPZRH/JmF8tX0hoHuAQ==</ds:X509Certificate></ds:X509Data></ds:KeyInfo><xenc:CipherData xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"><xenc:CipherValue>0grSplyWOao1tEshQRtSsQqcl8lKTOqg/AR+U2Dh/1ACl0nZcv18De8U0iySrKSHQNaWcm2YpvBGUMddf4yKn40eVvmNoqElJVgOIhc5rPykua2AEyt2ShXOpFaCtXindqyax1IxxyJi+6o62swx+Q5pIy3YDaFN6/lNCgSdLak=</xenc:CipherValue></xenc:CipherData></xenc:EncryptedKey></ds:KeyInfo><xenc:CipherData xmlns:xenc="http://www.w3.org/2001/04/xmlenc#"><xenc:CipherValue>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</xenc:CipherValue></xenc:CipherData></xenc:EncryptedData></saml2:EncryptedAssertion></saml2p:Response>
`)

	expectedSignedString := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><Response xmlns=\"urn:oasis:names:tc:SAML:2.0:protocol\" xmlns:_xmlns=\"xmlns\" _xmlns:saml2p=\"urn:oasis:names:tc:SAML:2.0:protocol\" Destination=\"https://15661444.ngrok.io/saml2/acs\" ID=\"_59c2431b4916af2018984213940ee675\" InResponseTo=\"id-3d21faf29a101222d740735fa512f161\" IssueInstant=\"2015-11-29T21:29:09.991Z\" Version=\"2.0\"><Issuer xmlns=\"urn:oasis:names:tc:SAML:2.0:assertion\" _xmlns:saml2=\"urn:oasis:names:tc:SAML:2.0:assertion\" Format=\"urn:oasis:names:tc:SAML:2.0:nameid-format:entity\">https://idp.testshib.org/idp/shibboleth</Issuer><Status xmlns=\"urn:oasis:names:tc:SAML:2.0:protocol\"><StatusCode xmlns=\"urn:oasis:names:tc:SAML:2.0:protocol\" Value=\"urn:oasis:names:tc:SAML:2.0:status:Success\"></StatusCode></Status><EncryptedAssertion xmlns=\"urn:oasis:names:tc:SAML:2.0:assertion\" _xmlns:saml2=\"urn:oasis:names:tc:SAML:2.0:assertion\"><saml2:Assertion xmlns:saml2=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:xs=\"http://www.w3.org/2001/XMLSchema\" ID=\"_f6f518e2c236c9c558f7a8bc6387b103\" IssueInstant=\"2015-11-29T21:29:09.991Z\" Version=\"2.0\"><saml2:Issuer Format=\"urn:oasis:names:tc:SAML:2.0:nameid-format:entity\">https://idp.testshib.org/idp/shibboleth</saml2:Issuer><ds:Signature xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><ds:SignedInfo><ds:CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"></ds:CanonicalizationMethod><ds:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"></ds:SignatureMethod><ds:Reference URI=\"#_f6f518e2c236c9c558f7a8bc6387b103\"><ds:Transforms><ds:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"></ds:Transform><ds:Transform Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"><ec:InclusiveNamespaces xmlns:ec=\"http://www.w3.org/2001/10/xml-exc-c14n#\" PrefixList=\"xs\"></ec:InclusiveNamespaces></ds:Transform></ds:Transforms><ds:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"></ds:DigestMethod><ds:DigestValue>VwEKsGObmOM6y22Nstadwz1fq6dnQ2aDmERPMuEteds=</ds:DigestValue></ds:Reference></ds:SignedInfo><ds:SignatureValue>gcROTzJ7HgTu/LQprki8v9J5y4et2np48hYspgmygZRvRawzxfQDgB0MBvDIBG78J5XSd401g7E999JUEh4JtSMAig1THbeWhyITGHU1Vpl2xAR5Ma0vCMLjVIleeuFHhStFBNqKirNfulfhEa7Q5THVGKrVsNuIaP/yc10Gf8AyHfCIOf/ZQGiU3Srp/pKZLXPkSKTEZIq5tAOl+pA0maFBvb4+EkMPB6E66HiXknHL9KdNh8bPcq+EkqjhtHWOy341F8W9iy6MJYGuO9ksxdiY6FK5SqmPHlgoJqXx7Et2vYME6opIgFYB6m1KW6kWgVcF0VyIzJbkXq3yTi0b5g==</ds:SignatureValue><ds:KeyInfo><ds:X509Data><ds:X509Certificate>MIIEDjCCAvagAwIBAgIBADANBgkqhkiG9w0BAQUFADBnMQswCQYDVQQGEwJVUzEVMBMGA1UECBMM\nUGVubnN5bHZhbmlhMRMwEQYDVQQHEwpQaXR0c2J1cmdoMREwDwYDVQQKEwhUZXN0U2hpYjEZMBcG\nA1UEAxMQaWRwLnRlc3RzaGliLm9yZzAeFw0wNjA4MzAyMTEyMjVaFw0xNjA4MjcyMTEyMjVaMGcx\nCzAJBgNVBAYTAlVTMRUwEwYDVQQIEwxQZW5uc3lsdmFuaWExEzARBgNVBAcTClBpdHRzYnVyZ2gx\nETAPBgNVBAoTCFRlc3RTaGliMRkwFwYDVQQDExBpZHAudGVzdHNoaWIub3JnMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEArYkCGuTmJp9eAOSGHwRJo1SNatB5ZOKqDM9ysg7CyVTDClcp\nu93gSP10nH4gkCZOlnESNgttg0r+MqL8tfJC6ybddEFB3YBo8PZajKSe3OQ01Ow3yT4I+Wdg1tsT\npSge9gEz7SrC07EkYmHuPtd71CHiUaCWDv+xVfUQX0aTNPFmDixzUjoYzbGDrtAyCqA8f9CN2txI\nfJnpHE6q6CmKcoLADS4UrNPlhHSzd614kR/JYiks0K4kbRqCQF0Dv0P5Di+rEfefC6glV8ysC8dB\n5/9nb0yh/ojRuJGmgMWHgWk6h0ihjihqiu4jACovUZ7vVOCgSE5Ipn7OIwqd93zp2wIDAQABo4HE\nMIHBMB0GA1UdDgQWBBSsBQ869nh83KqZr5jArr4/7b+QazCBkQYDVR0jBIGJMIGGgBSsBQ869nh8\n3KqZr5jArr4/7b+Qa6FrpGkwZzELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTET\nMBEGA1UEBxMKUGl0dHNidXJnaDERMA8GA1UEChMIVGVzdFNoaWIxGTAXBgNVBAMTEGlkcC50ZXN0\nc2hpYi5vcmeCAQAwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQUFAAOCAQEAjR29PhrCbk8qLN5M\nFfSVk98t3CT9jHZoYxd8QMRLI4j7iYQxXiGJTT1FXs1nd4Rha9un+LqTfeMMYqISdDDI6tv8iNpk\nOAvZZUosVkUo93pv1T0RPz35hcHHYq2yee59HJOco2bFlcsH8JBXRSRrJ3Q7Eut+z9uo80JdGNJ4\n/SJy5UorZ8KazGj16lfJhOBXldgrhppQBb0Nq6HKHguqmwRfJ+WkxemZXzhediAjGeka8nz8Jjwx\npUjAiSWYKLtJhGEaTqCYxCCX2Dw+dOTqUzHOZ7WKv4JXPK5G/Uhr8K/qhmFT2nIQi538n6rVYLeW\nj8Bbnl+ev0peYzxFyF5sQA==</ds:X509Certificate></ds:X509Data></ds:KeyInfo></ds:Signature><saml2:Subject><saml2:NameID Format=\"urn:oasis:names:tc:SAML:2.0:nameid-format:transient\" NameQualifier=\"https://idp.testshib.org/idp/shibboleth\" SPNameQualifier=\"https://15661444.ngrok.io/saml2/metadata\">_5c425656721b41a6cfa4a9c96225e082</saml2:NameID><saml2:SubjectConfirmation Method=\"urn:oasis:names:tc:SAML:2.0:cm:bearer\"><saml2:SubjectConfirmationData Address=\"75.144.86.91\" InResponseTo=\"id-3d21faf29a101222d740735fa512f161\" NotOnOrAfter=\"2015-11-29T21:34:09.991Z\" Recipient=\"https://15661444.ngrok.io/saml2/acs\"></saml2:SubjectConfirmationData></saml2:SubjectConfirmation></saml2:Subject><saml2:Conditions NotBefore=\"2015-11-29T21:29:09.991Z\" NotOnOrAfter=\"2015-11-29T21:34:09.991Z\"><saml2:AudienceRestriction><saml2:Audience>https://15661444.ngrok.io/saml2/metadata</saml2:Audience></saml2:AudienceRestriction></saml2:Conditions><saml2:AuthnStatement AuthnInstant=\"2015-11-29T21:29:09.715Z\" SessionIndex=\"_57adf921604642bd4e1dce7f308734f0\"><saml2:SubjectLocality Address=\"75.144.86.91\"></saml2:SubjectLocality><saml2:AuthnContext><saml2:AuthnContextClassRef>urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport</saml2:AuthnContextClassRef></saml2:AuthnContext></saml2:AuthnStatement><saml2:AttributeStatement><saml2:Attribute FriendlyName=\"uid\" Name=\"urn:oid:0.9.2342.19200300.100.1.1\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">myself</saml2:AttributeValue></saml2:Attribute><saml2:Attribute FriendlyName=\"eduPersonAffiliation\" Name=\"urn:oid:1.3.6.1.4.1.5923.1.1.1.1\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">Member</saml2:AttributeValue><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">Staff</saml2:AttributeValue></saml2:Attribute><saml2:Attribute FriendlyName=\"eduPersonPrincipalName\" Name=\"urn:oid:1.3.6.1.4.1.5923.1.1.1.6\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">myself@testshib.org</saml2:AttributeValue></saml2:Attribute><saml2:Attribute FriendlyName=\"sn\" Name=\"urn:oid:2.5.4.4\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">And I</saml2:AttributeValue></saml2:Attribute><saml2:Attribute FriendlyName=\"eduPersonScopedAffiliation\" Name=\"urn:oid:1.3.6.1.4.1.5923.1.1.1.9\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">Member@testshib.org</saml2:AttributeValue><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">Staff@testshib.org</saml2:AttributeValue></saml2:Attribute><saml2:Attribute FriendlyName=\"givenName\" Name=\"urn:oid:2.5.4.42\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">Me Myself</saml2:AttributeValue></saml2:Attribute><saml2:Attribute FriendlyName=\"eduPersonEntitlement\" Name=\"urn:oid:1.3.6.1.4.1.5923.1.1.1.7\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">urn:mace:dir:entitlement:common-lib-terms</saml2:AttributeValue></saml2:Attribute><saml2:Attribute FriendlyName=\"cn\" Name=\"urn:oid:2.5.4.3\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">Me Myself And I</saml2:AttributeValue></saml2:Attribute><saml2:Attribute FriendlyName=\"eduPersonTargetedID\" Name=\"urn:oid:1.3.6.1.4.1.5923.1.1.1.10\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue><saml2:NameID Format=\"urn:oasis:names:tc:SAML:2.0:nameid-format:persistent\" NameQualifier=\"https://idp.testshib.org/idp/shibboleth\" SPNameQualifier=\"https://15661444.ngrok.io/saml2/metadata\">8F+M9ovyaYNwCId0pVkVsnZYRDo=</saml2:NameID></saml2:AttributeValue></saml2:Attribute><saml2:Attribute FriendlyName=\"telephoneNumber\" Name=\"urn:oid:2.5.4.20\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:uri\"><saml2:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:type=\"xs:string\">555-5555</saml2:AttributeValue></saml2:Attribute></saml2:AttributeStatement></saml2:Assertion>\x01</EncryptedAssertion></Response>\n"
	actualSignedString, err := Decrypt(key, docStr)
	if err != nil {
		t.Errorf("decrypt: %s", err)
		return
	}

	if string(actualSignedString) != expectedSignedString {
		t.Errorf("decrypt: expected %q, got `%q`", expectedSignedString, string(actualSignedString))
		return
	}
}
