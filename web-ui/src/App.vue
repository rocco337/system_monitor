<template>
  <div id="app">
    <HostInfo/>
    <Temperatures/>
    <MemoryStat/>
    <CpuUsage/>
    <Processes/>
    
  </div>
</template>

<script>
import HostInfo from './components/HostInfo.vue'
import MemoryStat from './components/MemoryStat.vue'
import Processes from './components/Processes.vue'
import CpuUsage from './components/CpuUsage.vue'
import Temperatures from './components/Temperatures.vue'

export default {
  name: 'App',

  components: {
     HostInfo,
    MemoryStat,
    Processes,
    CpuUsage,
    Temperatures,
  },  
}
</script>

<style>
</style>
