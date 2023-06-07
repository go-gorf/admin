# Gorf Admin

### Installation

```bash
go get https://github.com/go-gorf/admin@latest
```

### Add to Gorf App

```go
var apps = []gorf.GorfApp{
	&admin.AdminApp,
}
```
### Install frontend dependencies
```bash
cd frontend && npm install
```

### Build frontend
```bash
cd frontend && npm run build
```
