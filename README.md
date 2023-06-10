# go-gin-basicauth-monolithic-template
Monolithic server side Golang template with gin web framework, postgres and basic authorzation middleware.


# Instructions to run the source code
Follow the instractions to run the code in your local machine.
MakeFile, Docker engine and docker compose should be installed in your local machine. 

To run the program
```
make compose_up
```
To stop the program
```
make compose_down
```

After program is started successfully, you can check if it is running using this address.
```
http://localhost:8000/v1/swagger/index.html
```