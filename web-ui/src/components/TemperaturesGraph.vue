

<script>
import { Line, mixins } from 'vue-chartjs'

const { reactiveProp } = mixins

export default {
  name: 'TemperaturesGraph',
  extends: Line,
  mixins: [reactiveProp],
  data () {
          return {
            options: { 
              animation: false,
              scales: {
                yAxes: [{
                  ticks: {
                    beginAtZero: true
                  },
                  gridLines: {
                    display: true
                  }
                }],
                xAxes: [ {
                  gridLines: {
                    display: false
                  }
                }]
              },
              legend: {
                display: true
              },
              responsive: true,
              maintainAspectRatio: false
            }
          }
        },
  mounted() {
    this.renderChart(this.chartData, this.options)
  }
};

</script>
