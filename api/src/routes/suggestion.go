package routes

import (
	 "net/http"
	 "go-api/src/controller"
)
var suggestionRoutes = []Route{
	{
		URI: "/suggestions",
		Method: http.MethodGet,
		Function: controller.GetSuggestionsWithFilters,
		AuthISRequired: false,
	},
	{
		URI: "/register",
		Method: http.MethodPost,
		Function: controller.CreateSuggestion,
		AuthISRequired: false,
	},
}