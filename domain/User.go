package domain

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/st3redstripe/termbank/parser"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const (
	primaryLoginGet   = "https://online.lloydsbank.co.uk/personal/logon/login.jsp"
	primaryLoginPost  = "https://online.lloydsbank.co.uk/personal/primarylogin"
	memorableInfoPost = "https://secure.lloydsbank.co.uk/personal/a/logon/entermemorableinformation.jsp"
)

type User struct {
	Accounts  []*Account
	client    *http.Client
	id        string
	password  string
	memorable string
}

func NewUser(credentials map[string]string) *User {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	return &User{
		Accounts:  []*Account{},
		client:    client,
		id:        credentials["id"],
		password:  credentials["password"],
		memorable: credentials["memorable"],
	}
}

func (u *User) Login() error {
	initial, _ := u.client.Get(primaryLoginGet)
	initialBody, _ := ioutil.ReadAll(initial.Body)
	submitToken := parser.SubmitToken(initialBody)

	body := url.Values{
		"frmLogin:strCustomerLogin_userID": []string{u.id},
		"frmLogin:strCustomerLogin_pwd":    []string{u.password},
		"submitToken":                      []string{submitToken},
	}.Encode()

	req, _ := http.NewRequest("POST", primaryLoginPost, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	loginResponse, _ := u.client.Do(req)
	defer loginResponse.Body.Close()

	memorableBody, _ := ioutil.ReadAll(loginResponse.Body)
	memorableChars := parser.MemorableCharacters(memorableBody, u.memorable)
	submitToken2 := parser.SubmitToken(memorableBody)

	var memInfoFormKey = "frmentermemorableinformation1:strEnterMemorableInformation_memInfo"
	formData := url.Values{
		memInfoFormKey + "1":                        []string{"&nbsp;" + memorableChars[0]},
		memInfoFormKey + "2":                        []string{"&nbsp;" + memorableChars[1]},
		memInfoFormKey + "3":                        []string{"&nbsp;" + memorableChars[2]},
		"frmentermemorableinformation1":             []string{"frmentermemorableinformation1"},
		"submitToken":                               []string{submitToken2},
		"frmentermemorableinformation1:btnContinue": []string{"null"},
	}.Encode()

	homePageRequest, _ := http.NewRequest("POST", memorableInfoPost, strings.NewReader(formData))
	homePageRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	homePageResponse, _ := u.client.Do(homePageRequest)
	defer homePageResponse.Body.Close()
	buildAccountsFromHomePage(homePageResponse, u)

	return nil
}

func buildAccountsFromHomePage(homePageResponse *http.Response, u *User) {
	homePageDocument, _ := goquery.NewDocumentFromResponse(homePageResponse)
	accountListItems := homePageDocument.Find(".myAccounts>li")

	u.Accounts = make([]*Account, accountListItems.Length())

	accountListItems.Each(func(i int, s *goquery.Selection) {
		u.Accounts[i] = new(Account)
		u.Accounts[i].User = u

		a := s.Find(".accountDetails h2 a")
		href, _ := a.Attr("href")

		u.Accounts[i].Name = a.Text()
		u.Accounts[i].Href = "https://secure.lloydsbank.co.uk" + href

		balanceSpans := s.Find(".balance span")

		// Two spans indicates we have a balance in the last span
		if balanceSpans.Length() == 2 {
			u.Accounts[i].Balance = balanceSpans.Last().Text()
		} else {
			// Otherwise the account has a nil ballance
			u.Accounts[i].Balance = "Empty"
		}
	})
}
