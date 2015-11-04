package cipher

type Key []byte
type Message []byte

type Cipher interface
{
  Generate(n ...int)
  Encrypt(m Message) Message
  Decrypt(m Message) Message
}
