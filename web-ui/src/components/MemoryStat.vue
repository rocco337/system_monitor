<template>
  <q-card flat bordered>
    <q-card-section>
      <strong class="text-h6">Memory</strong>

      <q-toggle v-model="IsEnabled" />
    </q-card-section>

    <q-card-section class="q-pt-none">
      Used:
      <strong class="text-h6" v-text="AvailableFormated"></strong>
      Total
      <strong class="text-h6" v-text="TotalFormated"></strong>
      <q-linear-progress size="25px" :value="UsedPercent" color="accent">
        <div class="absolute-full flex flex-center">
          <q-badge color="white" text-color="accent" :label="UsedPercentLabel" />
        </div>
      </q-linear-progress>
    </q-card-section>
  </q-card>
</template>

<script>
import { monitorService } from "../monitor.service";
export default {
  name: "MemoryStat",
  data() {
    return {
      IsEnabled: true,
      Total: "",
      Available: "",
      Used: "",
      UsedPercent: 0,
      Free: ""
    };
  },
  computed: {
    UsedPercentLabel() {
      return (this.UsedPercent * 100).toFixed(2) + "%";
    },
    AvailableFormated() {
      return bytesToSize(this.Total - this.Available);
    },
    TotalFormated() {
      return bytesToSize(this.Total);
    }
  },
  mounted() {
    this.loadData();
  },
  watch: {
    IsEnabled: function(newVal) {
      if (!newVal) {
        this.Processes = [];
      }else{
        this.loadData();
      }
    }
  },
  methods: {
    loadData() {
      var self = this;
      monitorService.GetMemoryStat(function(evt, conn) {
        if (!self.IsEnabled) {
          self.Total = "";
          self.Available = "";
          self.Used = "";
          self.UsedPercent = 0;
          self.Free = "";
          conn.close();
          return;
        }

        var parsed = JSON.parse(evt.data);
        self.Total = parsed.Total;
        self.Available = parsed.Available;
        self.Used = parsed.Used;
        self.UsedPercent = parsed.UsedPercent / 100;
        self.Free = parsed.Free;
      });
    }
  }
};

function bytesToSize(bytes) {
  var sizes = ["Bytes", "KB", "MB", "GB", "TB"];
  if (bytes == 0) return "0 Byte";
  var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)));
  return Math.round(bytes / Math.pow(1024, i), 2) + " " + sizes[i];
}
</script>
