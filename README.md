<h1 align="center">
  Pliki Server
</h1>

<h4 align="center">A temporary file storage server.</h4>

<p align="center"> <a href="https://github.com/pedrobarco/pliki-server/issues">
    <img src="https://img.shields.io/github/issues/pedrobarco/pliki-server"
         alt="issues">
  </a>
 <a href="https://github.com/pedrobarco/pliki-server/network/members">
    <img src="https://img.shields.io/github/forks/pedrobarco/pliki-server"
         alt="forks">
  </a>
 <a href="https://github.com/pedrobarco/pliki-server/stargazers">
    <img src="https://img.shields.io/github/stars/pedrobarco/pliki-server"
         alt="stargazers">
  </a>
 <a href="https://github.com/pedrobarco/pliki-server/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/pedrobarco/pliki-server"
         alt="LICENSE">
  </a>
</p>
<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#intro">Intro</a> •
  <a href="#install">Install</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#credits">Credits</a> •
  <a href="#contribute">Contribute</a> •
  <a href="#license">License</a>
</p>

## Key Features

- [x] upload/download files
- [ ] control data lifecycles
- [ ] apply file retention policies

## Intro

Pliki-server is a temporary file storage server.

The server is inspired by [catbox](https://catbox.moe/)'s "temporary storage" aspect, but no code comes from it. 

## Install

You can [download](https://github.com/pedrobarco/pliki-server/releases) the latest installable version of Pliki Server for Windows, macOS and Linux.

To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://golang.org/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/pedrobarco/pliki-server.git

# Go into the repository
$ cd pliki-server

# Build the app
$ make clean && make build

# Install it
$ sudo make install

# Run the CLI
$ pliki -h
```

## How to Use

TBD

## Configuration

TBD

## Credits

This software uses the following open source packages:

- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)
- [goreleaser](https://github.com/goreleaser/goreleaser)
- [nfpm](https://github.com/goreleaser/nfpm)
- [golanci-lint](https://github.com/golangci/golangci-lint)
- [pre-commit](https://github.com/pre-commit/pre-commit)

## Contribute

1. [Fork it](https://github.com/pedrobarco/pliki-server/fork)
2. Create your feature branch (`git checkout -b feature/myFeature`)
3. Commit your changes (`git commit -am 'feat: add something'`)
4. Push to the branch (`git push origin feature/myFeature`)
5. Create a new [Pull Request](https://github.com/pedrobarco/pliki-server/pulls)

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/pedrobarco/pliki-server/blob/main/LICENSE) file for details

---

> GitHub [@pedrobarco](https://github.com/pedrobarco)

