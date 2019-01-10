## Overview

tantan-simplify provides some simple match functions between two people.

## How to run

### compile

execute command below：

```shell
$ ./scripts/generate-version.sh v0.0.1 ; ./build.sh
```

after compiled, go to `_build`，run：

```shell
$ cd _build
$ ./bin/tantan-simplify -c conf/{dev/pre/release}/config.toml
```
execute follow command to build docker image：

```shell
$ ./scripts/generate-version.sh v0.0.1 ; ./build.sh
$ cd _build
$ docker build -t jayantxie/tantan-simplify:v0.0.1 .
```

docker run command：

```shell
$ docker run -it -p 8080:8080 -e TANTAN_SIMPLIFY_STARTUP_ARGS="-c conf/{dev/pre/release}/config.toml" ${Image ID}
```

### test

run all test

```
$ ./run-all-test.sh
```
### dependency upgrade

1. modify ` Gopkg.toml`；

2. execute：

    ```shell
    $ dep ensure --update
    ```
   
3. make sure compiled successfully：

    ```shell
    $ ./build.sh
    ```
  
4. commit new `vendor/` 和 `Gopkg*`；