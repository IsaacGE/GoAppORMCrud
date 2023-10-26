package handlers

import (
	"GoCrudORM/types"
	"io/ioutil"
	"strings"
)

/**
 * Get the html tamplate for the home view
 * @since 2023-10-20
 */
func GetHomeViewTemplate() string {
	htmlBytes, err := ioutil.ReadFile("templates/homeTemplate.html")
	if err != nil {
		return GetErrorViewTemplate(err)
	}

	htmlContent := string(htmlBytes)
	return htmlContent
}

/***
 * Get the html tamplate for the user created
 * @since 2023-10-20
 */
func GetUserCreatedViewTemplate(userData types.User) string {
	htmlBytes, err := ioutil.ReadFile("templates/userCreatedTemplate.html")
	if err != nil {
		return GetErrorViewTemplate(err)
	}

	htmlContent := string(htmlBytes)
	return htmlContent
}

/**
 * Get the html tamplate for the error view
 * @since 2023-10-20
 */
func GetErrorViewTemplate(inputError error) string {
	htmlBytes, err := ioutil.ReadFile("templates/errorTemplate.html")
	if err != nil {
		return `
			<!DOCTYPE html>
			<html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><title>Error</title>
			</head>
			<body>
				<h1>An error occurred while processing the request or the resource was not found.</h1>
				<h3>Also the error template was not found</h3>
				<p>Error details: ` + err.Error() + ` 
			</body>
			</html>
		`
	}

	htmlContent := string(htmlBytes)
	htmlContent = strings.Replace(htmlContent, "{errorDetails}", inputError.Error(), -1)
	return htmlContent
}
