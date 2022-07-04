import { Component, OnInit } from "@angular/core";
import { ApiService } from "src/services/api.service";
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
@Component({
  selector: "app-login",
  templateUrl: "./login.component.html",
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup
  constructor(private api:ApiService,private formBuilder: FormBuilder) {
    this.loginForm = this.formBuilder.group({
      email: [null,[Validators.required,Validators.email,Validators.minLength(8)]],
      password: [null, [Validators.required]],
    });
  }
  formData: any
  ngOnInit(): void {}

  SubmitLogin(){
    this.api.post('/auth/login',this.loginForm.value).subscribe((data)=>{
      console.log(data)
    })
  }
}
