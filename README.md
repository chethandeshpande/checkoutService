####To run tests:
```
make -B test
```

####To run the app:
```
make build
make run
```

The prompt will ask you to enter the products and quantity you want to buy
####Products Available:
1. Google Home: $49.99
2. Macbook Pro: $5399.99
3. Alexa Speaker: $109.50
4. Raspberry Pi: $30.00

####Promotions Available:
- Each sale of a MacBook Pro comes with a free Raspberry Pi B
- Buy 3 Google Homes for the price of 2
- Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers

Please select products and quantity:
Enter in ```1:quantity,2:quantiy,3:quantity``` format to select products and quantity:

####Example:
If you want to buy 2 Alexa speakers and one macbook pro, the command will look like
```3:2,2:1```

You can also add test cases to test your scenarios. Please add tests in test/checkout_test.go file.

####Assumptions made:
- Only backend is involved
- There is no inventory management
