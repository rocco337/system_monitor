<template>
  <q-card flat bordered>
    <q-card-section>
      <strong class="text-h6">Temperatures</strong>
      <q-toggle v-model="IsEnabled" />
    </q-card-section>

    <q-card-section class="q-pt-none">
      <temperatures-graph :chart-data="datacollection" />
    </q-card-section>
  </q-card>
</template>

<script>
import { monitorService } from "../monitor.service";
import TemperaturesGraph from "./TemperaturesGraph";

const maxGraphLength = 120;

export default {
  name: "Temperatures",
  components: {
    TemperaturesGraph
  },
  data() {
    return {
      IsEnabled: true,
      datacollection: {},
      Temperatures: null,
      elapsed: [],
      elapsedGrouped: []
    };
  },
  mounted() {
    this.loadData();
  },
  watch: {
    IsEnabled: function(newVal) {
      if (!newVal) {
        this.Processes = [];
      } else {
        this.loadData();
      }
    }
  },
  methods: {
    loadData() {
      var self = this;
      const startedReceivingData = new Date().getTime();
      monitorService.GetTemperatures(function(evt, conn) {
        if (!self.IsEnabled) {
          self.datacollection = {
            labels: [],
            datasets: []
          };
          conn.close();
          return;
        }
        var parsed = JSON.parse(evt.data);
        var datasets = [];
        if (self.Temperatures == null) {
          self.Temperatures = {};
          let i = 0;
          for (var key in parsed) {
            self.Temperatures[key] = {
              label: key,
              borderColor: colors[i++],
              fill: false,
              pointRadius: 0,
              data: [parsed[key]]
            };
          }
        } else {
          for (let parsedPropKey in parsed) {
            for (let propKey in self.Temperatures) {
              if (parsedPropKey === propKey) {
                let prop = self.Temperatures[propKey];

                pushLimitedArray(
                  maxGraphLength,
                  prop.data,
                  parsed[parsedPropKey]
                );

                datasets.push(prop);
              }
            }
          }
        }

        pushLimitedArray(
          maxGraphLength,
          self.elapsed,
          timeSince(startedReceivingData)
        );
        self.elapsedGrouped = groupElapsedTimes(self.elapsed, 10);

        self.datacollection = {
          labels: self.elapsedGrouped,
          datasets: datasets
        };
      });
    }
  }
};

function groupElapsedTimes(arrayToGroup, maxAllowedLabels) {
  let grouped = [];

  let groupRate = Math.round(arrayToGroup.length / maxAllowedLabels);

  let lastPushedIndex = 0;
  for (let ii = 0; ii < arrayToGroup.length; ii++) {
    if (ii == 0 || ii - lastPushedIndex >= groupRate) {
      grouped.push(arrayToGroup[ii]);
      lastPushedIndex = ii;
    } else {
      grouped.push("");
    }
  }

  return grouped;
}

function pushLimitedArray(limit, array, value) {
  if (array.length >= limit) {
    array.shift();
  }
  array.push(value);
}

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

const colors = [
  "aqua",
  "lime",
  "silver",
  "black",
  "maroon",
  "teal",
  "blue",
  "navy",
  "fuchsia",
  "olive",
  "yellow",
  "gray",
  "purple",
  "green",
  "red"
];
</script>
