# wc-tool
Custom version of the Unix command line tool `wc`.

## Build
Generate the `wc-tool` executable from the root directory for use:
```shell
go build -o wc-tool cmd/wc-tool/main.go
```

## Usage

### `-c` option
Outputs the number of bytes in a file, e.g.:

```shell
$ ./wc-tool -c test/test.txt
  342190 test/test.txt
```

### `-l` option
Outputs the number of lines in a file.

```shell
$ ./wc-tool -l test/test.txt
    7145 test/test.txt
```

### `-w` option
Outputs the number of words in a file.

```shell
$ ./wc-tool -w test/test.txt
   58164 test/test.txt
```

### `-m` option
Outputs the number of characters in a file. If the current locale does not support multibyte characters this will match the -c option.

```shell
$ ./wc-tool -m test/test.txt
  339292 test/test.txt
```

### Default option
No options provided, equivalent to `-c`, `-l`, and `-w` options.

```shell
$ ./wc-tool test/test.txt
    7145   58164  342190 test/test.txt
```

### Read from standard input if no file specified

```shell
$ cat test/test.txt | ./wc-tool -l
    7145
```