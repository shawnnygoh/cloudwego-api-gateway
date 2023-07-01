# CloudWeGo-API-Gateway
An API Gateway made with CloudWeGo frameworks [Hertz](https://github.com/cloudwego/hertz#readme) and [Kitex](https://github.com/cloudwego/kitex#readme).

## Basic Features
- Generic Call with Kitex
- Service Registry & Discovery with Nacos
- Load Balancing with Kitex
- Service Mapping

### Directory Organization 
```
├── hertz-gateway
├── idl 
│   └── echo.thrift
├── kitex-services
├── LICENSE
├── README.md
```

## Set-up
To get started, users will have to install Nacos [here](https://nacos.io/en-us/docs/v2/quickstart/quick-start.html) first in order to utilize the API Gateway. 