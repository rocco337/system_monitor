"use strict"

window.onload = function () {
    if (!window["WebSocket"]) {
        alert("Sockets not supported by browser!")
    }

    var baseUrl = "ws://" + document.location.host;

    listen(baseUrl + "/hostInfo", function (evt) {
        document.querySelector("#hostInfo").innerHTML = evt.data;
    });

    listen(baseUrl + "/memorystat", function (evt) {
        document.querySelector("#memorystat").innerHTML = evt.data;
    })

    listen(baseUrl + "/processes", renderProcesses)
    listen(baseUrl + "/cpuusage", renderCpuUsage)

};

function listen(url, onmessage) {
    var conn = new WebSocket(url);
    conn.onclose = function (evt) { connectionClosed(evt) }
    conn.onmessage = function (evt) {
        console.log("received: ", evt);
        onmessage(evt);
    };

}

function webSocketsNotAvailable() {
    alert("Your browser does not support WebSockets.");
}

function connectionClosed(evt) {
    console.log(evt);
    //alert("Connection closed!")
}

function renderProcesses(evt) {
    var parsed = JSON.parse(evt.data)

    var table = document.createElement("table")
    parsed.forEach(element => {
        var row = document.createElement("tr")

        for (var prop in element) {
            var td = document.createElement('td');
            td.innerHTML = prop + ": " + element[prop];
            row.appendChild(td);
        }
        table.appendChild(row);
    });
    document.querySelector("#processes").innerHTML = "";
    document.querySelector("#processes").appendChild(table)
}

function renderCpuUsage(evt) {
    var parsed = JSON.parse(evt.data)

    var table = document.createElement("table")

    for (let i = 0; i < parsed.length; i++) {
        const usage = parsed[i].toFixed(2);
        var row = document.createElement("tr")
        var td = document.createElement('td');

        td.innerHTML = i + ": " + renderUsage(parsed[i].toFixed(0)) + "    " + usage;
        row.appendChild(td);
        table.appendChild(row);
    }

    document.querySelector("#cpuusage").innerHTML = "";
    document.querySelector("#cpuusage").appendChild(table)
}

function renderUsage(usage) {
    var result = ""
    for (let i = 0; i < usage; i++) {
        result += "|"
    }
    return result
}