import { Component, OnInit } from "@angular/core";
import { ApiService } from "src/services/api.service";
import { FormGroup, FormBuilder, Validators } from "@angular/forms";
@Component({
  selector: "app-dashboard",
  templateUrl: "./dashboard.component.html",
})
export class DashboardComponent implements OnInit {
  tempForm: FormGroup;
  lights: boolean;
  constructor(private api: ApiService, private formBuilder: FormBuilder) {
    this.lights = false;
    this.tempForm = this.formBuilder.group({
      tempValue: [
        null,
        [Validators.required, Validators.max(35), Validators.min(10)],
      ],
    });
  }

  ngOnInit() {
    console.log("init dati");
  }

  sendDesideredTemp() {
    let tempCommand = {
      TrainId: 1,
      WagonId: 1,
      WagonCommand: 70,
      Temperature: this.tempForm.value.tempValue,
    };
    console.log("sending data");
    this.api.post("/comando/temperature", tempCommand).subscribe((data) => {
      console.log(data);
      if (data) {
        this.tempForm.reset();
      }
    });
  }

  setLights() {
    let lightCommand = {
      TrainId: 1,
      WagonId: 1,
      WagonCommand: 71,
      LightMode: this.lights,
    };
    console.log("sending data");
    this.api.post("/comando/lights", lightCommand).subscribe((data) => {
      console.log(data);
      if (data) {
        this.lights = !this.lights;
      }
    });
  }
}
