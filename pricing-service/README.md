# Pricing service
This service has only an api to list up-to-date list of current prices in USD of all crypto currencies.

Data source:
https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
Default limit is 100, so it is required to add limit while sending request. In our code, max value which 5000 is used.

For more information and all services by this source, please read: https://coinmarketcap.com/api/documentation/v1/

This service runs at port 1903. This could be changed in main/service.go.
There is no any intention to choose this port. It is just birth year of my favourite team :) 
 
