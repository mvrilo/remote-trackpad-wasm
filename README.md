# Remote Trackpad WASM

Remote trackpad built with go, websockets and wasm.

This is an unfinished yet working PoC implemented for Mac OS, using the CoreFoundation api. This was built as an experiment to mess around with cgo and wasm support in go.


## Building

`make`


## Certificate

Generate a certificate and install it on your mobile device. You can use [mkcert](https://github.com/FiloSottile/mkcert) for it. Make sure to [enable full trust](https://support.apple.com/en-nz/HT204477) after installing it.


## Example usage

```
./remote-trackpad-wasm \
  -addr 10.0.0.3:4430 \
  -cert ./cert.pem \
  -key ./key.pem
```

Access `https://10.0.0.3:4430` from a mobile device in your local network and drag around the webpage.
After the server start, the system will ask you to grant access to accessibility features.


## Author

Murilo Santana <<mvrilo@gmail.com>>
