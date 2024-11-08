package main

import(
    "net/http"
    "runtime/pprof"
)

func dumpGoroutines(w http.ResponseWriter, r *http.Request, t auth.Token) error {
    if !permission.Check(t, permission.PermDebug) {
        return permission.ErrUnauthorized
    }

    // ruleid: write-pprof-profile-output
    return pprof.Lookup("goroutine").WriteTo(w, 2)
}
