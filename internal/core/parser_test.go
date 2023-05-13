package core_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/suportebeloj/desafio-dev/internal/core"
	"reflect"
	"testing"
)

func TestNewTransactionParser(t *testing.T) {
	instance := core.NewTransactionParser()
	assert.NotNil(t, instance)

	v := reflect.ValueOf(instance)
	assert.True(t, v.MethodByName("Parser").IsValid())
}

func TestTransactionParser_Parser(t *testing.T) {
	instance := core.NewTransactionParser()
	transactionTest := "3201903010000014200096206760174753****3153153453JOÃO MACEDO   BAR DO JOÃO       "
	parsedObject, err := instance.Parser(transactionTest)
	assert.NoError(t, err)
	assert.NotNil(t, parsedObject)
}

func TestTransactionParser_GivenError_WhenParse_InvalidTransaction(t *testing.T) {
	instance := core.NewTransactionParser()
	transactionTest := "a201903010000014200096206760174753****3153153453JOÃO MACEDO   BAR DO JOÃO       "
	_, err := instance.Parser(transactionTest)
	assert.Error(t, err)
}
