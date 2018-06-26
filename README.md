# yamlsh - YAML Parsed in `#!` for shell

(I'm sure there is already a util out there that does this (I couldn't find one,
though admittedly I didn't search very hard), but I'm learning the
Go language and needed the practice, so...)

Parse a yaml file and pass as environment variables to a script from `#!`.

You can download a prebuilt binary from the project's releases page
[here](https://github.com/gbevan/yamlsh/releases/latest).

## Build and install
Requires GOPATH setup and this project cloned out to `src/`
```
go get github.com/golang/dep/cmd/dep
dep ensure -v --vendor-only
go build && sudo cp yamlsh /usr/local/bin
```

## Usage
To use within a shell script replace the #! line with:
```
#!/usr/local/bin/yamlsh --yaml=${MYFILE}
# ...rest of your script goes here...
```
`${MYFILE}` will be substituted with the environment variable `MYFILE`, e.g.:
```
MYFILE=test.yml ./test.sh
```
The `test.sh` script and `test.yml` file produces this output:
```
IN SCRIPT
YAMLSH_DICT1_NEST1_FloatNum=3.14
YAMLSH_VAR1=my var 1
YAMLSH_DICT1_NEST1_MultiLinePp=aaaaaaaaaaaaaaaaaaaaaa\nbbbbbbbbbbbbbbbbbbbbbb\ncccccccccccccccccccccc\n
YAMLSH_DICT1_NEST1_MultiLineGt=aaaaaaaaaaaaaaaaaaaaaa bbbbbbbbbbbbbbbbbbbbbb cccccccccccccccccccccc\n
YAMLSH_DICT1_NEST1_ARRAY2_1=array
YAMLSH_DICT1_NEST1_ARRAY2_0=an
YAMLSH_DICT1_NEST1_Number=100
YAMLSH_PREFIX=YAMLSH
YAMLSH_DICT1_NEST1_NEST2=value of NEST2
YAMLSH_ARRAY1_0=val1
YAMLSH_ARRAY1_1=val2
```
As can be seen, complex yaml structures are flattened for easy use in shell
scripts.

By default variables passed are prefixed with `YAMLSH_`.  To change this use
the `YAMLSH_PREFIX` environment variable, e.g.:
```
MYFILE=test.yml YAMLSH_PREFIX=FOO ./test.sh
```
The shell used defaults to `/bin/bash`, but you can override to use your
preferred shell via environment variable `YAMLSH_SHELL`

Likewise you can also set the YAML file to preload using environment
variable `YAMLSH_YAMLFILE`.  However, the `--yaml=file.yml` will take
precedence.
