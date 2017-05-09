package mapbench

import "testing"

func BenchmarkSyncMap(b *testing.B) {
	b.ResetTimer()
	makeSyncMap(b.N)
}
func BenchmarkConcurrentMap(b *testing.B) {
	b.ResetTimer()
	makeConcurrentMap(b.N)
}
