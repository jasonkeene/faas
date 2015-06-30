
## Fractals as a Service (FAAS)

To run the server do this:

```bash
go run server-zach/server.go
```

or:

```bash
go run server-nick/server.go
```

Once the server is running you can churn out fractals like there is no
tomorrow.

```bash
curl "http://127.0.0.1:9017/mandelbrot?width=1024&height=768&i=68&z=0.9"
```
