<template>
  <q-card flat bordered>
    <q-card-section>
      <strong class="text-h6">Cpu usage</strong>
    </q-card-section>

    <q-card-section class="q-pt-none">
        <div v-for="(usage, index) in Usages" v-bind:key="index" style="margin-bottom:2px;">
          <q-linear-progress size="20px" :value="usage/100" color="secondary">
            <div class="absolute-full flex flex-center">
              <q-badge color="white" text-color="accent" :label="usage + ' %'" />
            </div>
          </q-linear-progress>
        </div>
    </q-card-section>
  </q-card>
</template>

<script>
import { monitorService } from "../monitor.service";
export default {
  name: "CpuUsage",
  data() {
    return {
      Usages: []
    };
  },
  mounted() {
    var self = this;
    monitorService.GetCpuUsage(function(evt) {
      var parsed = JSON.parse(evt.data);
      self.Usages = parsed.map(element => {
        return element.toFixed(2);
      });
    });
  }
};
</script>
