//+build !noasm !appengine
// AUTO-GENERATED BY C2GOASM -- DO NOT EDIT

TEXT ·__resize_bilinear(SB), $32-48

	MOVQ dst+0(FP), DI
	MOVQ src+8(FP), SI
	MOVQ dst_h+16(FP), DX
	MOVQ dst_w+24(FP), CX
	MOVQ src_h+32(FP), R8
	MOVQ src_w+40(FP), R9
	ADDQ $8, SP

	WORD $0x854d; BYTE $0xc0
	JS   LBB0_1
	LONG $0x2afac1c4; BYTE $0xc0
	JMP  LBB0_3

LBB0_1:
	QUAD $0x8944e8d148c0894c; QUAD $0xc4c3094801e383c3
	QUAD $0xc058fac5c32afae1

LBB0_3:
	QUAD $0x4d01e38341d38941
	WORD $0xc985
	JS   LBB0_4
	LONG $0x2af2c1c4; BYTE $0xc9
	JMP  LBB0_6

LBB0_4:
	QUAD $0x8944ead149ca894d; QUAD $0xc4d3094c01e383cb
	QUAD $0xc958f2c5cb2af2e1

LBB0_6:
	QUAD $0x8548ebd148d38948
	BYTE $0xc9
	JS   LBB0_7
	LONG $0x2aeae1c4; BYTE $0xd1
	JMP  LBB0_9

LBB0_7:
	QUAD $0xc889ead149ca8949; QUAD $0xe1c4d0094c01e083
	LONG $0xc5d02aea; WORD $0x58ea; BYTE $0xd2

LBB0_9:
	LONG $0x48db0949; WORD $0xd285
	JS   LBB0_10
	LONG $0x2ae2e1c4; BYTE $0xda
	JMP  LBB0_12

LBB0_10:
	QUAD $0x58e2c5db2ae2c1c4
	BYTE $0xdb

LBB0_12:
	JE   LBB0_18
	WORD $0x8548; BYTE $0xc9
	JE   LBB0_18
	QUAD $0xca5ef2c5c35efac5; QUAD $0x0c244489ff408d41
	QUAD $0x08c78348ff798d45; QUAD $0x000000008d048d48
	QUAD $0x2444894840048d48
	LONG $0xf6314510

LBB0_15:
	QUAD $0x59eac5d62a5ac1c4; QUAD $0x4d9848c22cfac5d0
	QUAD $0x44470fc0394cc389; QUAD $0xaf0f4de0634c0c24
	LONG $0xf88948e1; WORD $0x3145; BYTE $0xed

LBB0_16:
	QUAD $0x59eac5d52a5ac1c4; QUAD $0xdb6348da2cfac5d1
	QUAD $0x48df470f41cb394c; QUAD $0x048d4ce3014cdb63
	QUAD $0x50894486148b465b; QUAD $0x588904865c8b42f8
	QUAD $0x188908865c8b42fc; QUAD $0x4c0cc08348c5ff49
	WORD $0xe939
	JNE  LBB0_16
	QUAD $0x10247c0348c6ff49
	LONG $0x4dd63949; WORD $0xd889
	JNE  LBB0_15

LBB0_18:
	SUBQ $8, SP
	RET
