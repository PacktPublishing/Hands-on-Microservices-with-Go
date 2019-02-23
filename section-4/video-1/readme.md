# Packt Publishing - Hands on Microservices with Go
# Section 4 - Video 1 - HTTPS and TLS

## Disclaimer

The contents of the following video are intended as an **introduction** to securing an http endpoint with https. It is not intended as a definitive guide, if your application deals with sensitive data consider consulting the advice of a security professional. Also, consider that new vulnerabilities are discovered every day in security practices and software libraries, so the advice provided here might be obsolete and insecure by the time you read it.

## Generate Self Signed Certificates

```

cd ~
mkdir certs
cd certs

openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key

openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650 -subj "/C=AR/ST=CABA/L=CABA/O=Example Org/OU=IT Department/CN=*"

openssl genrsa -out client.key 2048
openssl ecparam -genkey -name secp384r1 -out client.key

openssl req -new -x509 -sha256 -key client.key -out client.pem -days 3650 -subj "/C=GB/ST=LONDON/L=LONDON/O=Another Org/OU=IT Department/CN=*"

```

**Warning:** Self signed certificates are useful for testing they should never be used in production. Use a CA signed certificate on production applications.


## Learn More

#### Public Key Cryptography

[Wikipedia - Public Key Cryptography](https://en.wikipedia.org/wiki/Public-key_cryptography)

[IBM - Public Key Cryptography](https://www.ibm.com/support/knowledgecenter/en/SSB23S_1.1.0.13/gtps7/s7pkey.html)

[IBM - Digital Signatures](https://www.ibm.com/support/knowledgecenter/SSB23S_1.1.0.13/gtps7/s7dsign.html)

#### TLS

[Mozilla Guide lines for TLS](https://wiki.mozilla.org/Security/Server_Side_TLS)

[Wikipedia - Public Key Certificate](https://en.wikipedia.org/wiki/Public_key_certificate)

[IBM - An overview of the TLS Handshake](https://www.ibm.com/support/knowledgecenter/en/SSFKSJ_7.1.0/com.ibm.mq.doc/sy10660_.htm)

[Tech Radar - How SSL and TLS Works](https://www.techradar.com/news/software/how-ssl-and-tls-works-1047412)

[The SSL/TLS Handshake: an Overview](https://www.ssl.com/article/ssl-tls-handshake-overview/)

[Caddy](https://caddyserver.com/)
