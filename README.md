[![Build Status](https://travis-ci.com/chz8494/gored.svg?branch=master)](https://travis-ci.com/chz8494/gored)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/chz8494/gored/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/chz8494/gored?status.svg)](https://godoc.org/github.com/chz8494/gored)
[![Coverage Status](http://codecov.io/github/chz8494/gored/coverage.svg?branch=master)](http://codecov.io/github/chz8494/gored?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/chz8494/gored)](https://goreportcard.com/report/github.com/chz8494/gored)

A Realtime-Exchange-Data SDK is supporting multiple exchanges written in Golang.

**Please note that this SDK is under heavily development and is not ready for production!**

## Community

Join our telegram to discuss all things related to GoRed! [GoRed Telegram](https://t.me/goredtalk)

## Exchange Support Table

LIST ON https://www.coindatapage.com

We are aiming to support the top 100 highest volume exchanges based off the [CoinMarketCap exchange data](https://coinmarketcap.com/exchanges/volume/24-hour/).

** NA means not applicable as the Exchange does not support the feature.

## Current Features

+ Unified all symbols / pairs into Bitontop standard.
+ Support for all Exchange fiat and digital currencies, with the ability to individually toggle them on/off.

## Planned Features

Planned features can be found on our [community Trello page](https://trello.com/gored).

## Contribution

Please feel free to submit any pull requests or suggest any desired features to be added.

When submitting a PR, please abide by our coding guidelines:

+ Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
+ Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
+ Code must adhere to our [coding style](https://github.com/chz8494/gored/blob/master/.github/CONTRIBUTING.md).
+ Pull requests need to be based on and opened against the `master` branch.

## Compiling instructions

Download and install Go from [Go Downloads](https://golang.org/dl/) for your
platform.

### Linux/OSX

gored is built using [Go Modules](https://github.com/golang/go/wiki/Modules) and requires Go 1.11 or above
Using Go Modules you now clone this repository **outside** your GOPATH

```bash
git clone https://github.com/chz8494/gored.git
cd gored
go build
mkdir ~/.gored

```

### Windows

```bash
git clone https://github.com/chz8494/gored.git
cd gored
go build

```

## Donations

<img src="" hspace="70">

If this framework helped you in any way, or you would like to support the developers working on it, please donate Bitcoin to:

***Bitcoin Address***
16sXCkKN5KegwvC53xvsahtCpPsernzQoC

***AIB Address***
ALVNJC94McFjN7enZsY2XaR6fY3GWDY3Wh




## Contributor List

### A very special thank you to all who have contributed to this program:

|User|Github|Contribution Amount|
|--|--|--|
| chunlee1991 | https://github.com/chunlee1991 
| tony0408 | https://github.com/tony0408 
| 9cat | https://github.com/9cat 
| botemple | https://github.com/botemple 
| iobond | https://github.com/iobond 

## Who used GoRED
[CoinDataPage](https://www.coindatapage.com)
