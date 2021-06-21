package helper_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/iamseki/grpc-vs-rest/helper"

	"github.com/stretchr/testify/assert"
)

func TestJSONFileToStruct(t *testing.T) {
	filename := "testing.json"
	ioutil.WriteFile(filename, []byte("{\"company\": \"Gorila\", \"employers\": 17}"), 0644)

	type sut struct {
		Company   string `json:"company"`
		Employers int    `json:"employers"`
	}
	var target sut
	expected := sut{Company: "Gorila", Employers: 17}

	helper.JSONFileToStruct(filename, &target)
	assert.Equal(t, expected, target)

	os.Remove(filename)
	err := helper.JSONFileToStruct(filename, &target)
	assert.NotNil(t, err)
}
