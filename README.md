# gonel

## How to install

> Warning: you need `goimports` installed.

```bash
$ git clone https://github.com/xrash/gonel.git
$ cd gonel
$ make
$ sudo cp bin/gonel /usr/local/bin
```

## Usage

```bash
Usage:
  gonel [flags]

  Flags:
    -h, --help                 help for gonel
    -i, --import stringArray   package name to import
```

## Examples

```bash
$ gonel "fmt.Println(123)"
123

$ gonel "fmt.Println(333)" "fmt.Println(999)"
333
999
	
$ gonel 'fmt.Println(123)
> fmt.Println(333)
> fmt.Println(999)'
123
333
999
```

## More information

1. You must have goimports installed in your machine.

2. `gonel` will automatically import your packages using `goimports`, so you can write code as if you were inside `main()` with everything set up already.

3. If you still need to explicitly import a package, use the flag `-i`, like this: `$ gonel -i fmt -i math 'fmt.Println(math.Max(20, 30))'`.

4. You can use multiple lines/pieces of code as multiple arguments, or just break lines in the same argument.

5. STDIN passed to `gonel` will be forwarded to the program, so you'll be able to access it through `os.Stdin`.
