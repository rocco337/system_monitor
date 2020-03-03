function getUrl(){
  var url =  "ws://" + location.host;
  console.log(url);
  return url;
}

export const monitorService = {

    GetHostInfo(onMessage){
        this.listen(getUrl() + "/hostInfo",onMessage)
    },
    GetMemoryStat(onMessage){
        this.listen(getUrl() + "/memorystat",onMessage)
    },
    GetListOfProcesses(onMessage){
        this.listen(getUrl() +"/processes", onMessage)
    },
    GetCpuUsage(onMessage){
      this.listen(getUrl() +"/cpuusage", onMessage)
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