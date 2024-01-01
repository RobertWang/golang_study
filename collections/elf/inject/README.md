### 编译动态链接库 so 文件

要在命令行编译一个 Golang 源代码为动态链接库，请按照以下步骤操作：

首先确保已经安装了 Go 语言环境。如果没有安装，可以访问 Go 官网 下载并安装。

打开命令行，进入到源代码所在的目录。

使用 go build 命令编译源代码。例如，如果源代码文件名为 main.go，则输入以下命令：

```bash
go build -o libfunc.so main.go
```

这将编译 main.go 文件并生成一个名为 libexample.so 的动态链接库。如果需要指定其他选项，可以在 go build 命令后添加相应的参数。

### 注入

首先，需要使用 go tool 工具中的 objcopy 命令将已经存在的 so 文件中的 helloworld 方法导出为一个符号。假设该 so 文件名为 libhello.so，导出的符号名为 `_helloworld`。

```bash
$ objcopy -t elf64-x86-64 --symbol=_helloworld libhello.so libhello_hook.so
```

```
objcopy -t elf64-x86-64 --symbol=_hellox funcs.so funcs_hook.so
objcopy: invalid option -- t
Usage: objcopy [option(s)] in-file [out-file]
 Copies a binary file, possibly transforming it in the process
 The options are:
  -I --input-target <bfdname>      Assume input file is in format <bfdname>
  -O --output-target <bfdname>     Create an output file in format <bfdname>
  -B --binary-architecture <arch>  Set output arch, when input is arch-less
  -F --target <bfdname>            Set both input and output format to <bfdname>
     --debugging                   Convert debugging information, if possible
  -p --preserve-dates              Copy modified/access timestamps to the output
  -D --enable-deterministic-archives
                                   Produce deterministic output when stripping archives (default)
  -U --disable-deterministic-archives
                                   Disable -D behavior
  -j --only-section <name>         Only copy section <name> into the output
     --add-gnu-debuglink=<file>    Add section .gnu_debuglink linking to <file>
  -R --remove-section <name>       Remove section <name> from the output
     --remove-relocations <name>   Remove relocations from section <name>
     --strip-section-headers              Strip section header from the output
  -S --strip-all                   Remove all symbol and relocation information
  -g --strip-debug                 Remove all debugging symbols & sections
     --strip-dwo                   Remove all DWO sections
     --strip-unneeded              Remove all symbols not needed by relocations
  -N --strip-symbol <name>         Do not copy symbol <name>
     --strip-unneeded-symbol <name>
                                   Do not copy symbol <name> unless needed by
                                     relocations
     --only-keep-debug             Strip everything but the debug information
     --extract-dwo                 Copy only DWO sections
     --extract-symbol              Remove section contents but keep symbols
     --keep-section <name>         Do not strip section <name>
  -K --keep-symbol <name>          Do not strip symbol <name>
     --keep-section-symbols        Do not strip section symbols
     --keep-file-symbols           Do not strip file symbol(s)
     --localize-hidden             Turn all ELF hidden symbols into locals
  -L --localize-symbol <name>      Force symbol <name> to be marked as a local
     --globalize-symbol <name>     Force symbol <name> to be marked as a global
  -G --keep-global-symbol <name>   Localize all symbols except <name>
  -W --weaken-symbol <name>        Force symbol <name> to be marked as a weak
     --weaken                      Force all global symbols to be marked as weak
  -w --wildcard                    Permit wildcard in symbol comparison
  -x --discard-all                 Remove all non-global symbols
  -X --discard-locals              Remove any compiler-generated symbols
  -i --interleave[=<number>]       Only copy N out of every <number> bytes
     --interleave-width <number>   Set N for --interleave
  -b --byte <num>                  Select byte <num> in every interleaved block
     --gap-fill <val>              Fill gaps between sections with <val>
     --pad-to <addr>               Pad the last section up to address <addr>
     --set-start <addr>            Set the start address to <addr>
    {--change-start|--adjust-start} <incr>
                                   Add <incr> to the start address
    {--change-addresses|--adjust-vma} <incr>
                                   Add <incr> to LMA, VMA and start addresses
    {--change-section-address|--adjust-section-vma} <name>{=|+|-}<val>
                                   Change LMA and VMA of section <name> by <val>
     --change-section-lma <name>{=|+|-}<val>
                                   Change the LMA of section <name> by <val>
     --change-section-vma <name>{=|+|-}<val>
                                   Change the VMA of section <name> by <val>
    {--[no-]change-warnings|--[no-]adjust-warnings}
                                   Warn if a named section does not exist
     --set-section-flags <name>=<flags>
                                   Set section <name>'s properties to <flags>
     --set-section-alignment <name>=<align>
                                   Set section <name>'s alignment to <align> bytes
     --add-section <name>=<file>   Add section <name> found in <file> to output
     --update-section <name>=<file>
                                   Update contents of section <name> with
                                   contents found in <file>
     --dump-section <name>=<file>  Dump the contents of section <name> into <file>
     --rename-section <old>=<new>[,<flags>] Rename section <old> to <new>
     --long-section-names {enable|disable|keep}
                                   Handle long section names in Coff objects.
     --change-leading-char         Force output format's leading character style
     --remove-leading-char         Remove leading character from global symbols
     --reverse-bytes=<num>         Reverse <num> bytes at a time, in output sections with content
     --redefine-sym <old>=<new>    Redefine symbol name <old> to <new>
     --redefine-syms <file>        --redefine-sym for all symbol pairs
                                     listed in <file>
     --srec-len <number>           Restrict the length of generated Srecords
     --srec-forceS3                Restrict the type of generated Srecords to S3
     --strip-symbols <file>        -N for all symbols listed in <file>
     --strip-unneeded-symbols <file>
                                   --strip-unneeded-symbol for all symbols listed
                                     in <file>
     --keep-symbols <file>         -K for all symbols listed in <file>
     --localize-symbols <file>     -L for all symbols listed in <file>
     --globalize-symbols <file>    --globalize-symbol for all in <file>
     --keep-global-symbols <file>  -G for all symbols listed in <file>
     --weaken-symbols <file>       -W for all symbols listed in <file>
     --add-symbol <name>=[<section>:]<value>[,<flags>]  Add a symbol
     --alt-machine-code <index>    Use the target's <index>'th alternative machine
     --writable-text               Mark the output text as writable
     --readonly-text               Make the output text write protected
     --pure                        Mark the output file as demand paged
     --impure                      Mark the output file as impure
     --prefix-symbols <prefix>     Add <prefix> to start of every symbol name
     --prefix-sections <prefix>    Add <prefix> to start of every section name
     --prefix-alloc-sections <prefix>
                                   Add <prefix> to start of every allocatable
                                     section name
     --file-alignment <num>        Set PE file alignment to <num>
     --heap <reserve>[,<commit>]   Set PE reserve/commit heap to <reserve>/
                                   <commit>
     --image-base <address>        Set PE image base to <address>
     --section-alignment <num>     Set PE section alignment to <num>
     --stack <reserve>[,<commit>]  Set PE reserve/commit stack to <reserve>/
                                   <commit>
     --subsystem <name>[:<version>]
                                   Set PE subsystem to <name> [& <version>]
     --compress-debug-sections[={none|zlib|zlib-gnu|zlib-gabi|zstd}]
                                   Compress DWARF debug sections
     --decompress-debug-sections   Decompress DWARF debug sections using zlib
     --elf-stt-common=[yes|no]     Generate ELF common symbols with STT_COMMON
                                     type
     --verilog-data-width <number> Specifies data width, in bytes, for verilog output
  -M  --merge-notes                Remove redundant entries in note sections
      --no-merge-notes             Do not attempt to remove redundant notes (default)
  -v --verbose                     List all object files modified
  @<file>                          Read options from <file>
  -V --version                     Display this program's version number
  -h --help                        Display this output
     --info                        List object formats & architectures supported
objcopy: supported targets: mach-o-x86-64 elf32-bigaarch64 elf32-littleaarch64 elf64-bigaarch64 elf64-bigaarch64-cloudabi elf64-littleaarch64 elf64-littleaarch64-cloudabi mach-o-arm64 pe-aarch64-little pei-aarch64-little elf64-amdgcn ecoff-littlealpha elf64-alpha elf64-alpha-freebsd vms-alpha vms-libtxt elf32-am33lin elf32-bigarc elf32-littlearc elf32-bigarm elf32-littlearm elf32-bigarm-fdpic elf32-littlearm-fdpic elf32-bigarm-vxworks elf32-littlearm-vxworks mach-o-arm pe-arm-big pe-arm-little pe-arm-wince-big pe-arm-wince-little pei-arm-big pei-arm-little pei-arm-wince-big pei-arm-wince-little elf32-avr elf32-bfin elf32-bfinfdpic elf32-cr16 a.out-cris elf32-cris elf32-us-cris elf32-crx elf32-csky-big elf32-csky-little elf32-d10v elf32-d30v elf32-dlx elf32-big elf32-little elf64-big elf64-little elf32-epiphany elf32-fr30 elf32-frv elf32-frvfdpic elf32-h8300 elf32-h8300-linux elf32-hppa elf32-hppa-linux elf32-hppa-netbsd elf64-hppa elf64-hppa-linux som a.out-i386 a.out-i386-bsd a.out-i386-lynx coff-i386 coff-go32 coff-go32-exe coff-i386-lynx elf32-i386 elf32-i386-freebsd elf32-i386-sol2 elf32-i386-vxworks mach-o-i386 msdos pe-i386 pe-bigobj-i386 pei-i386 elf32-iamcu elf64-bpfbe elf64-bpfle elf32-ia64-hpux-big elf64-ia64-big elf64-ia64-little elf64-ia64-hpux-big elf64-ia64-vms pei-ia64 elf32-ip2k elf32-iq2000 elf32-lm32 elf32-m32c elf32-m32r elf32-m32rle elf32-m32r-linux elf32-m32rle-linux elf32-m68hc11 elf32-m68hc12 elf32-m68k elf32-s12z mach-o-be mach-o-le mach-o-fat elf32-mcore-big elf32-mcore-little pe-mcore-big pe-mcore-little pei-mcore-big pei-mcore-little elf32-mep elf32-metag elf32-microblaze ecoff-bigmips ecoff-littlemips ecoff-biglittlemips elf32-bigmips elf32-littlemips elf32-nbigmips elf32-nlittlemips elf32-ntradbigmips elf32-ntradlittlemips elf32-ntradbigmips-freebsd elf32-ntradlittlemips-freebsd elf32-tradbigmips elf32-tradlittlemips elf32-tradbigmips-freebsd elf32-tradlittlemips-freebsd elf32-bigmips-vxworks elf32-littlemips-vxworks elf64-bigmips elf64-littlemips elf64-tradbigmips elf64-tradlittlemips elf64-tradbigmips-freebsd elf64-tradlittlemips-freebsd elf64-mmix mmo elf32-mn10200 elf32-mn10300 elf32-bigmoxie elf32-littlemoxie elf32-msp430 elf32-msp430 elf32-mt elf32-nds32be elf32-nds32le elf32-nds32be-linux elf32-nds32le-linux elf64-nfp elf32-bignios2 elf32-littlenios2 a.out-pc532-mach a.out-ns32k-netbsd elf32-or1k pdb a.out-pdp11 pef pef-xlib elf32-pj elf32-pjl ppcboot elf32-powerpc elf32-powerpcle elf32-powerpc-freebsd elf32-powerpc-vxworks elf64-powerpc elf64-powerpcle elf64-powerpc-freebsd elf64-powerpcle-freebsd elf32-pru elf32-littleriscv elf64-littleriscv elf32-bigriscv elf64-bigriscv elf32-rl78 aixcoff64-rs6000 aix5coff64-rs6000 aixcoff-rs6000 elf32-rx-be elf32-rx-be-ns elf32-rx-le elf32-s390 elf64-s390 elf32-bigscore elf32-littlescore coff-sh coff-shl coff-sh-small coff-shl-small elf32-sh elf32-shl elf32-shbig-fdpic elf32-sh-fdpic elf32-sh-linux elf32-shbig-linux elf32-sh-nbsd elf32-shl-nbsd elf32-sh-vxworks elf32-shl-vxworks pe-shl pei-shl elf32-sparc elf32-sparc-sol2 elf32-sparc-vxworks elf64-sparc elf64-sparc-freebsd elf64-sparc-sol2 elf32-spu sym coff-tic30 coff0-beh-c54x coff0-c54x coff1-beh-c54x coff1-c54x coff2-beh-c54x coff2-c54x elf32-tic6x-be elf32-tic6x-le elf32-tilegx-be elf32-tilegx-le elf64-tilegx-be elf64-tilegx-le elf32-tilepro elf32-ft32 elf32-v850-rh850 elf32-v850 a.out-vax1k-netbsd a.out-vax-netbsd elf32-vax elf32-visium wasm elf32-wasm32 coff-x86-64 elf32-x86-64 elf64-x86-64 elf64-x86-64-cloudabi elf64-x86-64-freebsd elf64-x86-64-sol2 pe-x86-64 pe-bigobj-x86-64 pei-x86-64 elf32-xgate elf32-xstormy16 elf32-xtensa-be elf32-xtensa-le coff-z80 elf32-z80 coff-z8k elf32-loongarch elf64-loongarch pei-loongarch64 srec symbolsrec verilog tekhex binary ihex plugin
```

接下来，可以使用以下代码向 ELF 二进制文件程序入口处注入一个 hook：

```golang
package main

import (
    "fmt"
    "syscall"
)

func main() {
    // 打开 ELF 二进制文件
    file, err := os.Open("/path/to/binary")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // 获取 ELF 二进制文件的程序入口地址
    entry, err := getEntry(file)
    if err != nil {
        fmt.Println(err)
        return
    }

    // 在程序入口处写入跳转指令，跳转到 hook 函数中
    hookAddress, err := getHookAddress(file)
    if err != nil {
        fmt.Println(err)
        return
    }
    if _, err := syscall.Write(int(entry), []byte("\xc3")); err != nil {
        fmt.Println(err)
        return
    }
    if _, err := syscall.Write(int(entry), []byte("\x8b")); err != nil {
        fmt.Println(err)
        return
    }
    if _, err := syscall.Write(int(entry), []byte("movl $0x" + string(hookAddress) + "\n")); err != nil {
        fmt.Println(err)
        return
    }
    if _, err := syscall.Write(int(entry), []byte("ret")); err != nil {
        fmt.Println(err)
        return
    }
}

// 获取 ELF 二进制文件的程序入口地址
func getEntry(file *os.File) (uintptr, error) {
    var entry uintptr
    if err := binary.Read(file, binary.LittleEndian, &entry); err != nil {
        return 0, err
    }
    return entry, nil
}

// 获取 hook 函数的地址
func getHookAddress(file *os.File) (uintptr, error) {
    var hookAddr uintptr
    if err := binary.Read(file, binary.LittleEndian, &hookAddr); err != nil {
        return 0, err
    }
    return hookAddr, nil
}
```

以上代码会将 ELF 二进制文件的程序入口地址修改为跳转到 hook 函数的地址，并在 hook 函数中调用已经存在的 so 文件中的 helloworld 方法。

### OSX 下读二进制程序文件

```bash
brew install binutils
```

首先，mac 系统下的文件格式是 mach-o，并不支持 objdump 和 readelf ；
退而求其次，使用 brew update && brew install binutils，然后用 greadelf 和 gobjdump；

注意，使用使用 brew update && brew install binutils 会有如下提示：

```bash
binutils is keg-only, which means it was not symlinked into /usr/local,
because because Apple provides the same tools and binutils is poorly supported on macOS.

If you need to have binutils first in your PATH run:
  echo 'export PATH="/usr/local/opt/binutils/bin:$PATH"' >> ~/.bash_profile

For compilers to find binutils you may need to set:
  export LDFLAGS="-L/usr/local/opt/binutils/lib"
  export CPPFLAGS="-I/usr/local/opt/binutils/include"
```

### go help build

```
usage: go build [-o output] [build flags] [packages]

Build compiles the packages named by the import paths,
along with their dependencies, but it does not install the results.

If the arguments to build are a list of .go files from a single directory,
build treats them as a list of source files specifying a single package.

When compiling packages, build ignores files that end in '\_test.go'.

When compiling a single main package, build writes
the resulting executable to an output file named after
the first source file ('go build ed.go rx.go' writes 'ed' or 'ed.exe')
or the source code directory ('go build unix/sam' writes 'sam' or 'sam.exe').
The '.exe' suffix is added when writing a Windows executable.

When compiling multiple packages or a single non-main package,
build compiles the packages but discards the resulting object,
serving only as a check that the packages can be built.

The -o flag forces build to write the resulting executable or object
to the named output file or directory, instead of the default behavior described
in the last two paragraphs. If the named output is an existing directory or
ends with a slash or backslash, then any resulting executables
will be written to that directory.

The -i flag installs the packages that are dependencies of the target.
The -i flag is deprecated. Compiled packages are cached automatically.

The build flags are shared by the build, clean, get, install, list, run,
and test commands:

        -a
                force rebuilding of packages that are already up-to-date.
        -n
                print the commands but do not run them.
        -p n
                the number of programs, such as build commands or
                test binaries, that can be run in parallel.
                The default is GOMAXPROCS, normally the number of CPUs available.
        -race
                enable data race detection.
                Supported only on linux/amd64, freebsd/amd64, darwin/amd64, darwin/arm64, windows/amd64,
                linux/ppc64le and linux/arm64 (only for 48-bit VMA).
        -msan
                enable interoperation with memory sanitizer.
                Supported only on linux/amd64, linux/arm64
                and only with Clang/LLVM as the host C compiler.
                On linux/arm64, pie build mode will be used.
        -asan
                enable interoperation with address sanitizer.
                Supported only on linux/arm64, linux/amd64.
                Supported only on linux/amd64 or linux/arm64 and only with GCC 7 and higher
                or Clang/LLVM 9 and higher.
        -v
                print the names of packages as they are compiled.
        -work
                print the name of the temporary work directory and
                do not delete it when exiting.
        -x
                print the commands.

        -asmflags '[pattern=]arg list'
                arguments to pass on each go tool asm invocation.
        -buildmode mode
                build mode to use. See 'go help buildmode' for more.
        -buildvcs
                Whether to stamp binaries with version control information
                ("true", "false", or "auto"). By default ("auto"), version control
                information is stamped into a binary if the main package, the main module
                containing it, and the current directory are all in the same repository.
                Use -buildvcs=false to always omit version control information, or
                -buildvcs=true to error out if version control information is available but
                cannot be included due to a missing tool or ambiguous directory structure.
        -compiler name
                name of compiler to use, as in runtime.Compiler (gccgo or gc).
        -gccgoflags '[pattern=]arg list'
                arguments to pass on each gccgo compiler/linker invocation.
        -gcflags '[pattern=]arg list'
                arguments to pass on each go tool compile invocation.
        -installsuffix suffix
                a suffix to use in the name of the package installation directory,
                in order to keep output separate from default builds.
                If using the -race flag, the install suffix is automatically set to race
                or, if set explicitly, has _race appended to it. Likewise for the -msan
                and -asan flags. Using a -buildmode option that requires non-default compile
                flags has a similar effect.
        -ldflags '[pattern=]arg list'
                arguments to pass on each go tool link invocation.
        -linkshared
                build code that will be linked against shared libraries previously
                created with -buildmode=shared.
        -mod mode
                module download mode to use: readonly, vendor, or mod.
                By default, if a vendor directory is present and the go version in go.mod
                is 1.14 or higher, the go command acts as if -mod=vendor were set.
                Otherwise, the go command acts as if -mod=readonly were set.
                See https://golang.org/ref/mod#build-commands for details.
        -modcacherw
                leave newly-created directories in the module cache read-write
                instead of making them read-only.
        -modfile file
                in module aware mode, read (and possibly write) an alternate go.mod
                file instead of the one in the module root directory. A file named
                "go.mod" must still be present in order to determine the module root
                directory, but it is not accessed. When -modfile is specified, an
                alternate go.sum file is also used: its path is derived from the
                -modfile flag by trimming the ".mod" extension and appending ".sum".
        -overlay file
                read a JSON config file that provides an overlay for build operations.
                The file is a JSON struct with a single field, named 'Replace', that
                maps each disk file path (a string) to its backing file path, so that
                a build will run as if the disk file path exists with the contents
                given by the backing file paths, or as if the disk file path does not
                exist if its backing file path is empty. Support for the -overlay flag
                has some limitations: importantly, cgo files included from outside the
                include path must be in the same directory as the Go package they are
                included from, and overlays will not appear when binaries and tests are
                run through go run and go test respectively.
        -pkgdir dir
                install and load all packages from dir instead of the usual locations.
                For example, when building with a non-standard configuration,
                use -pkgdir to keep generated packages in a separate location.
        -tags tag,list
                a comma-separated list of additional build tags to consider satisfied
                during the build. For more information about build tags, see
                'go help buildconstraint'. (Earlier versions of Go used a
                space-separated list, and that form is deprecated but still recognized.)
        -trimpath
                remove all file system paths from the resulting executable.
                Instead of absolute file system paths, the recorded file names
                will begin either a module path@version (when using modules),
                or a plain import path (when using the standard library, or GOPATH).
        -toolexec 'cmd args'
                a program to use to invoke toolchain programs like vet and asm.
                For example, instead of running asm, the go command will run
                'cmd args /path/to/asm <arguments for asm>'.
                The TOOLEXEC_IMPORTPATH environment variable will be set,
                matching 'go list -f {{.ImportPath}}' for the package being built.

The -asmflags, -gccgoflags, -gcflags, and -ldflags flags accept a
space-separated list of arguments to pass to an underlying tool
during the build. To embed spaces in an element in the list, surround
it with either single or double quotes. The argument list may be
preceded by a package pattern and an equal sign, which restricts
the use of that argument list to the building of packages matching
that pattern (see 'go help packages' for a description of package
patterns). Without a pattern, the argument list applies only to the
packages named on the command line. The flags may be repeated
with different patterns in order to specify different arguments for
different sets of packages. If a package matches patterns given in
multiple flags, the latest match on the command line wins.
For example, 'go build -gcflags=-S fmt' prints the disassembly
only for package fmt, while 'go build -gcflags=all=-S fmt'
prints the disassembly for fmt and all its dependencies.

For more about specifying packages, see 'go help packages'.
For more about where packages and binaries are installed,
run 'go help gopath'.
For more about calling between Go and C/C++, run 'go help c'.

Note: Build adheres to certain conventions such as those described
by 'go help gopath'. Not all projects can follow these conventions,
however. Installations that have their own conventions or that use
a separate software build system may choose to use lower-level
invocations such as 'go tool compile' and 'go tool link' to avoid
some of the overheads and design decisions of the build tool.

See also: go install, go get, go clean.
```

### 参考资料

- [GoLang 调用 .so 文件](https://www.kandaoni.com/news/25186.html)
- [golang的动态库(so)生成与使用](https://blog.csdn.net/arv002/article/details/122955707)
- [Mac上使用objdump和readelf](https://www.jianshu.com/p/4b93b0665a2b)
- [Go程序性能分析工具和方法](https://segmentfault.com/a/1190000040053014)


