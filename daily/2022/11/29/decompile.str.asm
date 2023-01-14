main.main STEXT size=66 args=0x0 locals=0x18 funcid=0x0 align=0x0
	0x0000 00000 (str.go:3)	TEXT	main.main(SB), ABIInternal, $24-0
	0x0000 00000 (str.go:3)	CMPQ	SP, 16(R14)
	0x0004 00004 (str.go:3)	PCDATA	$0, $-2
	0x0004 00004 (str.go:3)	JLS	57
	0x0006 00006 (str.go:3)	PCDATA	$0, $-1
	0x0006 00006 (str.go:3)	SUBQ	$24, SP
	0x000a 00010 (str.go:3)	MOVQ	BP, 16(SP)
	0x000f 00015 (str.go:3)	LEAQ	16(SP), BP
	0x0014 00020 (str.go:3)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0014 00020 (str.go:3)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0014 00020 (str.go:6)	PCDATA	$1, $0
	0x0014 00020 (str.go:6)	CALL	runtime.printlock(SB)
	0x0019 00025 (str.go:6)	LEAQ	go.string."test"(SB), AX
	0x0020 00032 (str.go:6)	MOVL	$4, BX
	0x0025 00037 (str.go:6)	CALL	runtime.printstring(SB)
	0x002a 00042 (str.go:6)	CALL	runtime.printunlock(SB)
	0x002f 00047 (str.go:7)	MOVQ	16(SP), BP
	0x0034 00052 (str.go:7)	ADDQ	$24, SP
	0x0038 00056 (str.go:7)	RET
	0x0039 00057 (str.go:7)	NOP
	0x0039 00057 (str.go:3)	PCDATA	$1, $-1
	0x0039 00057 (str.go:3)	PCDATA	$0, $-2
	0x0039 00057 (str.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x003e 00062 (str.go:3)	PCDATA	$0, $-1
	0x003e 00062 (str.go:3)	NOP
	0x0040 00064 (str.go:3)	JMP	0
	0x0000 49 3b 66 10 76 33 48 83 ec 18 48 89 6c 24 10 48  I;f.v3H...H.l$.H
	0x0010 8d 6c 24 10 e8 00 00 00 00 48 8d 05 00 00 00 00  .l$......H......
	0x0020 bb 04 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48  ...............H
	0x0030 8b 6c 24 10 48 83 c4 18 c3 e8 00 00 00 00 66 90  .l$.H.........f.
	0x0040 eb be                                            ..
	rel 21+4 t=7 runtime.printlock+0
	rel 28+4 t=14 go.string."test"+0
	rel 38+4 t=7 runtime.printstring+0
	rel 43+4 t=7 runtime.printunlock+0
	rel 58+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
main..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
go.string."123" SRODATA dupok size=3
	0x0000 31 32 33                                         123
go.string."test" SRODATA dupok size=4
	0x0000 74 65 73 74                                      test
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
