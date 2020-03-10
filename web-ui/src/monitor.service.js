function getUrl(){
  // var url =  "ws://localhost:8081"
  var url =  "ws://" + location.host;
  return url;
}

export const monitorService = {

    GetHostInfo(onMessage){
        this.listen(getUrl() + "/hostInfo",onMessage)
    },
    GetTemperatures(onMessage){
      this.listen(getUrl() + "/temperatures",onMessage)
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
          console.log('close:', evt, new Date());
        };
        
        conn.onmessage = function(evt) {
          onmessage(evt, conn);
        };

        conn.onerror = function(evt){
          console.log("web socket error", evt)
        }
      }
};