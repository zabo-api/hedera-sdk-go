# Hedera Go SDK
> Hedera SDK for Go

## Requirements

 * Go 1.11+

 * [Git LFS](https://git-lfs.github.com) – Git LFS is used to version the built libraries in [`./libs`](./libs) (built from `hedera-sdk-rust`). This must be installed on your system prior to `go get`.

    - MacOS (homebrew)

        ```sh
        $ brew install git-lfs
        $ git lfs install
        ```

    - Ubuntu / Debian

        ```sh
        $ sudo apt install git-lfs
        $ git lfs install
        ```

    - Windows

        Download and run [git-lfs.exe](https://github.com/git-lfs/git-lfs/releases/download/v2.6.0/git-lfs-windows-v2.6.0.exe). Open your git console ( under `Git Bash` if you installed Git through https://git-scm.com/ ).

        ```sh
        $ git lfs install
        ```

## Install

```sh
$ go get github.com/hashgraph/hedera-sdk-go
```

## Development

### Requirements

 * [rustup](https://rustup.rs/)

 * musl

 * MinGW

#### MacOS

```sh
$ brew install FiloSottile/musl-cross/musl-cross
$ brew install mingw-w64
```

### Build

```sh
$ ./x.py
```

## License

Licensed under Apache License,
Version 2.0 ([LICENSE](LICENSE) or http://www.apache.org/licenses/LICENSE-2.0).

## Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in the work by you, as defined in the Apache-2.0 license, shall be
licensed as above, without any additional terms or conditions.
