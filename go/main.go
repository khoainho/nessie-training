package main

import (
		"bytes"
		"encoding/json"
		"fmt"
		"io/ioutil"
		"net/http"
	)

var apiKey = "" // API key val

type Address struct {
		StreetNumber string `json:"street_number"`
		StreetName   string `json:"street_name"`
		City         string `json:"city"`
		State        string `json:"state"`
		Zip          string `json:"zip"`
	}

type Customer struct {
		FirstName string  `json:"first_name"`
		LastName  string  `json:"last_name"`
		Address   Address `json:"address"`
	}

type Account struct {
		Type         string `json:"type"`
		Nickname     string `json:"nickname"`
		Rewards      int    `json:"rewards"`
		Balance      int    `json:"balance"`
		AccountNumber string `json:"account_number"`
	}

type Deposit struct {
		Medium         string `json:"medium"`
		TransactionDate string `json:"transaction_date"`
		Amount         int    `json:"amount"`
		Description    string `json:"description"`
	}

func getAtms() {
		pageNumber := 1
		for {
				url := fmt.Sprintf("http://api.nessieisreal.com/atms?key=%s&page=%d", apiKey, pageNumber)
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					break
				}
				defer resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
						var data map[string]interface{}
						body, _ := ioutil.ReadAll(resp.Body)
						json.Unmarshal(body, &data)
						if len(data["data"].([]interface{})) == 0 {
								break
						}
						fmt.Println(data["data"])
						pageNumber++
				} else {
						fmt.Println(resp.StatusCode)
						break
				}
			}
}

func createCustomer(customerData Customer) *string {
		url := fmt.Sprintf("http://api.nessieisreal.com/customers?key=%s", apiKey)
		payload, _ := json.Marshal(customerData)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
				fmt.Printf("Error: %v\n", err)
				return nil
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusCreated {
				var data map[string]interface{}
				body, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(body, &data)
				fmt.Println("customer created")
				id := data["objectCreated"].(map[string]interface{})["_id"].(string)
				return &id
		} else {
				fmt.Println(resp.StatusCode)
				return nil
		}
}

func createAccount(customerId string, accountData Account) *string {
		url := fmt.Sprintf("http://api.nessieisreal.com/customers/%s/accounts?key=%s", customerId, apiKey)
		payload, _ := json.Marshal(accountData)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
				fmt.Printf("Error: %v\n", err)
				return nil
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusCreated {
				var data map[string]interface{}
				body, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(body, &data)
				fmt.Println("account created")
				id := data["objectCreated"].(map[string]interface{})["_id"].(string)
				return &id
		} else {
				fmt.Println(resp.StatusCode)
				return nil
		}
}

func getBill(accountId string) {
		url := fmt.Sprintf("http://api.nessieisreal.com/accounts/%s/bills?key=%s", accountId, apiKey)
		resp, err := http.Get(url)
		if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
				var data map[string]interface{}
				body, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(body, &data)
				fmt.Println(data)
		} else {
				fmt.Println(resp.StatusCode)
		}
}

func getPurchases(accountId string) {
		url := fmt.Sprintf("http://api.nessieisreal.com/accounts/%s/purchases?key=%s", accountId, apiKey)
		resp, err := http.Get(url)
		if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
				var data map[string]interface{}
				body, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(body, &data)
				fmt.Println(data)
		} else {
				fmt.Println(resp.StatusCode)
		}
}

func depositMoney(accountId string, depositData Deposit) {
		url := fmt.Sprintf("http://api.nessieisreal.com/accounts/%s/deposits?key=%s", accountId, apiKey)
		payload, _ := json.Marshal(depositData)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusCreated {
				fmt.Println("deposit made")
		} else {
				fmt.Println(resp.StatusCode)
		}
}

func getEnterpriseAccounts() {
		url := fmt.Sprintf("http://api.nessieisreal.com/enterprise/accounts?key=%s", apiKey)
		resp, err := http.Get(url)
		if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
				var data map[string]interface{}
				body, _ := ioutil.ReadAll(resp.Body)
				json.Unmarshal(body, &data)
				fmt.Println(data)
		} else {
				fmt.Println(resp.StatusCode)
		}
}

func deleteData(entityType string, val string) {
		url := fmt.Sprintf("http://api.nessieisreal.com/data?type=%s&key=%s", entityType, val)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
				fmt.Printf("Error creating DELETE request: %v\n", err)
				return
		}
		req.Header.Set("Accept", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
				fmt.Printf("Error making DELETE request: %v\n", err)
				return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
				fmt.Println("Data deleted successfully.")
		} else {
				fmt.Printf("Failed to delete data. Status code: %d\n", resp.StatusCode)
		}
}

func main() {
// paginated get ATMs call
		getAtms()
		// create the lineage customer->account->bill, purchase, deposit
		customer := createCustomer(Customer{
				FirstName: "John",
				LastName:  "Doe",
				Address: Address{
						StreetNumber: "1680",
						StreetName:   "Capital One Drive",
						City:         "McLean",
						State:        "VA",
						Zip:          "22102",
				},
		})
		if customer != nil {
				account := createAccount(*customer, Account{
						Type:         "Credit Card",
						Nickname:     "spending",
						Rewards:      150,
						Balance:      2000,
						AccountNumber: "1234123412341234",
				})
				if account != nil {
						getBill(*account)
						getPurchases(*account)
						depositMoney(*account, Deposit{
								Medium:         "balance",
								TransactionDate: "2016-08-29",
								Amount:         100,
								Description:    "paycheck",
						})
						deleteData("Accounts", "1234123412341234")
				}
		}
		getEnterpriseAccounts()
}