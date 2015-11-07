package onetimepad

//import "fmt"
import "math/rand"
import "time"
import "github.com/seehuhn/mt19937"

import . "github.com/pbenner/0100101101/key"
import . "github.com/pbenner/0100101101/message"

type OneTimePad struct
{
  privateKey Key
}

func (cipher *OneTimePad) Generate(n ...int) {

  length := 8;

  if len(n) > 1 {
    panic("Generate(): Invalid argument")
  }
  if len(n) == 1 {
    length = n[0]
  }

  cipher.privateKey= make(Key, length)

  rng := rand.New(mt19937.New())
  rng.Seed(time.Now().UnixNano())

  for i := range cipher.privateKey {
    cipher.privateKey[i] = byte(rng.Int())
  }
}

func (cipher *OneTimePad) Encrypt(m Message) Message {

  result := make(Message, len(m))
  length := len(cipher.privateKey)

  if (len(cipher.privateKey) != length) {
    panic("No private key available!")
  }
  mt  := mt19937.New()
  mt.SeedFromSlice(Uint64Slice(cipher.privateKey))
  rng := rand.New(mt)

  for i := 0; i < len(m); i++ {
    result[i] = m[i] ^ byte(rng.Int())
  }

  return result
}

func (cipher *OneTimePad) Decrypt(m Message) Message {

  return cipher.Encrypt(m)
}
