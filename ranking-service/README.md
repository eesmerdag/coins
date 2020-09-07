# Ranking service
This service has only an api to list up-to-date list of top assests.

Data source:
https://min-api.cryptocompare.com/documentation?key=Toplists&cat=TopTotalMktCapEndpointFull
(note: I am not %100 sure, this is the correct api :S. But I feel the solution is more important than data)

There is two parameters we have to add to uri:
1. tsym: The currency symbol (statically defined as USD in code as)
2. limit: number of top coins 
Default limit is 10, so it is required to add limit while sending request. In our code, max value which 5000 is used. Luckily, the number of all crypto currencies less than it.

For more information and all services by this source, please read: https://min-api.cryptocompare.com/documentation

This service runs at port 1904. This could be changed in main/service.go.
 
