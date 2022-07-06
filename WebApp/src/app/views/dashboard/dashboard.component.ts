import { Component, OnInit } from "@angular/core";
import { ApiService } from "src/services/api.service";

@Component({
  selector: "app-dashboard",
  templateUrl: "./dashboard.component.html",
})
export class DashboardComponent implements OnInit {
  constructor(private api:ApiService) {}

  ngOnInit() {
    console.log("init dati")
  }
}
