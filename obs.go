package main

import (
	"net/http"

	"html/template"
)

var displayTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script src="assets.js"></script>
<script>  
var defaultFont = "arial";
var defaultBgColor = "#007030";
var secondsDuration = "30s"; 
window.addEventListener("load", function(evt) {
    var ws;
    ws = new WebSocket("{{.}}");
    ws.onopen = function(evt) {
        console.log("Connection established");
    }
    ws.onclose = function(evt) {
        alert("Cannot establish the connection with the local server.");
        console.log("Error encountered. Event:");
        console.log(evt);
    }
    ws.onerror = function(evt) {
        console.log("Encountered: " + evt.data);
    }
    ws.onmessage = function(evt) {
        console.log("Got: " + evt.data);
    }

    });
</script>
`))

func serveObs(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "obs.html")
	//displayTemplate.Execute(w, "ws://"+r.Host+"/ws")
}
