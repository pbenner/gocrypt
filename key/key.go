
package key

//import "fmt"
import "math"
import "encoding/binary"

type Key []byte

func Uint64Slice(k Key) []uint64 {

  const step int = 8;

  length := int(math.Ceil(float64(len(k))/float64(step)))
  tmp    := make([]byte,   step);
  result := make([]uint64, length)

  for i := 0; i < len(k); i += step {
    for j := 0; j < step; j++ {
      if i+j < len(k) {
        tmp[j] = k[i+j]
      } else {
        tmp[j] = 0
      }
    }
    result[i/step] = binary.LittleEndian.Uint64(tmp)
  }
  return (result)
}
