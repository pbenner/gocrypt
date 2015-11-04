package onetimepad

//import "fmt"
import "encoding/binary"
import "math/rand"
import "time"
import "github.com/seehuhn/mt19937"

import . "github.com/pbenner/0100101101/cipher"

const KeyLength = 8

type OneTimePad struct
{
  privateKey Key
}

func (cipher *OneTimePad) Generate(n ...int) {

  if len(n) != 0 {
    panic("Generate(): Invalid argument")
  }

  cipher.privateKey= make(Key, KeyLength)

  rng := rand.New(mt19937.New())
  rng.Seed(time.Now().UnixNano())

  for i := range cipher.privateKey {
    cipher.privateKey[i] = byte(rng.Int())
  }
}

func (cipher *OneTimePad) Encrypt(m Message) Message {

  result := make(Message, len(m))

  if (len(cipher.privateKey) != KeyLength) {
    panic("No private key available!")
  }
  rng := rand.New(mt19937.New())
  rng.Seed(int64(binary.LittleEndian.Uint64(cipher.privateKey[0:KeyLength])))

  for i := 0; i < len(m); i++ {
    result[i] = m[i] ^ byte(rng.Int())
  }

  return result
}

func (cipher *OneTimePad) Decrypt(m Message) Message {

  return cipher.Encrypt(m)
}
