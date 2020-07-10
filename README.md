# godump

godump dumps AST and SSA IR of given package.

## Usage

### Dump AST

```sh
$ godump /tmp/main.go
/tmp/main.go
File
├── Doc
├── Package = /tmp/main.go:1:1
├── Name
│   └── Ident
│       ├── NamePos = /tmp/main.go:1:9
│       ├── Name = main
│       └── Obj
...
```

### Dump SSA IR

```sh
$ godump -mode=ssa /tmp/main.go
command-line-arguments.main
        Block 0
                *ssa.Call       println("hello, world":string)
                *ssa.Return     return
```
