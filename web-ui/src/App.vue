<template>
  <div id="app">
     <div class="row">
      <div class="col-12">
         <HostInfo />
      </div>
    </div>
    <div class="row">
      <div class="col-6">
         <Temperatures />
      </div>
      <div class="col-6">
        <CpuUsage />
        <MemoryStat />
      </div>
    </div>
     <div class="row">
      <div class="col-12">
        <Processes />
      </div>
    </div>
  </div>
</template>

<script>
import HostInfo from "./components/HostInfo.vue";
import MemoryStat from "./components/MemoryStat.vue";
import Processes from "./components/Processes.vue";
import CpuUsage from "./components/CpuUsage.vue";
import Temperatures from "./components/Temperatures.vue";

export default {
  name: "App",

  components: {
    HostInfo,
    MemoryStat,
    Processes,
    CpuUsage,
    Temperatures
  }
};
</script>

<style>
</style>
