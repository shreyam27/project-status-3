package main

func statusHandler(w http.ResponseWriter, r *http.Request) {
	statusTemplate.Execute(w, fmt.Sprintf("Project status 2"))
}
