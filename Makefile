run:
	templ generate
	#tailwindcss -i static/css/input.css -o static/css/output.css
	#open http://localhost:42069
	@go run cmd/main.go

