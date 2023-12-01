package appContext

type AppContext struct {
	Key string
}

var (
	CurrentUserKey = &AppContext{Key: "CurrentUser"}
)
