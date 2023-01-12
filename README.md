# Auth serverless for Confess app
## Description
- An Authen serverless run on AWS Lambda support for confess app
## API
- Register: register new account and return token for client
- Login: login to account and return token for client
- Verify token: verify token is usable and expire time
- Logout //maybe not

## Todo/Advanced feature
- One step verify by email when register
- Control token with caching (disable login multi device)