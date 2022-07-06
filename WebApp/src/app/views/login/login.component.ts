import { Component, OnInit } from "@angular/core";
import { ApiService } from "src/services/api.service";
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { AuthService } from "src/services/auth.service";
import { Router } from "@angular/router";
@Component({
  selector: "app-login",
  templateUrl: "./login.component.html",
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup
  constructor(private api:ApiService,private formBuilder: FormBuilder,private auth:AuthService,private router:Router) {
    this.loginForm = this.formBuilder.group({
      email: [null,[Validators.required,Validators.email,Validators.minLength(8)]],
      password: [null, [Validators.required]],
    });
  }
  formData: any
  ngOnInit(): void {
    if (this.checkIfTheUserIsAlreadySignupped) {
      console.log('non torni true?');
      this.router.navigateByUrl('/dashboard/TrainData/1');
    }
  }

  async checkIfTheUserIsAlreadySignupped() {
    if (await this.auth.isAuthenticated()) {
      return true;
    } else {
      return false;
    }
  }

  SubmitLogin(){
    this.auth.login(this.loginForm.get('email').value,this.loginForm.get('password').value)
  }
}
