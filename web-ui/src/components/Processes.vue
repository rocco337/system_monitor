<template>
  <div>
    <ul>
       <li v-for="(process,index) in Processes" v-bind:key="process.Pid" >
       {{index}} {{ process.Pid }}{{ process.Name }}{{ process.Status }}{{ process.CPUPercent }}{{ process.MemPercent }}{{ process.Username }}{{ process.Uptime }}
      </li>
    </ul>
  </div>
</template>

<script>
import {monitorService} from "../monitor.service"
export default {
  name: "Processes",
  data() {
    return {
        Processes:[]
      };
    },
 mounted() {
    var self = this;
    monitorService.GetListOfProcesses(function(evt) {
      var parsed = JSON.parse(evt.data);
      self.Processes = parsed;
    });
    
  },
  
};



</script>
