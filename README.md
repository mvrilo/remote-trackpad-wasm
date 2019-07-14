# Remote Trackpad WASM

Application for a remote trackpad using cgo, websockets and wasm.

This is an unfinished yet working PoC tested on Mac OS.


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

After the server start, the system will ask you to grant access to accessibility features.


## Author

Murilo Santana <<mvrilo@gmail.com>>
