# API
This is the only service which we assume it has a connection to outside as API.

This service runs at port 1902. This could be changed in main/service.go.

When there is a request (GET), 

after running api on your local, you can test : http://localhost:1902/list/?limit=10

#Algorithm

##1. Validation of Limit:
###case 1: 
if limit is not defined, error again. Error message 'limit is not defined'
###case 2: 
if limit is not integer,Error message is 'limit must be integer'
###case 3: 
if limit is set less than 10 or more than 100, you should not see error. Error message: 'limit must be between 10 and 100'
###case 4: 
if limit is an integer between 10 and 100

##2. Calling Ranking service:
###case 1: 
if response from ranking service is not successful, show the error to the api user
###case 2: 
if response from ranking service is successful, keep the data and allow program for the flow

##3. Calling Pricing service:
###case 1: 
if response from ranking service is not successful, show the error to the api user
###case 2: 
if response from ranking service is successful, keep the data and allow program for the flow

##3. Combine data from services
Combination algorithm:
1. convert prices from array to map. It is because of lookup with map is faster than array.
2. for each rank, lookup in prices map and match and value price,
   If rank is not in lookup map, then price is -1
3. User should see combination result


#What could be improved?

1. Unit tests! even though i tried to cover all cases, there are some cases to be tested. 
2. Maybe naming conventions should be better?
3. Algorithm is not the best for sure. The best way is to fetch related information from one of those ranking or pricing servi
   ce. However, this is not requirement.
   But it could be better if we can add parameters to the pricing service. To explain it better, we maybe able to ask only related coins prices to pricing service.
   However, it is not possible, or i miss something?  
4. A very good code reviewer would help a lot :) 
5. Logging is omitted for the task, there must be logs but i don't have enough time :( 
