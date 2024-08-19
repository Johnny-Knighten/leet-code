package main

func reverseBits(num uint32) uint32 {
	num = (num&0xffff0000)>>16 | (num&0x0000ffff)<<16
	num = (num&0xff00ff00)>>8 | (num&0x00ff00ff)<<8
	num = (num&0xf0f0f0f0)>>4 | (num&0x0f0f0f0f)<<4
	num = (num&0xCCCCCCCC)>>2 | (num&0x33333333)<<2
	num = (num&0xAAAAAAAA)>>1 | (num&0x55555555)<<1
	return num
}
