<template>
  <div>
    <span v-text="Hostname"></span>
    <span v-text="Uptime"></span>
    <span v-text="Procs"></span>
    <span v-text="OS"></span>
  </div>
</template>

<script>
import {monitorService} from "../monitor.service"
export default {
  name: "HostInfo",
  data() {
    return {
        Hostname:"",
        Uptime:"",
        Procs:"",
        OS:""
      };
    },
 mounted() {
    var self = this;
    monitorService.GetHostInfo(function(evt) {
      var parsed = JSON.parse(evt.data);
      self.Hostname = parsed.Hostname;
      self.Uptime = parsed.Uptime;
      self.Procs = parsed.Procs;
      self.OS = parsed.OS;
    });
    
  },
  
};



</script>
