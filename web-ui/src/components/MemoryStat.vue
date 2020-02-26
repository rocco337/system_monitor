<template>
  <div>
      <span v-text="Total"></span>
      <span v-text="Available"></span>
      <span v-text="Used"></span>
      <span v-text="UsedPercent"></span>
      <span v-text="Free"></span>    
  </div>
</template>

<script>
import {monitorService} from "../monitor.service"
export default {
  name: "MemoryStat",
  data() {
    return {
        Total: "",
        Available: "",
        Used: "",
        UsedPercent: "",
        Free: "",
      };
    },
 mounted() {
    var self = this;
    monitorService.GetMemoryStat(function(evt) {
      var parsed = JSON.parse(evt.data);
      self.Total = parsed.Total;
      self.Available = parsed.Available;
      self.Used = parsed.Used;
      self.UsedPercent = parsed.UsedPercent;
      self.Free = parsed.Free;
    });    
  },
  
};
</script>
