# touchnote
1. Create a database with name sample (any name you choose)
1.1 If are using different name then please make changes in connection/connection.go.
1.2 I have used my loginID password in that connection please make changes and put your mysql password.(mostly it is root).

2. There are 5 different API:
2.1 "/hello" :
2.2.1. This API is for enrollment of customer into the Database.
2.2.2 Sample JSON request is :
{
   "Name" :"amey",
   "EmailId":"ameymanjkr@gmail.com",
   "Balance" : 100,
   "Uuid" : 800156531
}
2.2.3 Note kindly use UUID 9 digit number. (it is set to be unique in database).


2.2	"/credits":
2.2.1 This API is for crediting the amount to your account.
2.2.2 Sample JSON request is :
{
 "name":"credit",
 "payload":{
 "uuid":800156531,
 "amount":15,
 "type":"subscription",
 "priority":1,
 "expiry":"1032323230"
 }
}

2.3	"/debits" : 
2.3.1 This API is for debiting the amount from your account
2.3.2 Sample request i :
{
 "name":"debit",
 "payload":{
 "uuid":800156531,
 "amount":10
 }
}
2.4 "/getAll" : This is a Get request which feteches all the transaction of all users

2.5	"getAll/:uuid" : This is a get request of particular user using UUID.
