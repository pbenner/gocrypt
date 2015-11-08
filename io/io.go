package io

import "strings"
import "bytes"

func CodeBlock(s string) string {

  var buffer bytes.Buffer

  reader := strings.NewReader(s)

  for i := 0;; i++ {

    if i != 0 && i % 80 == 0 {
      buffer.WriteByte('\n')
    }
    b, err := reader.ReadByte()
    if err != nil {
      break
    }
    buffer.WriteByte(b)

  }

  return (buffer.String())
}
