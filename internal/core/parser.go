package core

import (
	"encoding/json"
	"fmt"
	"github.com/suportebeloj/desafio-dev/internal/db/postgres"
	"github.com/suportebeloj/desafio-dev/internal/utils/cerrors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type TransactionParser struct {
}

func (t *TransactionParser) Parser(transactionString string) (*postgres.CreateTransactionParams, error) {
	newTransaction := new(postgres.CreateTransactionParams)

	parsedTransaction, err := t.parser(transactionString)
	if err != nil {
		return nil, err
	}

	if err = t.validateFields(parsedTransaction); err != nil {
		return nil, err
	}

	if err = t.mapToStruct(parsedTransaction, newTransaction); err != nil {
		return nil, err
	}

	return newTransaction, nil
}

func (t TransactionParser) parser(text string) (map[string]interface{}, error) {

	parsedTransaction := map[string]interface{}{}

	patter := `(?P<type>[\d]{1})(?P<date>[\d]{8})(?P<value>[\d]{10})(?P<cpf>[\d]{11})(?P<card>[\d\S]{12})(?P<time>[\d]{6})(?P<owner>[\s\S]{14})(?P<market>[\s\S]{18})`
	re := regexp.MustCompile(patter)
	groupNames := re.SubexpNames()[1:]

	match := re.FindStringSubmatch(text)
	if match == nil {
		return nil, cerrors.TransactionNotMatchError{}
	}
	match = match[1:]

	for i, s := range groupNames {
		parsedTransaction[s] = match[i]
	}
	return parsedTransaction, nil
}

func (t TransactionParser) mapToStruct(m map[string]interface{}, s interface{}) error {
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(j, s); err != nil {
		return err
	}

	return nil
}

func (t TransactionParser) validateFields(m map[string]any) error {
	for key, value := range m {
		if key == "date" {
			date, err := t.convertStrToTime(t.formatDate(value.(string)), "2006-01-02")
			if err != nil {
				return err
			}
			m[key] = date
		}

		if key == "value" {
			converted, err := strconv.ParseFloat(value.(string), 64)
			if err != nil {
				return err
			}

			m[key] = converted
		}

		if key == "time" {
			date, err := t.convertStrToTime(t.formateTime(value.(string)), "15:04:05")
			if err != nil {
				return err
			}

			m[key] = date
		}

		if key == "owner" {
			m[key] = strings.TrimSpace(value.(string))
		}

		if key == "market" {
			m[key] = strings.TrimSpace(value.(string))
		}

	}

	return nil
}

func (t TransactionParser) formatDate(value string) string {
	pattern := `(?P<year>[\d]{4})(?P<month>[\d]{2})(?P<day>[\d]{2})`
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(value)[1:]

	return fmt.Sprintf("%s-%s-%s", match[0], match[1], match[2])
}

func (t TransactionParser) formateTime(value string) string {
	pattern := `(?P<hours>[\d]{2})(?P<minutes>[\d]{2})(?P<seconds>[\d]{2})`
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(value)[1:]

	return fmt.Sprintf("%s:%s:%s", match[0], match[1], match[2])
}

func (t TransactionParser) convertStrToTime(values, layout string) (time.Time, error) {
	parsed, err := time.Parse(layout, values)
	if err != nil {
		return time.Time{}, err
	}

	return parsed, nil
}

func NewTransactionParser() *TransactionParser {
	return &TransactionParser{}
}
