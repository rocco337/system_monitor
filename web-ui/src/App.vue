<template>
  <div id="app">
    <div class="q-pa-md">
      <div class="row">
        <div class="col-12">
          <HostInfo />
        </div>
      </div>
    </div>
    <div class="q-pa-md">
      <div class="row">
        <div class="col-6">
          <div class="q-pa-md">
            <Temperatures />
          </div>
        </div>
        <div class="col-6">
          <div class="q-pa-md">
            <div class="column">
              <div class="col-auto">
                <CpuUsage />
              </div>
              <div class="col col-md-2">
                <MemoryStat />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Processes -->
    <div class="q-pa-md">
      <div class="row">
        <div class="col-12">
          <Processes />
        </div>
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
