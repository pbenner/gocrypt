package main

import   "fmt"
import . "github.com/pbenner/0100101101/cipher"
import . "github.com/pbenner/0100101101/key"
import . "github.com/pbenner/0100101101/message"
import . "github.com/pbenner/0100101101/onetimepad"
import . "github.com/pbenner/0100101101/io"


func f(cipher Cipher) Cipher {

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
  cipher.Generate(1024);

  fmt.Println(CodeBlock(KeyToString(cipher.GetKey())));

  f(&cipher)

}
