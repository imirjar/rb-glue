# RB_Glue

## Authentification
### Use it in combination with [rb_auth](https://github.com/imirjar/rb-auth)  service. 
### Use User struct to authentificate ID, Roles and Groups wich expexted to be.
```golang
router.Use(authentication.Authenticate(authpath, authentication.User{}))
```

## Contype 
### Use for catch unsuported requests types.
```golang
router.Use(contype.REST("application/json"))
```

## Logger
### Use for logging your requests.
```golang
router.Use(logger.Logger())
```