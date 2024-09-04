import requests
import json

apiKey = ""

customerExample = {
"first_name": "Khoa",
"last_name": "Nguyen",
"address": {
    "street_number": "1680",
    "street_name": "Capital One Drive ",
    "city": "McLean",
    "state": "VA",
    "zip": "22102"
}
}

accountExample = {
"type": "Credit Card",
"nickname": "spending",
"rewards": 100,
"balance": 2000,
"account_number": "1234123412341234"
}

depositExample = {
"medium": "balance",
"transaction_date": "2016-08-29",
"amount": 100,
"description": "paycheck"
}


def createCustomer(customerData):
    payload = customerData
    url = 'http://api.reimaginebanking.com/customers?key={}'.format(apiKey)
    response = requests.post(
    url,
    data=json.dumps(payload),
    headers={'content-type':'application/json'},
    )
    if response.status_code == 201:
        print('customer created')
        return response.json()['objectCreated']['_id']
    else:
        print (response.status_code)
    return None

def createAccount(customerId, accountData):
    payload = accountData
    url = "http://api.reimaginebanking.com/customers/{}/accounts?key={}".format(customerId, apiKey)
    response = requests.post(
    url,
    data=json.dumps(payload),
    headers={'content-type':'application/json'},
    )
    if response.status_code == 201:
        print('account created')
        return response.json()['objectCreated']['_id']
    else:
        print (response.status_code)
    return None

def getAtms():
    pageNumber = 1
    while (True):
        url = 'http://api.reimaginebanking.com/atms?key={}&page={}'.format(apiKey,pageNumber)
        response = requests.get(url)
        if response.status_code == 200:
            if (response.json()['data'] == []):
                break
            print (response.json()['data'])
            pageNumber += 1
        else:
            print (response.status_code)
        break

def getBill(accountId):
    url = "http://api.reimaginebanking.com/accounts/{}/bills?key={}".format(accountId, apiKey)
    response = requests.get(url)
    if (response.status_code == 200):
        print (response.json())
    else:
        print (response.status_code)

def getPurchases(accountId):
    url = "http://api.reimaginebanking.com/accounts/{}/purchases?key={}".format(accountId, apiKey)
    response = requests.get(url)
    if (response.status_code == 200):
        print (response.json())
    else:
        print (response.status_code)

def depositMoney(accountId, depositData):
    payload = depositData
    url = "http://api.reimaginebanking.com/accounts/{}/deposits?key={}".format(accountId, apiKey)
    response = requests.post(
    url,
    data=json.dumps(payload),
    headers={'content-type':'application/json'},
    )
    if response.status_code == 201:
        print('deposit made')
    else:
        print (response.status_code)

def getEnterpriseAccounts():
    url = "http://api.reimaginebanking.com/enterprise/accounts?key={}".format(apiKey)
    response = requests.get(url)
    if (response.status_code == 200):
        print (response.json())
    else:
        print (response.status_code)

getAtms()
customer = createCustomer(customerExample)
account = createAccount(customer, accountExample)
getBill(account)
getPurchases(account)
depositMoney(account, depositExample)
getEnterpriseAccounts()