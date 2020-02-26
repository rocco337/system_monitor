const baseUrl = "ws://localhost:8081";

export const monitorService = {

    GetHostInfo(onMessage){
        this.listen(baseUrl + "/hostInfo",onMessage)
    },
    listen(url, onmessage) {
        var conn = new WebSocket(url);
        conn.onclose = function(evt) {
          console.log(evt);
          // connectionClosed(evt);
        };
        conn.onmessage = function(evt) {
          console.log("received: ", evt);
          onmessage(evt);
        };
      }
};