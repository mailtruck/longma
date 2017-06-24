![longma: dragon horse](https://upload.wikimedia.org/wikipedia/commons/6/60/Waddell-p410-Chinese-LONG-Horse-Or-Horse-Dragon.jpg)

# Longma 🐉🐎

Streaming CSV file splitter/chunker written in golang

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

## Todo

* fix license wars