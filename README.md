[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

# Go Session Handling Example

This is a simple "application" to demonstrate the use of the [Gorilla Session](http://www.gorillatoolkit.org/pkg/sessions) package.

Check out the [live demo](https://go-sessions-demo.herokuapp.com/) and [read the docs](https://devcenter.heroku.com/articles/go-sessions).

## Startup

```sh
export SESSION_AUTHENTICATION_KEY="your-authentication-key"
export SESSION_ENCRYPTION_KEY="0123456789abcdef0123456789abcdef"
export PORT="18080"
```

```sh
go run main.go
```

## Session Cookies

Default values for `sessions.NewCookieStore`

* Path="/" 
* Domain=""
* MaxAge=2592000
* Secure=true
* Partitioned=false
* HttpOnly=false
* SameSite=4(None)
