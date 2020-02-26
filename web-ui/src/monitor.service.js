const baseUrl = "ws://localhost:8081";

export const monitorService = {

    GetHostInfo(onMessage){
        this.listen(baseUrl + "/hostInfo",onMessage)
    },
    GetMemoryStat(onMessage){
        this.listen(baseUrl + "/memorystat",onMessage)
    },
    GetListOfProcesses(onMessage){
        this.listen(baseUrl +"/processes", onMessage)
    },
    GetCpuUsage(onMessage){
      this.listen(baseUrl +"/cpuusage", onMessage)
  },
    listen(url, onmessage) {
        var conn = new WebSocket(url);
        conn.onclose = function(evt) {
          console.log(evt);
          // connectionClosed(evt);
        };
        conn.onmessage = function(evt) {
          onmessage(evt);
        };
      }
};