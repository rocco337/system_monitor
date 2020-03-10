<template>
  <q-card flat bordered>
    <q-card-section>
      <strong class="text-h6">Processes</strong>
      <q-toggle v-model="IsEnabled" />
    </q-card-section>

    <q-card-section class="q-pt-none">
      <q-table
        style="height: 400px"
        :data="Processes"
        :columns="Columns"
        row-key="Pid"
        hide-bottom
        :pagination.sync="pagination"
        :rows-per-page-options="[0]"
        :dense="$q.screen.lt.md"
      />
    </q-card-section>
  </q-card>
</template>

<script>
import { monitorService } from "../monitor.service";
export default {
  name: "Processes",
  data() {
    return {
      Processes: [],
      IsEnabled: true,
      pagination: {
        page: 1,
        rowsPerPage: 20
      },
      Columns: [
        { name: "Pid", label: "Pid", field: "Pid" },
        { name: "Name", label: "Name", field: "Name" },
        { name: "CPUPercent", label: "CPU", field: "CPUPercent" },
        { name: "MemPercent", label: "Mem", field: "MemPercent" },
        { name: "Username", label: "Username", field: "Username" },
        { name: "Uptime", label: "Uptime", field: "Uptime" }
      ]
    };
  },
  mounted() {
   this.loadData();
  },
  watch: {
    IsEnabled: function(newVal) {
      if(!newVal){
        this.Processes = [];
      }else{
        this.loadData();
      }
    }
  },
  methods: {
    loadData() {
    var self = this;

       monitorService.GetListOfProcesses(function(evt, conn) {
      if(!self.IsEnabled){
        self.Processes = [];
        conn.close();
        return;
      }
      var parsed = JSON.parse(evt.data);
      self.Processes = parsed;
    });
    }
  }
};
</script>
