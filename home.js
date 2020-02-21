"use strict"

window.onload = function () {
    listenHostInfo();
    listenProcesses();
    
};

function listenHostInfo(){
    var conn;
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/hostInfo");
        conn.onclose = function(evt){connectionClosed(evt)}
        conn.onmessage = function (evt) {
            console.log(evt.data)
            document.querySelector("#hostInfo").innerHTML = evt.data;
        };
    } else {       
        webSocketsNotAvailable()
    }
}

function listenProcesses(){
    var conn;
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/processes");
        conn.onclose = function(evt){connectionClosed(evt)}
        conn.onmessage = function (evt) {
            
            parseProcesses(evt.data)

        };
    } else {       
        webSocketsNotAvailable()
    }
}

function webSocketsNotAvailable(){
    alert("Your browser does not support WebSockets.");
}

function connectionClosed(evt){
    console.log(evt);
    //alert("Connection closed!")
}

function parseProcesses(json){
    var parsed = JSON.parse(json)
    
    var table = document.createElement("table")
    parsed.forEach(element => {
        var row = document.createElement("tr")

        for(var prop in element){
            var td = document.createElement('td');
            td.innerHTML = prop + ": " + element[prop];
            row.appendChild(td);
        }
        table.appendChild(row);
    });
    document.querySelector("#processes").innerHTML = "";
    document.querySelector("#processes").appendChild(table)
}