<template>
  <q-card flat bordered>
    <q-card-section>
      <div class="text-h6" v-text="Hostname"></div>
    </q-card-section>

    <q-card-section class="q-pt-none">
      #Tasks: <strong v-text="Procs"></strong>
      OS: <strong v-text="OS"></strong>
      Uptime: <strong v-text="Uptime"></strong>
    </q-card-section>
  </q-card>
</template>

<script>
import { monitorService } from "../monitor.service";
export default {
  name: "HostInfo",
  data() {
    return {
      Hostname: "",
      Uptime: "",
      Procs: "",
      OS: ""
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
      document.title = parsed.Hostname + " (" + parsed.OS + ")";
    });
  }
};
</script>
