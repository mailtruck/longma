# Longma 🐉🐎

![longma: dragon horse](https://upload.wikimedia.org/wikipedia/commons/6/60/Waddell-p410-Chinese-LONG-Horse-Or-Horse-Dragon.jpg)

Streaming CSV file splitter/chunker written in golang

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


## Roadmap

* flag for lines per file
* create and split files into a new directory
* support splitting csv.gz files
* flag for csv.gz or csv output
* ability to output to tar.gz or zip
* flag for lazy quotes on/off
* flag for srict records per row flags
* flag to include or exclude header/columns in file chunks
* command to combine file chunks back into a single large csv file
* make binaries available to download?
* gui?

## Todo

* fix license wars

---

⚓
