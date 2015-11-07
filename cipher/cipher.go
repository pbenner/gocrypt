package cipher

import . "github.com/pbenner/0100101101/message"

type Cipher interface
{
  Generate(n ...int)
  Encrypt(m Message) Message
  Decrypt(m Message) Message
}
