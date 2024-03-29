rm *.pem

SERVER_CN=localhost

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "//CN=${SERVER_CN}"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "//CN=${SERVER_CN}"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "//CN=${SERVER_CN}"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text

# 6. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout email-server-key.pem -out email-server-req.pem -subj "//CN=${SERVER_CN}"

# 7. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in email-server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out email-server-cert.pem -extfile email-server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in email-server-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout yt-client-key.pem -out yt-client-req.pem -subj "//CN=${SERVER_CN}"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in yt-client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out yt-client-cert.pem -extfile yt-client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in yt-client-cert.pem -noout -text