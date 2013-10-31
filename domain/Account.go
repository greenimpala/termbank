package domain

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"termbank/parser"
)

type Account struct {
	User    *User
	Name    string
	Href    string
	Balance string
}

func (a *Account) StatementPretty() []byte {
	// Obtain our cookie for the account page
	_, err := a.User.client.Get(a.Href)

	if err != nil {
		panic(err)
	}

	// Now get the export popup HTML to obtain our submit token
	res, _ := a.User.client.Get("https://secure.lloydsbank.co.uk/personal/a/viewproductdetails/m44_exportstatement.jsp")
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	submitToken := parser.SubmitToken(resBody)

	formData := url.Values{
		"frmTest:rdoDateRange":            []string{"0"},
		"frmTest:dtSearchFromDate":        []string{"-"},
		"frmTest:dtSearchFromDate.month":  []string{"-"},
		"frmTest:dtSearchFromDate.year":   []string{"-"},
		"frmTest:dtSearchToDate":          []string{"-"},
		"frmTest:dtSearchToDate.month":    []string{"-"},
		"frmTest:dtSearchToDate.year":     []string{"-"},
		"frmTest:strExportFormatSelected": []string{"Internet banking text/spreadsheet (.CSV)"},
		"frmTest":                         []string{"frmTest"},
		"submitToken":                     []string{submitToken},
		"frmTest:btn_Export":              []string{"null"},
	}.Encode()

	req, _ := http.NewRequest("POST", "https://secure.lloydsbank.co.uk/personal/a/viewproductdetails/m44_exportstatement.jsp", strings.NewReader(formData))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	memResponse, _ := a.User.client.Do(req)
	defer memResponse.Body.Close()

	statement, _ := ioutil.ReadAll(memResponse.Body)

	return statement
}
