import { Component, OnInit, AfterViewInit } from "@angular/core";
import Chart from "chart.js";
import { ApiService } from "src/services/api.service";

interface Data {
  timestamp: string;
  value: string;
  item: number;
}
@Component({
  selector: "humidity-line-chart",
  templateUrl: "./humidity-line-chart.component.html",
})
export class HumidityLineChartComponent implements OnInit {
  dato: Array<Data>;
  constructor(private api: ApiService) {}

  ngOnInit(): void {}
  ngAfterViewInit() {
    this.api.get("/dato").subscribe((data) => {
      let newArray = data.filter((item) => {
        return item._field === "Humidity";
      });

      this.dato = newArray.map((i) => {
        return { timestamp: i._time, item: i._field, value: i._value };
      });
      console.log(this.dato);
      this.buildGraph();
    });
  }
  buildGraph() {
    var config = {
      type: "line",
      data: {
        labels: this.dato.map((i) => {
          return i.timestamp;
        }),
        datasets: [
          {
            label: new Date().getFullYear(),
            backgroundColor: "#4c51bf",
            borderColor: "#4c51bf",
            data: this.dato.map((i) => {
              return i.value;
            }),
            fill: false,
          },
        ],
      },
      options: {
        maintainAspectRatio: false,
        responsive: true,
        title: {
          display: false,
          text: "Sales Charts",
          fontColor: "white",
        },
        legend: {
          labels: {
            fontColor: "white",
          },
          align: "end",
          position: "bottom",
        },
        tooltips: {
          mode: "index",
          intersect: false,
        },
        hover: {
          mode: "nearest",
          intersect: true,
        },
        scales: {
          xAxes: [
            {
              ticks: {
                fontColor: "rgba(255,255,255,.7)",
              },
              display: true,
              scaleLabel: {
                display: false,
                labelString: "Month",
                fontColor: "white",
              },
              gridLines: {
                display: false,
                borderDash: [2],
                borderDashOffset: [2],
                color: "rgba(33, 37, 41, 0.3)",
                zeroLineColor: "rgba(0, 0, 0, 0)",
                zeroLineBorderDash: [2],
                zeroLineBorderDashOffset: [2],
              },
            },
          ],
          yAxes: [
            {
              ticks: {
                fontColor: "rgba(255,255,255,.7)",
              },
              display: true,
              scaleLabel: {
                display: false,
                labelString: "Value",
                fontColor: "white",
              },
              gridLines: {
                borderDash: [3],
                borderDashOffset: [3],
                drawBorder: false,
                color: "rgba(255, 255, 255, 0.15)",
                zeroLineColor: "rgba(33, 37, 41, 0)",
                zeroLineBorderDash: [2],
                zeroLineBorderDashOffset: [2],
              },
            },
          ],
        },
      },
    };
    let ctx: any = document.getElementById("line-chart") as HTMLCanvasElement;
    ctx = ctx.getContext("2d");
    new Chart(ctx, config);
  }
}
