# Logcool [![Version Status](https://img.shields.io/badge/release-v0.1.0-orange.svg)](https://github.com/admpub/logcool/releases/tag/v0.1.0)

[![Build Status](https://travis-ci.org/wgliang/logcool.svg?branch=master)](https://travis-ci.org/wgliang/logcool)
[![GoDoc](https://godoc.org/github.com/admpub/logcool?status.svg)](https://godoc.org/github.com/admpub/logcool)
[![Join the chat at https://gitter.im/logcool/Lobby](https://badges.gitter.im/logcool/Lobby.svg)](https://gitter.im/logcool/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Code Health](https://landscape.io/github/wgliang/logcool/master/landscape.svg?style=flat)](https://landscape.io/github/wgliang/logcool/master)
[![Code Issues](https://www.quantifiedcode.com/api/v1/project/98b2cb0efd774c5fa8f9299c4f96a8c5/badge.svg)](https://www.quantifiedcode.com/app/project/98b2cb0efd774c5fa8f9299c4f96a8c5)
[![Go Report Card](https://goreportcard.com/badge/github.com/admpub/logcool)](https://goreportcard.com/report/github.com/admpub/logcool)
[![License](https://img.shields.io/badge/LICENSE-Apache2.0-ff69b4.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)


Logcool is an open source project to collect, filter ,transfer and response any log or event-flow data as a lightweight system.[中文](./docs/README_ZH.md)

![Logcool](./logcool.jpg)

Logcool's design learn from Heka and Logstash and it's implementation was inspired by gogstash. What's more, the logcool's goal is to be a completely independent project and not much rely on other non-standard libiaries.

Because it is difficult to fully meet the needs of different services, this repository provides basic plugins, such as encryption and decryption of data, compression and decompression of data, data format conversion, support files, command line, http, or the output of any system or redis, influxDB, MySQL database and so on. Importantly, you can easily develop a plugin according to your needs, and easily use it.

You can use logcool in any way.

## Getting started

Logcool can collect all-types los or event-flow data, and support some input/output types.Besides,you can  new your's plugs if you need it. To get started, [check out the installation instructions in the documentation](https://godoc.org/github.com/admpub/logcool).

## Using Example

A easy stdin2stdout example. 
![Logcool](./logcool.gif)

## Plugins

Some plugins that had finished and will develope in the future.

### input
- [file](https://github.com/admpub/logcool/tree/master/input/file) source data from files,such as log file.
- [stdin](https://github.com/admpub/logcool/tree/master/input/stdin) get data from the console, debugging and example will need it.
- [http](https://github.com/admpub/logcool/tree/master/input/stdin) get data from the network, support post, get, etc.
- [collectd](https://github.com/admpub/logcool/tree/master/input/collectd) monitor and control system performance data, such as CPU, memory, network, hard disk, etc.

### filter
- [zeus](https://github.com/admpub/logcool/tree/master/filter/zeus) simple label filter.
- [metrics](https://github.com/admpub/logcool/tree/master/filter/metrics) dot counter, can be used for alarm and dashboard generation.
- [grok](https://github.com/admpub/logcool/tree/master/filter/grok) regular filtering data, support multi pattern matching.
- [split](https://github.com/admpub/logcool/tree/master/filter/split) split command parameter based on the separator.

### output
- [stdout](https://github.com/admpub/logcool/tree/master/output/stdout) output to console.
- [redis](https://github.com/admpub/logcool/tree/master/output/redis) enter data into the redis database.
- influxdb data import influxdb, this is useful for timing data
- [email](https://github.com/admpub/logcool/tree/master/output/email) send messages via email, such as alerts and service exception notifications.
- [lexec](https://github.com/admpub/logcool/tree/master/output/lexec) send a message to execute a command or script.
- mysql Write data to MySQL
- pg Write data to pg

## Versions

[versions](https://github.com/admpub/logcool/blob/master/docs/VERSION_UPDATE.md)

## Other Contributor

Logcool learn from gogstash much. Thank you for your contribution, and I also learn a lot from your project. @tsaikd

## Licensing

Logcool is licensed under the Apache License, Version 2.0. See LICENSE for the full license text.

## Welcome to Contribute

1.Fork it
2.Create your feature branch
3.Commit your changes (git commit -am 'Add some feature'),and no test error.
4.Push to the branch
5.Create new Pull Request

Documentation or correcting comments are also welcome.
