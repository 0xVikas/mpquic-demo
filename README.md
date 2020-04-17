# MPQuic Demo

Uses [qdeconinck's mpquic](https://github.com/qdeconinck/mp-quic/) library

MPQUIC Setup instructions: https://multipath-quic.org/2017/12/09/artifacts-available.html

Install latest version of Golang

- Required changes/fixes:
> Replace the hardcoded url "quic.clemente.io" at quic-go/internal/handshake/crypto_setup_client.go in mp-quic project


 - Using the example server from quic-go with mp-quic:
  ```
  Generate the necessary pem files and then replace them with the existing ones at ~/go/src/github.com/lucas-clemente/quic-go/example/.
  $ cd ~/go/src/github.com/lucas-clemente/quic-go
  $ go run example/main.go -www "path to folder with files to serve"
  ```
  - On the client side,
  ```
  cd ~/go/src/github.com/0xVikas/mpquic-client
  go run client.go
  ```
