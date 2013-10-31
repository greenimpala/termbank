# termbank

Bank like it's 1960. View account details and statements in the terminal.

## Supported Banks

* Lloyds

## Usage

Just run `termbank` and follow the instructions.

Typical program run:

```bash
$ termbank
User ID: *******
Password: *******
Memorable Info: *******

Logged in!

1) Account 1 - £xxxx.xx
2) Account 2 - £xxxx.xx
3) Account 3 - £xxxx.xx

Enter valid account number: 1

Account: Account 1
Balance: £xxxx.xx
Fetching statement...

Transaction 	Transaction 	Sort Code	Account Numb	Transaction 	Debit Amount	Balance
30/10/2013  	DEB         	'99-99-99	12345678    	TESCO STORES	5.15        	1006.64
29/10/2013  	DEB         	'99-99-99	12345678    	TESCO STORES 	5.63        	1011.79
...
```