<template>
  <q-card flat bordered>
    <q-card-section>
      <strong class="text-h6">Cpu usage</strong>
      <q-toggle v-model="IsEnabled" />
    </q-card-section>

    <q-card-section class="q-pt-none">
      <div v-for="(usage, index) in Usages" v-bind:key="index" style="margin-bottom:2px;">
        <q-linear-progress size="20px" :value="usage/100" color="info">
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
      IsEnabled: true,
      Usages: []
    };
  },
  mounted() {
    this.loadData();
  },
  watch: {
    IsEnabled: function(newVal) {
      if (!newVal) {
        this.Usages = [];
      }else{
        this.loadData();
      }
    }
  },
  methods: {
    loadData() {
      var self = this;
      monitorService.GetCpuUsage(function(evt, conn) {
        if (!self.IsEnabled) {
          self.Usages = [];
          conn.close();
          return;
        }

        var parsed = JSON.parse(evt.data);
        self.Usages = parsed.map(element => {
          return element.toFixed(2);
        });
      });
    }
  }
};
</script>
