import { Component, OnInit } from "@angular/core";

import { ApiService } from "src/services/api.service";
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
@Component({
  selector: "app-register",
  templateUrl: "./register.component.html",
})
export class RegisterComponent implements OnInit {
  registrationForm: FormGroup
  constructor(private formBuilder:FormBuilder,private api:ApiService) {
    this.registrationForm = this.formBuilder.group({
      username: [null,[Validators.required,Validators.maxLength(18),Validators.minLength(2)]],
      email: [undefined,[Validators.required,Validators.email,Validators.minLength(5)]],
      password: [undefined, [Validators.required]],
    });
  }

  ngOnInit(): void {}

  async SubmitRegistration(){
    console.log(this.registrationForm)
    this.api.post('/auth/register',this.registrationForm.value).subscribe(async (data:Object)=>{
      console.log(data)
    })
  }
}
