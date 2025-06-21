# acrun

## build

### docker image
```sh
$ make docker-image
```

### docker run
```sh
$ docker run --rm --mount type=bind,source="$(pwd)",target=/work/src acrun:latest ./acrun [-c {コンテスト名(eg. abcXXX}] [-e {実行ファイル名}] [-f {ソースコード}] [-t {タスク名(eg. abcXXX)}] {言語} {問題}
// タスク名は問題URLパス末尾のabcXXX
```

## command

### usage
```
acrun [...options] lang question

options:
  -c string
        contest name (default: current dir name)
  -e string
        exec comand name
  -f string
        file name
  -t string
        contest question task (default: value of question)
```
