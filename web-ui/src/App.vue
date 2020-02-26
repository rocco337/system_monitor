<template>
  <div id="app">
    <HostInfo/>
    <MemoryStat/>

  </div>
</template>

<script>
import HostInfo from './components/HostInfo.vue'
import MemoryStat from './components/MemoryStat.vue'


export default {
  name: 'App',
  components: {
    HostInfo,
    MemoryStat
  },  
}


</script>

