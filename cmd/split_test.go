package cmd

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/mailtruck/longma/generate"
	"github.com/mailtruck/longma/utils"
	"github.com/stretchr/testify/assert"
)

// incomplete test
func TestSlit(t *testing.T) {
	testFile := "test.csv"
	err := generate.GenerateCSV(testFile, int64(10000))
	assert.Equal(t, nil, err)

	folder, err := split(testFile, 1001, false)

	_, err = os.Stat(folder)

	assert.Equal(t, true, !os.IsNotExist(err))
}

// This should be split up
func TestGetFolderName(t *testing.T) {
	filename := "filename.csv"

	expected := "filename-split"

	//                        ↓ trying out this style... I think it makes the tests less understandable
	assert.Equal(t, expected, getFolderName(filename))

	filename = "dingo.csv"
	expected = "dingo-split"
	assert.Equal(t, expected, getFolderName(filename))

	filename = "dingo.csv.gz"
	assert.Equal(t, expected, getFolderName(filename))

	if _, err = os.Stat("dingo-split"); err == nil {
		log.Fatal(errors.New("for this test you really want to make sure this file isn't here: " + filename))
	}

	err := os.Mkdir("dingo-split", 0777)
	assert.Equal(t, nil, err)
	defer utils.Cleanup("dingo-split")

	err = generate.GenerateCSV(filename, 2)
	assert.Equal(t, nil, err)
	defer utils.Cleanup(filename)

	expected = "dingo-split_1"

	assert.Equal(t, expected, getFolderName(filename))

	if _, err = os.Stat("dingo-split_1"); err == nil {
		log.Fatal(errors.New("for this test you really want to make sure this file isn't here: " + filename))
	}

	err = os.Mkdir("dingo-split_1", 0777)
	assert.Equal(t, nil, err)
	defer utils.Cleanup("dingo-split_1")

	expected = "dingo-split_2"

	assert.Equal(t, expected, getFolderName(filename))

}

func TestGetBaseBase(t *testing.T) {
	base := "dingo.csv"
	expected := "dingo"
	bb, err := getBaseBase(base)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, bb)

	base = "dingo.csv.gz"
	bb, err = getBaseBase(base)
	assert.Equal(t, expected, bb)
	assert.Equal(t, nil, err)
}
