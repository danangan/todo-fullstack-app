package appContext

type AppContext struct {
	key string
}

var (
	CurrentUserKey = &AppContext{key: "CurrentUser"}
)
