package main

import   "fmt"
import . "github.com/pbenner/0100101101/cipher"
import . "github.com/pbenner/0100101101/onetimepad"


func f(cipher Cipher) Cipher {

  cipher.Generate();
  m := Message("Hello World!")
  a := cipher.Encrypt(m)
  b := cipher.Decrypt(a)

  fmt.Println(string(m));
  fmt.Println(string(a));
  fmt.Println(string(b));

  return (cipher)

}

func main() {

  cipher := OneTimePad{};

  f(&cipher)

}
