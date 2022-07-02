import { Component, OnInit } from "@angular/core";
import { ApiService } from "src/services/api.service";

@Component({
  selector: "app-login",
  templateUrl: "./login.component.html",
})
export class LoginComponent implements OnInit {
  constructor(private api:ApiService) {}
  formData: any
  ngOnInit(): void {}

  SubmitLogin(){
    this.api.post('/auth/login',this.formData).subscribe((data)=>{
      console.log(data)
    })
  }
}
