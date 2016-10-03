type Widget struct {
    Description string
    Price       int
}

func handle(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
    q := datastore.NewQuery("Widget")

       q= q.Filter("Price <", 1000).
        q= q.Order("-Price")
    b := new(bytes.Buffer)
    for t := q.Run(ctx); ; {
        var x Widget
        key, err := t.Next(&x)
        if err == datastore.Done {
            break
        }
        if err != nil {
            serveError(ctx, w, err)
            return
        }
        fmt.Fprintf(b, "Key=%v\nWidget=%#v\n\n", key, x)
    }
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    io.Copy(w, b)
}