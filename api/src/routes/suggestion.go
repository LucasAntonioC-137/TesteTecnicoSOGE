package routes

import (
	 "net/http"
	 "go-api/src/controller"
)

var suggestionRoutes = []Route{
	{
		URI: "/suggestions",
		Method: http.MethodGet,
		Function: controller.GetSuggestions,
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
	{
		URI: "/suggestions/grouped-by-status",
		Method: http.MethodGet,
		Function: controller.GetSuggestionsGroupedByStatus,
	},
	{
		URI: "/suggestions/grouped-by-sector",
		Method: http.MethodGet,
		Function: controller.GetSuggestionsGroupedBySector,
	},	
}