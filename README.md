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
cd ui && npm install
```

### Build frontend
```bash
cd ui && npm run build
```
