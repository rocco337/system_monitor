"use strict"

window.onload = function () {
    listenHostInfo();
    listenProcesses();
    listenCpuUsage();

};

function listenHostInfo() {
    var conn;
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/hostInfo");
        conn.onclose = function (evt) { connectionClosed(evt) }
        conn.onmessage = function (evt) {
            console.log(evt.data)
            document.querySelector("#hostInfo").innerHTML = evt.data;
        };
    } else {
        webSocketsNotAvailable()
    }
}

function listenProcesses() {
    var conn;
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/processes");
        conn.onclose = function (evt) { connectionClosed(evt) }
        conn.onmessage = function (evt) {

            renderProcesses(evt.data)

        };
    } else {
        webSocketsNotAvailable()
    }
}

function listenCpuUsage() {
    var conn;
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/cpuusage");
        conn.onclose = function (evt) { connectionClosed(evt) }
        conn.onmessage = function (evt) {
            renderCpuUsage(evt.data)
        };
    } else {
        webSocketsNotAvailable()
    }
}


function webSocketsNotAvailable() {
    alert("Your browser does not support WebSockets.");
}

function connectionClosed(evt) {
    console.log(evt);
    //alert("Connection closed!")
}

function renderProcesses(json) {
    var parsed = JSON.parse(json)

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

function renderCpuUsage(json) {
    var parsed = JSON.parse(json)

    var table = document.createElement("table")

    for (let i = 0; i < parsed.length; i++) {
        const usage = parsed[i].toFixed(2);
        var row = document.createElement("tr")
        var td = document.createElement('td');

        td.innerHTML = i + ": " + renderUsage(parsed[i].toFixed(0)) + "    " +  usage;
        row.appendChild(td);
        table.appendChild(row);
    }

    document.querySelector("#cpuusage").innerHTML = "";
    document.querySelector("#cpuusage").appendChild(table)
}

function renderUsage(usage){
    var result=""
    for (let i = 0; i < usage; i++) {
        result+="|"
    }
    return result
}