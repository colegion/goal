package controllers

// Smth is not an action as it doesn't return action.Result.
func (c *Controller) Smth() bool {
	return true
}

// SmthElse is not an action as it returns nothing.
func (c App) SmthElse() {
}

// SmthElse1 is not an action as it doesn't return action.Result.
func (c *App) SmthElse1() {
}

// Init should be ignored as the number of arguments is incorrect.
func Init() {
}

func init() {
}
