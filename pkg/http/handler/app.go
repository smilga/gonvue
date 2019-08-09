package handler

import (
	"gonvue/pkg/config"
	"fmt"
	"net/http"
)

func AppHandler(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="en">
				<head>
					<title>%s</title>
					<meta charset="utf-8">
					<meta http-equiv="X-UA-Compatible" content="IE=edge">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
				</head>
				<body>
					<div id="app">
						<router-view></router-view>
					</div>
					<script type="text/javascript" src="%s"></script>
				</body>
			</html>
		`, config.Name, config.JSPath)
	}
}
