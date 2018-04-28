package queue

import "testing"

func TestLinkedMPSC(t *testing.T) {
	test(t, func(int) Queue { return NewLinkedMPSC() })
}

func BenchmarkLinkedMPSC(b *testing.B) {
	bench(b, func(int) Queue { return NewLinkedMPSC() })
}

func TestLinkedMPSCi(t *testing.T) {
	test(t, func(int) Queue { return NewLinkedMPSCIntrusive() })
}

func BenchmarkLinkedMPSCi(b *testing.B) {
	bench(b, func(int) Queue { return NewLinkedMPSCIntrusive() })
}
