vercel-dev:
	vercel dev

tailwind-watch:
	tailwindcss --watch -o assets/built.css

go-dev:
	PORT=3000 go run main.go

docker-build:
	docker build -t blogo .

docker-run:
	docker run -p 13109:13109 blogo