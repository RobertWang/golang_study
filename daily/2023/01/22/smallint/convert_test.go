package smallint

import (
	// "test/smallint"
	"testing"
)

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// result := smallint.Convert(12)
		result := Convert(12)
		_ = result
	}
}
