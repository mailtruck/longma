# Longma 🐉🐎

[![Build Status](https://travis-ci.org/mailtruck/longma.svg?branch=master)](https://travis-ci.org/mailtruck/longma)

![longma: dragon horse](https://upload.wikimedia.org/wikipedia/commons/6/60/Waddell-p410-Chinese-LONG-Horse-Or-Horse-Dragon.jpg)

A true dragon horse has wings at its sides and walks upon the water without sinking.

## About

Longma gives you the power to split a large CSV file into multiple files. Currently Longma includes the header/columns in every output file chunk. If you need file chunks without the header/columns open an issue and I will prioritize that feature.

## Status

Early early

## Install

Have go installed and set up.

```
go get github.com/mailtruck/longma
cd $GOPATH/src/github.com/mailtruck/longma
go install
```

## Usage

Generate a sample csv to test longma:

```
longma generateSample
```

Split a CSV file into chunks:

```
longma split example.csv
```

# Features

* support splitting csv.gz files
* flag for lines per file: -r

## Roadmap

* flag for csv.gz or csv output
* ability to output to tar.gz or zip
* flag for lazy quotes on/off
* flag for srict records per row flags
* flag to include or exclude header/columns in file chunks
* command to combine file chunks back into a single large csv file
* make binaries available to download?
  * do releases
* gui?  - wip

## To test:

What happens if you run a sample on something that has a non csv file?
What happens if you try to split on a gz thats not a csv?
What happens if you run sample on a dir twice?

---

⚓
