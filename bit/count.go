package bit

const (
	m1  = 0x5555555555555555 //binary: 0101...
	m2  = 0x3333333333333333 //binary: 00110011..
	m4  = 0x0f0f0f0f0f0f0f0f //binary:  4 zeros,  4 ones ...
	m8  = 0x00ff00ff00ff00ff //binary:  8 zeros,  8 ones ...
	m16 = 0x0000ffff0000ffff //binary: 16 zeros, 16 ones ...
	m32 = 0x00000000ffffffff //binary: 32 zeros, 32 ones

	b32  = 0xFFFFFFFF
	ms1  = m1 & b32
	ms2  = m2 & b32
	ms4  = m4 & b32
	ms8  = m8 & b32
	ms16 = m16 & b32
	ms32 = m32 & b32
)

func Count(x uint64) int {
	x = (x & m1) + ((x >> 1) & m1)
	x = (x & m2) + ((x >> 2) & m2)
	x = (x & m4) + ((x >> 4) & m4)
	x = (x & m8) + ((x >> 8) & m8)
	x = (x & m16) + ((x >> 16) & m16)
	x = (x & m32) + ((x >> 32) & m32)
	return int(x)
}
