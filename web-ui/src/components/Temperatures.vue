<template>
  <q-card flat bordered>
    <q-card-section>
      <strong class="text-h6">Temperatures</strong>
    </q-card-section>

    <q-card-section class="q-pt-none">
      <temperatures-graph :chart-data="datacollection" />
    </q-card-section>
  </q-card>
</template>

<script>
import { monitorService } from "../monitor.service";
import TemperaturesGraph from "./TemperaturesGraph";

export default {
  name: "Temperatures",
  components: {
    TemperaturesGraph
  },
  data() {
    return {
      datacollection: {},
      Temperatures: null,
      elapsed: []
    };
  },
  mounted() {
    var self = this;
    const currentTimeStamp = new Date().getTime();
    monitorService.GetTemperatures(function(evt) {
      var parsed = JSON.parse(evt.data);
console.log('received');
      var datasets = [];
      if (self.Temperatures == null) {
        self.Temperatures={};
        let i=0;
        for (var key in parsed) {
          self.Temperatures[key] = {
                label: key,
                borderColor: colors[i++],
                 fill: false,
                 pointRadius: 0,
                data: [parsed[key]]
              }              
        }
      } else {
        for (let parsedPropKey in parsed) {
          for (let propKey in self.Temperatures) {
            if (parsedPropKey === propKey) {
              let prop = self.Temperatures[propKey];

              prop.data.push(parsed[parsedPropKey]);

              datasets.push(prop);
            }
          }
        }
      }

      self.elapsed.push(timeSince(currentTimeStamp));

      self.datacollection = {
        labels: self.elapsed,
        datasets: datasets
      };
    });
  }
};

function timeSince(timeStamp) {
  var now = new Date(),
    secondsPast = (now.getTime() - timeStamp) / 1000;
  if (secondsPast < 60) {
    return parseInt(secondsPast) + "s";
  }
  if (secondsPast < 3600) {
    return parseInt(secondsPast / 60) + "m";
  }
  if (secondsPast <= 86400) {
    return parseInt(secondsPast / 3600) + "h";
  }
  if (secondsPast > 86400) {
    let day = timeStamp.getDate();
    let month = timeStamp
      .toDateString()
      .match(/ [a-zA-Z]*/)[0]
      .replace(" ", "");
    let year =
      timeStamp.getFullYear() == now.getFullYear()
        ? ""
        : " " + timeStamp.getFullYear();
    return day + " " + month + year;
  }
}

const colors = ["aqua"  	
,"lime"  	
,"silver"  
,"black"  	
,"maroon"  	
,"teal"  
,"blue"  	
,"navy"  	
,"fuchsia"  	
,"olive"  	
,"yellow"  
,"gray"  	
,"purple"  
,"green"  	
,"red"];
</script>
