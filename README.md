# touchnote
1. Create a database with name sample (any name you choose)
2. If are using different name then please make changes in connection/connection.go.
4. I have used my loginID password in that connection please make changes and put your mysql password.(mostly it is root).
5. To run the server use go run main.go from src/api folder.
6. The server will up on port 8080.
7. There are 5 different API:
8. "/hello" :
9. This API is for enrollment of customer into the Database.
10. Sample JSON request is :{
   "Name" :"amey",
   "EmailId":"ameymanjkr@gmail.com",
   "Balance" : 100,
   "Uuid" : 800156531
}
8. Note kindly use UUID 9 digit number. (it is set to be unique in database).
9.	"/credits":
10. This API is for crediting the amount to your account.
11. Sample JSON request is :{
 "name":"credit",
 "payload":{
 "uuid":800156531,
 "amount":15,
 "type":"subscription",
 "priority":1,
 "expiry":"1032323230"
 }
}
12.	"/debits" : 
13. This API is for debiting the amount from your account
14. Sample request for Debit :{
 "name":"debit",
 "payload":{
 "uuid":800156531,
 "amount":10
 }
}
15."/getAll" : This is a Get request which feteches all the transaction of all users
16	"getAll/:uuid" : This is a get request of particular user using UUID.
