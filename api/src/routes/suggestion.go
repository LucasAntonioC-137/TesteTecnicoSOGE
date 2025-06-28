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
	},
	{
		URI: "/register",
		Method: http.MethodPost,
		Function: controller.CreateSuggestion,
	},
	{
		URI: "/suggestions/{id}/status",
		Method: http.MethodPut,
		Function: controller.UpdateSuggestionStatus,
	},
}