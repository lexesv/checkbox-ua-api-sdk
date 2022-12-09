# checkbox-ua-api-sdk

### (under development)

Реалізовані деякі базові функції, яких достатньо для відкриття зміни і створення фіскальних чеків.
Implemented some basic functions.

Alternative: [Auto generated Checkbox SDK](https://github.com/lexesv/checkbox-ua-api-sdk-autogen)

### Usage

Example:

```go
package main

import (
	"log"
	"github.com/lexesv/checkbox-ua-api-sdk"
)

func main() {
	LicenseKey := "<LicenseKey>"
	CashierLogin := "<CashierLogin>"
	CashierPassword := "<CashierPassword>"
	CashierPinCode := "<CashierPinCode>"

	ChBox := checkbox.New(LicenseKey, CashierLogin, CashierPassword, CashierPinCode)

	sign, err := ChBox.SignIn()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[Cashier] AccessToken:", sign.AccessToken)

	profile, err := ChBox.GetCashierProfile(sign.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[Cashier] ID: %s; FullName: %s\n", profile.Id, profile.FullName)

	rinfo, err := ChBox.GetCashRegistersInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[CashRegisters] Has open Shift:", rinfo.HasShift)

	shift, err := ChBox.CreateShift(sign.AccessToken)
	if err != nil {
		//log.Fatal(err)
		log.Println(err)
	}
	if shift.Serial > 0 {
		log.Println("Shift success created")
	}

	req := ChBox.NewSimpleReceipt("1", "Продаж ПЗ", 10000, 1000, "CARD", 10000)
	receipt, err := ChBox.CreateReceipt(sign.AccessToken, req)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Serial > 0 {
		log.Println("Receipt success created")
	}
}

```
