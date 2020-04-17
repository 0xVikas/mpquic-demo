# MPQuic Demo

Uses [qdeconinck's mpquic](https://github.com/qdeconinck/mp-quic/) library

MPQUIC Setup instructions: https://multipath-quic.org/2017/12/09/artifacts-available.html

Install latest version of Golang

- Required changes/fixes:
> Replace the hardcoded url "quic.clemente.io" at quic-go/internal/handshake/crypto_setup_client.go in mp-quic project


 - To run the server:
  ```
  Generate the necessary pem files and then replace them with the existing ones at ./assets/certificates
  $ cd ~/go/src/github.com/0xVikas/mpquic-demo
  $ go run server.go
  ```
  - To run the proxy client:
  ```
  $ cd ~/go/src/github.com/0xVikas/mpquic-demo
  $ go run client.go
  ```
