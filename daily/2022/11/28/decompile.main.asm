main.increaseA STEXT size=172 args=0x0 locals=0x38 funcid=0x0 align=0x0
	0x0000 00000 (main.go:7)	TEXT	main.increaseA(SB), ABIInternal, $56-0
	0x0000 00000 (main.go:7)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:7)	PCDATA	$0, $-2
	0x0004 00004 (main.go:7)	JLS	162
	0x000a 00010 (main.go:7)	PCDATA	$0, $-1
	0x000a 00010 (main.go:7)	SUBQ	$56, SP
	0x000e 00014 (main.go:7)	MOVQ	BP, 48(SP)
	0x0013 00019 (main.go:7)	LEAQ	48(SP), BP
	0x0018 00024 (main.go:7)	MOVQ	$0, R13
	0x001f 00031 (main.go:7)	MOVQ	R13, 40(SP)
	0x0024 00036 (main.go:7)	FUNCDATA	$0, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0024 00036 (main.go:7)	FUNCDATA	$1, gclocals·wdmTuppZUxZYftR7OCq88Q==(SB)
	0x0024 00036 (main.go:7)	FUNCDATA	$2, main.increaseA.stkobj(SB)
	0x0024 00036 (main.go:7)	FUNCDATA	$4, main.increaseA.opendefer(SB)
	0x0024 00036 (main.go:7)	MOVB	$0, main..autotmp_4+7(SP)
	0x0029 00041 (main.go:7)	MOVQ	$0, main.~r0+8(SP)
	0x0032 00050 (main.go:8)	MOVQ	$0, main.i+16(SP)
	0x003b 00059 (main.go:9)	MOVUPS	X15, main..autotmp_3+24(SP)
	0x0041 00065 (main.go:9)	LEAQ	main.increaseA.func1(SB), AX
	0x0048 00072 (main.go:9)	MOVQ	AX, main..autotmp_3+24(SP)
	0x004d 00077 (main.go:9)	LEAQ	main.i+16(SP), AX
	0x0052 00082 (main.go:9)	MOVQ	AX, main..autotmp_3+32(SP)
	0x0057 00087 (main.go:9)	LEAQ	main..autotmp_3+24(SP), AX
	0x005c 00092 (main.go:9)	MOVQ	AX, main..autotmp_5+40(SP)
	0x0061 00097 (main.go:9)	MOVB	$1, main..autotmp_4+7(SP)
	0x0066 00102 (main.go:12)	MOVQ	main.i+16(SP), AX
	0x006b 00107 (main.go:12)	MOVQ	AX, main.~r0+8(SP)
	0x0070 00112 (main.go:12)	MOVB	$0, main..autotmp_4+7(SP)
	0x0075 00117 (main.go:12)	MOVQ	main..autotmp_5+40(SP), DX
	0x007a 00122 (main.go:12)	MOVQ	(DX), AX
	0x007d 00125 (main.go:12)	PCDATA	$1, $1
	0x007d 00125 (main.go:12)	CALL	AX
	0x007f 00127 (main.go:12)	MOVQ	main.~r0+8(SP), AX
	0x0084 00132 (main.go:12)	MOVQ	48(SP), BP
	0x0089 00137 (main.go:12)	ADDQ	$56, SP
	0x008d 00141 (main.go:12)	RET
	0x008e 00142 (main.go:12)	CALL	runtime.deferreturn(SB)
	0x0093 00147 (main.go:12)	MOVQ	main.~r0+8(SP), AX
	0x0098 00152 (main.go:12)	MOVQ	48(SP), BP
	0x009d 00157 (main.go:12)	ADDQ	$56, SP
	0x00a1 00161 (main.go:12)	RET
	0x00a2 00162 (main.go:12)	NOP
	0x00a2 00162 (main.go:7)	PCDATA	$1, $-1
	0x00a2 00162 (main.go:7)	PCDATA	$0, $-2
	0x00a2 00162 (main.go:7)	CALL	runtime.morestack_noctxt(SB)
	0x00a7 00167 (main.go:7)	PCDATA	$0, $-1
	0x00a7 00167 (main.go:7)	JMP	0
	0x0000 49 3b 66 10 0f 86 98 00 00 00 48 83 ec 38 48 89  I;f.......H..8H.
	0x0010 6c 24 30 48 8d 6c 24 30 49 c7 c5 00 00 00 00 4c  l$0H.l$0I......L
	0x0020 89 6c 24 28 c6 44 24 07 00 48 c7 44 24 08 00 00  .l$(.D$..H.D$...
	0x0030 00 00 48 c7 44 24 10 00 00 00 00 44 0f 11 7c 24  ..H.D$.....D..|$
	0x0040 18 48 8d 05 00 00 00 00 48 89 44 24 18 48 8d 44  .H......H.D$.H.D
	0x0050 24 10 48 89 44 24 20 48 8d 44 24 18 48 89 44 24  $.H.D$ H.D$.H.D$
	0x0060 28 c6 44 24 07 01 48 8b 44 24 10 48 89 44 24 08  (.D$..H.D$.H.D$.
	0x0070 c6 44 24 07 00 48 8b 54 24 28 48 8b 02 ff d0 48  .D$..H.T$(H....H
	0x0080 8b 44 24 08 48 8b 6c 24 30 48 83 c4 38 c3 e8 00  .D$.H.l$0H..8...
	0x0090 00 00 00 48 8b 44 24 08 48 8b 6c 24 30 48 83 c4  ...H.D$.H.l$0H..
	0x00a0 38 c3 e8 00 00 00 00 e9 54 ff ff ff              8.......T...
	rel 68+4 t=14 main.increaseA.func1+0
	rel 125+0 t=10 +0
	rel 143+4 t=7 runtime.deferreturn+0
	rel 163+4 t=7 runtime.morestack_noctxt+0
main.increaseA.func1 STEXT nosplit size=8 args=0x0 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (main.go:9)	TEXT	main.increaseA.func1(SB), NOSPLIT|NEEDCTXT|ABIInternal, $0-0
	0x0000 00000 (main.go:9)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (main.go:9)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (main.go:9)	MOVQ	8(DX), AX
	0x0004 00004 (main.go:10)	INCQ	(AX)
	0x0007 00007 (main.go:11)	RET
	0x0000 48 8b 42 08 48 ff 00 c3                          H.B.H...
main.increaseB STEXT size=148 args=0x0 locals=0x30 funcid=0x0 align=0x0
	0x0000 00000 (main.go:15)	TEXT	main.increaseB(SB), ABIInternal, $48-0
	0x0000 00000 (main.go:15)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:15)	PCDATA	$0, $-2
	0x0004 00004 (main.go:15)	JLS	138
	0x000a 00010 (main.go:15)	PCDATA	$0, $-1
	0x000a 00010 (main.go:15)	SUBQ	$48, SP
	0x000e 00014 (main.go:15)	MOVQ	BP, 40(SP)
	0x0013 00019 (main.go:15)	LEAQ	40(SP), BP
	0x0018 00024 (main.go:15)	MOVQ	$0, R13
	0x001f 00031 (main.go:15)	MOVQ	R13, 32(SP)
	0x0024 00036 (main.go:15)	FUNCDATA	$0, gclocals·J5F+7Qw7O7ve2QcWC7DpeQ==(SB)
	0x0024 00036 (main.go:15)	FUNCDATA	$1, gclocals·wdmTuppZUxZYftR7OCq88Q==(SB)
	0x0024 00036 (main.go:15)	FUNCDATA	$2, main.increaseB.stkobj(SB)
	0x0024 00036 (main.go:15)	FUNCDATA	$4, main.increaseB.opendefer(SB)
	0x0024 00036 (main.go:15)	MOVB	$0, main..autotmp_3+7(SP)
	0x0029 00041 (main.go:15)	MOVQ	$0, main.r+8(SP)
	0x0032 00050 (main.go:16)	MOVUPS	X15, main..autotmp_2+16(SP)
	0x0038 00056 (main.go:16)	LEAQ	main.increaseB.func1(SB), AX
	0x003f 00063 (main.go:16)	MOVQ	AX, main..autotmp_2+16(SP)
	0x0044 00068 (main.go:16)	LEAQ	main.r+8(SP), AX
	0x0049 00073 (main.go:16)	MOVQ	AX, main..autotmp_2+24(SP)
	0x004e 00078 (main.go:16)	LEAQ	main..autotmp_2+16(SP), AX
	0x0053 00083 (main.go:16)	MOVQ	AX, main..autotmp_4+32(SP)
	0x0058 00088 (main.go:19)	MOVB	$0, main..autotmp_3+7(SP)
	0x005d 00093 (main.go:19)	MOVQ	main..autotmp_4+32(SP), DX
	0x0062 00098 (main.go:19)	MOVQ	(DX), AX
	0x0065 00101 (main.go:19)	PCDATA	$1, $1
	0x0065 00101 (main.go:19)	CALL	AX
	0x0067 00103 (main.go:19)	MOVQ	main.r+8(SP), AX
	0x006c 00108 (main.go:19)	MOVQ	40(SP), BP
	0x0071 00113 (main.go:19)	ADDQ	$48, SP
	0x0075 00117 (main.go:19)	RET
	0x0076 00118 (main.go:19)	CALL	runtime.deferreturn(SB)
	0x007b 00123 (main.go:19)	MOVQ	main.r+8(SP), AX
	0x0080 00128 (main.go:19)	MOVQ	40(SP), BP
	0x0085 00133 (main.go:19)	ADDQ	$48, SP
	0x0089 00137 (main.go:19)	RET
	0x008a 00138 (main.go:19)	NOP
	0x008a 00138 (main.go:15)	PCDATA	$1, $-1
	0x008a 00138 (main.go:15)	PCDATA	$0, $-2
	0x008a 00138 (main.go:15)	CALL	runtime.morestack_noctxt(SB)
	0x008f 00143 (main.go:15)	PCDATA	$0, $-1
	0x008f 00143 (main.go:15)	JMP	0
	0x0000 49 3b 66 10 0f 86 80 00 00 00 48 83 ec 30 48 89  I;f.......H..0H.
	0x0010 6c 24 28 48 8d 6c 24 28 49 c7 c5 00 00 00 00 4c  l$(H.l$(I......L
	0x0020 89 6c 24 20 c6 44 24 07 00 48 c7 44 24 08 00 00  .l$ .D$..H.D$...
	0x0030 00 00 44 0f 11 7c 24 10 48 8d 05 00 00 00 00 48  ..D..|$.H......H
	0x0040 89 44 24 10 48 8d 44 24 08 48 89 44 24 18 48 8d  .D$.H.D$.H.D$.H.
	0x0050 44 24 10 48 89 44 24 20 c6 44 24 07 00 48 8b 54  D$.H.D$ .D$..H.T
	0x0060 24 20 48 8b 02 ff d0 48 8b 44 24 08 48 8b 6c 24  $ H....H.D$.H.l$
	0x0070 28 48 83 c4 30 c3 e8 00 00 00 00 48 8b 44 24 08  (H..0......H.D$.
	0x0080 48 8b 6c 24 28 48 83 c4 30 c3 e8 00 00 00 00 e9  H.l$(H..0.......
	0x0090 6c ff ff ff                                      l...
	rel 59+4 t=14 main.increaseB.func1+0
	rel 101+0 t=10 +0
	rel 119+4 t=7 runtime.deferreturn+0
	rel 139+4 t=7 runtime.morestack_noctxt+0
main.increaseB.func1 STEXT nosplit size=8 args=0x0 locals=0x0 funcid=0x0 align=0x0
	0x0000 00000 (main.go:16)	TEXT	main.increaseB.func1(SB), NOSPLIT|NEEDCTXT|ABIInternal, $0-0
	0x0000 00000 (main.go:16)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (main.go:16)	FUNCDATA	$1, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
	0x0000 00000 (main.go:16)	MOVQ	8(DX), AX
	0x0004 00004 (main.go:17)	INCQ	(AX)
	0x0007 00007 (main.go:18)	RET
	0x0000 48 8b 42 08 48 ff 00 c3                          H.B.H...
main.main STEXT size=174 args=0x0 locals=0x50 funcid=0x0 align=0x0
	0x0000 00000 (main.go:22)	TEXT	main.main(SB), ABIInternal, $80-0
	0x0000 00000 (main.go:22)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:22)	PCDATA	$0, $-2
	0x0004 00004 (main.go:22)	JLS	164
	0x000a 00010 (main.go:22)	PCDATA	$0, $-1
	0x000a 00010 (main.go:22)	SUBQ	$80, SP
	0x000e 00014 (main.go:22)	MOVQ	BP, 72(SP)
	0x0013 00019 (main.go:22)	LEAQ	72(SP), BP
	0x0018 00024 (main.go:22)	FUNCDATA	$0, gclocals·ykHN0vawYuq1dUW4zEe2gA==(SB)
	0x0018 00024 (main.go:22)	FUNCDATA	$1, gclocals·8xBwPYmLuQFhxCB+X15wvg==(SB)
	0x0018 00024 (main.go:22)	FUNCDATA	$2, main.main.stkobj(SB)
	0x0018 00024 (main.go:23)	PCDATA	$1, $0
	0x0018 00024 (main.go:23)	CALL	main.increaseA(SB)
	0x001d 00029 (main.go:23)	MOVUPS	X15, main..autotmp_13+56(SP)
	0x0023 00035 (main.go:23)	PCDATA	$1, $1
	0x0023 00035 (main.go:23)	CALL	runtime.convT64(SB)
	0x0028 00040 (main.go:23)	LEAQ	type.int(SB), CX
	0x002f 00047 (main.go:23)	MOVQ	CX, main..autotmp_13+56(SP)
	0x0034 00052 (main.go:23)	MOVQ	AX, main..autotmp_13+64(SP)
	0x0039 00057 (<unknown line number>)	NOP
	0x0039 00057 ($GOROOT/src/fmt/print.go:294)	MOVQ	os.Stdout(SB), BX
	0x0040 00064 ($GOROOT/src/fmt/print.go:294)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0047 00071 ($GOROOT/src/fmt/print.go:294)	MOVL	$1, DI
	0x004c 00076 ($GOROOT/src/fmt/print.go:294)	MOVQ	DI, SI
	0x004f 00079 ($GOROOT/src/fmt/print.go:294)	LEAQ	main..autotmp_13+56(SP), CX
	0x0054 00084 ($GOROOT/src/fmt/print.go:294)	PCDATA	$1, $0
	0x0054 00084 ($GOROOT/src/fmt/print.go:294)	CALL	fmt.Fprintln(SB)
	0x0059 00089 (main.go:24)	CALL	main.increaseB(SB)
	0x005e 00094 (main.go:24)	MOVUPS	X15, main..autotmp_16+40(SP)
	0x0064 00100 (main.go:24)	PCDATA	$1, $2
	0x0064 00100 (main.go:24)	CALL	runtime.convT64(SB)
	0x0069 00105 (main.go:24)	LEAQ	type.int(SB), CX
	0x0070 00112 (main.go:24)	MOVQ	CX, main..autotmp_16+40(SP)
	0x0075 00117 (main.go:24)	MOVQ	AX, main..autotmp_16+48(SP)
	0x007a 00122 (<unknown line number>)	NOP
	0x007a 00122 ($GOROOT/src/fmt/print.go:294)	MOVQ	os.Stdout(SB), BX
	0x0081 00129 ($GOROOT/src/fmt/print.go:294)	LEAQ	go.itab.*os.File,io.Writer(SB), AX
	0x0088 00136 ($GOROOT/src/fmt/print.go:294)	LEAQ	main..autotmp_16+40(SP), CX
	0x008d 00141 ($GOROOT/src/fmt/print.go:294)	MOVL	$1, DI
	0x0092 00146 ($GOROOT/src/fmt/print.go:294)	MOVQ	DI, SI
	0x0095 00149 ($GOROOT/src/fmt/print.go:294)	PCDATA	$1, $0
	0x0095 00149 ($GOROOT/src/fmt/print.go:294)	CALL	fmt.Fprintln(SB)
	0x009a 00154 (main.go:25)	MOVQ	72(SP), BP
	0x009f 00159 (main.go:25)	ADDQ	$80, SP
	0x00a3 00163 (main.go:25)	RET
	0x00a4 00164 (main.go:25)	NOP
	0x00a4 00164 (main.go:22)	PCDATA	$1, $-1
	0x00a4 00164 (main.go:22)	PCDATA	$0, $-2
	0x00a4 00164 (main.go:22)	CALL	runtime.morestack_noctxt(SB)
	0x00a9 00169 (main.go:22)	PCDATA	$0, $-1
	0x00a9 00169 (main.go:22)	JMP	0
	0x0000 49 3b 66 10 0f 86 9a 00 00 00 48 83 ec 50 48 89  I;f.......H..PH.
	0x0010 6c 24 48 48 8d 6c 24 48 e8 00 00 00 00 44 0f 11  l$HH.l$H.....D..
	0x0020 7c 24 38 e8 00 00 00 00 48 8d 0d 00 00 00 00 48  |$8.....H......H
	0x0030 89 4c 24 38 48 89 44 24 40 48 8b 1d 00 00 00 00  .L$8H.D$@H......
	0x0040 48 8d 05 00 00 00 00 bf 01 00 00 00 48 89 fe 48  H...........H..H
	0x0050 8d 4c 24 38 e8 00 00 00 00 e8 00 00 00 00 44 0f  .L$8..........D.
	0x0060 11 7c 24 28 e8 00 00 00 00 48 8d 0d 00 00 00 00  .|$(.....H......
	0x0070 48 89 4c 24 28 48 89 44 24 30 48 8b 1d 00 00 00  H.L$(H.D$0H.....
	0x0080 00 48 8d 05 00 00 00 00 48 8d 4c 24 28 bf 01 00  .H......H.L$(...
	0x0090 00 00 48 89 fe e8 00 00 00 00 48 8b 6c 24 48 48  ..H.......H.l$HH
	0x00a0 83 c4 50 c3 e8 00 00 00 00 e9 52 ff ff ff        ..P.......R...
	rel 3+0 t=23 type.int+0
	rel 3+0 t=23 type.*os.File+0
	rel 3+0 t=23 type.int+0
	rel 3+0 t=23 type.*os.File+0
	rel 25+4 t=7 main.increaseA+0
	rel 36+4 t=7 runtime.convT64+0
	rel 43+4 t=14 type.int+0
	rel 60+4 t=14 os.Stdout+0
	rel 67+4 t=14 go.itab.*os.File,io.Writer+0
	rel 85+4 t=7 fmt.Fprintln+0
	rel 90+4 t=7 main.increaseB+0
	rel 101+4 t=7 runtime.convT64+0
	rel 108+4 t=14 type.int+0
	rel 125+4 t=14 os.Stdout+0
	rel 132+4 t=14 go.itab.*os.File,io.Writer+0
	rel 150+4 t=7 fmt.Fprintln+0
	rel 165+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.producer.<unlinkable> SDWARFCUINFO dupok size=0
	0x0000 72 65 67 61 62 69                                regabi
go.cuinfo.packagename.main SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 13  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 13 6e 00 01 00 00 00 00 13  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=22 type.[]interface {}+0
	rel 0+0 t=22 type.error+0
	rel 0+0 t=22 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
main..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 5a 22 ee 60 00 00 00 00 00 00 00 00 00 00 00 00  Z".`............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 3b fc f8 8f 08 08 08 36 00 00 00 00 00 00 00 00  ;......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 39 7a 09 0f 02 08 08 14 00 00 00 00 00 00 00 00  9z..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9d 9c 0e 59 08 08 08 36 00 00 00 00 00 00 00 00  ...Y...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 76 de 99 0d 02 08 08 17 00 00 00 00 00 00 00 00  v...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=18
	0x0000 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65 20  ..*[1]interface 
	0x0010 7b 7d                                            {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a8 0e 57 36 08 08 08 36 00 00 00 00 00 00 00 00  ..W6...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 6e 20 6a 3d 02 08 08 11 00 00 00 00 00 00 00 00  n j=............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=-32763 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*func()- SRODATA dupok size=9
	0x0000 00 07 2a 66 75 6e 63 28 29                       ..*func()
type.*func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 75 ac 29 27 08 08 08 36 00 00 00 00 00 00 00 00  u.)'...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 48+8 t=1 type.func()+0
type.func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 fe fa b9 80 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00                                      ....
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 44+4 t=-32763 type.*func()+0
type..namedata.*struct { F uintptr; i *int }- SRODATA dupok size=31
	0x0000 00 1d 2a 73 74 72 75 63 74 20 7b 20 46 20 75 69  ..*struct { F ui
	0x0010 6e 74 70 74 72 3b 20 69 20 2a 69 6e 74 20 7d     ntptr; i *int }
type.*struct { F uintptr; main.i *int } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 d7 b5 2a ea 08 08 08 36 00 00 00 00 00 00 00 00  ..*....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; i *int }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; main.i *int }+0
type..importpath.main. SRODATA dupok size=6
	0x0000 00 04 6d 61 69 6e                                ..main
type..namedata..F- SRODATA dupok size=4
	0x0000 00 02 2e 46                                      ...F
type..namedata.i- SRODATA dupok size=3
	0x0000 00 01 69                                         ..i
type.noalg.struct { F uintptr; main.i *int } SRODATA dupok size=128
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 92 fd 49 a5 02 08 08 19 00 00 00 00 00 00 00 00  ..I.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; i *int }-+0
	rel 44+4 t=-32763 type.*struct { F uintptr; main.i *int }+0
	rel 48+8 t=1 type..importpath.main.+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; main.i *int }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.i-+0
	rel 112+8 t=1 type.*int+0
type..namedata.*struct { F uintptr; r *int }- SRODATA dupok size=31
	0x0000 00 1d 2a 73 74 72 75 63 74 20 7b 20 46 20 75 69  ..*struct { F ui
	0x0010 6e 74 70 74 72 3b 20 72 20 2a 69 6e 74 20 7d     ntptr; r *int }
type.*struct { F uintptr; main.r *int } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 33 33 ed d9 08 08 08 36 00 00 00 00 00 00 00 00  33.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; r *int }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; main.r *int }+0
type..namedata.r- SRODATA dupok size=3
	0x0000 00 01 72                                         ..r
type.noalg.struct { F uintptr; main.r *int } SRODATA dupok size=128
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 79 09 f8 01 02 08 08 19 00 00 00 00 00 00 00 00  y...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; r *int }-+0
	rel 44+4 t=-32763 type.*struct { F uintptr; main.r *int }+0
	rel 48+8 t=1 type..importpath.main.+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; main.r *int }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.r-+0
	rel 112+8 t=1 type.*int+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·J5F+7Qw7O7ve2QcWC7DpeQ== SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·wdmTuppZUxZYftR7OCq88Q== SRODATA dupok size=10
	0x0000 02 00 00 00 03 00 00 00 00 04                    ..........
main.increaseA.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 e8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
main.increaseA.opendefer SRODATA dupok size=3
	0x0000 29 01 08                                         )..
gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
main.increaseB.stkobj SRODATA static size=24
	0x0000 01 00 00 00 00 00 00 00 e8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
main.increaseB.opendefer SRODATA dupok size=3
	0x0000 21 01 08                                         !..
gclocals·ykHN0vawYuq1dUW4zEe2gA== SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·8xBwPYmLuQFhxCB+X15wvg== SRODATA dupok size=11
	0x0000 03 00 00 00 04 00 00 00 00 08 02                 ...........
main.main.stkobj SRODATA static size=40
	0x0000 02 00 00 00 00 00 00 00 e0 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
	0x0020 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
	rel 36+4 t=5 runtime.gcbits.02+0
