# Blogo

Blog in Go

## Usage

1. Clone the entire project from GitHub and install denpendencies:

        go get github.com/MrHuxu/blogo
        cd $GOPATH && cd src/github.com/MrHuxu/blogo/
        npm i
        go get github.com/codegangsta/gin

2. Launch the server in DEV mode:

        npm run dev

   - The Go server will be hot reloaded due to any changes;
   - The server will listen localhost at port `8283`.


3. Launch the server in PRD mode:

        npm run prd

   - The frontend files will be compressed and the Go server will be compiled;
   - The server will listen localhost at port `13109`.

