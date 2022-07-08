import { Component, OnInit, Input } from "@angular/core";
import { Router } from "@angular/router";

@Component({
  selector: "app-card-table",
  templateUrl: "./card-table.component.html",
})
export class CardTableComponent implements OnInit {
  @Input()
  get color(): string {
    return this._color;
  }
  set color(color: string) {
    this._color = color !== "light" && color !== "dark" ? "light" : color;
  }
  private _color = "light";
  list = [
    { TrainID: 1, TransportType: "Civil", TrainStatus: "In Transit" },
    { TrainID: 2, TransportType: "Cargo", TrainStatus: "Stopped" },
    { TrainID: 3, TransportType: "Cargo", TrainStatus: "In Transit" },
    { TrainID: 4, TransportType: "Civil", TrainStatus: "In Transit" },
    { TrainID: 5, TransportType: "Civil", TrainStatus: "In Transit" },
    { TrainID: 6, TransportType: "Civil", TrainStatus: "In Transit" },
    { TrainID: 7, TransportType: "Cargo", TrainStatus: "Stopped" },
  ];
  constructor(private router: Router) {}

  ngOnInit(): void {}
  viewTrainData(trainID: number) {
    console.log("ecco dati");
    this.router.navigate(["TrainData", trainID]);
  }
}
