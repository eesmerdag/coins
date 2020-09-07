# Coin Api Repository 

There is 3 different services provided by this repository:
1. api (this one is assumed only api to outside of repository)
2. pricing service
3. ranking service

For detailed information, please read read.md files in sub directories!

## How to start the application on local?
## On terminal
go to directory and execute these commands in different tabs by this sort:
1. go run ranking-service/main/service.go
2. go run pricing-service/main/service.go
3. go run api/main/service.go
## Docker
go to directory and execute these commands to create images:

```docker build -f ./Dockerfile_ranking -t ranking:latest .```

```docker build -f ./Dockerfile_pricing -t pricing:latest .```

```docker build -f ./Dockerfile_api -t api:latest .```

you can run images as containers with these commands:

 ```docker run  -p  127.0.0.1:1903:1903 container_id_of_pricing```
 
  ```docker run  -p  127.0.0.1:1904:1904 container_id_of_ranking```
  
  ```docker run  -p  127.0.0.1:1902:1902 container_id_of_api```
  
Note: Dockercompose should be used for orchestrating. This needs to be configured but now I don't have enough time.