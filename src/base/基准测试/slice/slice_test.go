package bench_test

import (
	"testing"
)

const TotalTimes = 1 << 16

func StaticCapacity() {

	// 提前一次性分配好slice所需内存空间，中间不需要再扩容，len为0，cap为1000000
	var s = make([]byte, 0, TotalTimes)

	for i := 0; i < TotalTimes; i++ {
		s = append(s, 0)
		//fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	}
}

func DynamicCapacity() {

	// 依赖slice底层自动扩容，中间会有很多次扩容，每次都从新分配一段新的内存空间，
	// 然后把数据拷贝到新的slice内存空间，然后释放旧空间，导致引发不必要的GC
	var s []byte

	for i := 0; i < TotalTimes; i++ {
		s = append(s, 0)
		//fmt.Printf("len = %d, cap = %d\n", len(s), cap(s
	}
}

func BenchmarkStaticCapacity(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StaticCapacity()
	}
}

func BenchmarkDynamicCapacity(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		DynamicCapacity()
	}
}
