# cookiecutter-golambda

cookiecutter-golambda is an AWS Lambda function generator. It creates opinionated structure for serverless functions written in Golang. cookiecutter-golang follows Hexagonal Architecture aka Ports and Adapters pattern, you can read more [here](http://codingcanvas.com/hexagonal-architecture/). 

# Requirements 
* [Cookiecutter](https://github.com/audreyr/cookiecutter)
* [Docker](https://www.docker.com)
* [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html) 
* Golang 1.12.x and above

# Usage
```cookiecutter https://github.com/Ryanair/cookiecutter-golambda```
![cookiecutter-golambda](https://github.com/Ryanair/cookiecutter-golambda/blob/master/cookiecutter-golambda.gif)

# Features
* Run your AWS Lambda function locally and use API Gateway as a trigger. 
* Use either DynamoDB or in memory storage adapter for your persistence layer.
* Execute by using pre-defined set of Makefile commands. 
* Handle environment variables by using [envconfig](https://github.com/kelseyhightower/envconfig)
* Use Uber's [zap](https://github.com/uber-go/zap) logging library.

# Built with
[Cookiecutter](https://github.com/audreyr/cookiecutter)

# License
cookiecutter-golambda is under the MIT license
```The MIT License (MIT)
Copyright (c) 2019 Ryanair Labs
 
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
associated documentation files (the "Software"), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge, publish, distribute,
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software
is furnished to do so, subject to the following conditions:
 
The above copyright notice and this permission notice shall be included in all copies or
substantial portions of the Software.
 
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

