<template>
  <div>
    <apexchart
      :height="height"
      type="line"
      :options="options"
      :series="series"
    ></apexchart>
  </div>
</template>

<script>
import defaultSeries from "@/static/defaultSeries.json";
import defaultOptions from "@/static/defaultOptions.json";

export default {
  name: "LineChart",

  props: {
    height: {
      type: String,
      default: "200",
    },

    series: {
      type: Array,
      default: () => defaultSeries,
    },

    xaxis: {
      type: Array,
      default: () => [],
    },
  },

  computed: {
    options() {
      if (this.xaxis.length) {
        return {
          chart: defaultOptions.chart,
          xaxis: {
            categories: this.xaxis,
          },
          yaxis: {
            labels: {
              formatter: (val) => {
                return val;
              },
            },
          },
        };
      } else {
        return defaultOptions;
      }
    },
  },
};
</script>
