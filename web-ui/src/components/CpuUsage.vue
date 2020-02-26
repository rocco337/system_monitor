<template>
  <div>
    <ul>
       <li v-for="(usage,index) in Usages" v-bind:key="usage" >
       {{index}} {{ usage }}
      </li>
    </ul>
  </div>
</template>

<script>
import {monitorService} from "../monitor.service"
export default {
  name: "CpuUsage",
  data() {
    return {
        Usages:[]
      };
    },
 mounted() {
    var self = this;
    monitorService.GetCpuUsage(function(evt) {
      var parsed = JSON.parse(evt.data);
      self.Usages = parsed;
    });
    
  },
  
};



</script>
