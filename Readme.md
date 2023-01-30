# Go Emoji Calculator

This code implements a calculator able to parse expressions containing  Emoji unicode characters like :100: :keycap_ten: :plus: or :8ball:


the `/parser` package was generated using [ANTLR4](https://github.com/antlr/antlr4) 

# requisites

* Go 1.19

* Install ANTLR4 runtime go package

```
go get github.com/antlr/antlr4/runtime/Go/antlr/v4
```

# building

On Unix/Linux/Macos:

```
bash build.sh
```

On Windows:

```
build.bat
```

# running

On Unix/Linux/Macos:

```
echo -n "4️2️+🎱+25✖️2" | ./emoji-calculator
```

On Windows:

```
echo "4️2️+🎱+25✖️2" | ./emoji-calculator
```

# running tests

On Unix/Linux/Macos:

```
bash run_tests.sh
```

On Windows:

```
run_tests.bat
```



