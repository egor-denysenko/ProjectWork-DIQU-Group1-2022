import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { ApiService } from './api.service';
const jwtHelper = new JwtHelperService();

@Injectable()
export class AuthService {

  constructor(
    private api: ApiService
  ) {}


  public async isAuthenticated(){
    // Get token from localstorage
    const token = localStorage.getItem('token');
    console.log(token)
    // Check if token is null or valorized
    if (token !== null) {
      console.log("token non verif")
      if (this.verifyToken(token)) {
        return jwtHelper.isTokenExpired(token);
      }
    }
      return false;
  }

  public async login(email, password) {
    if (localStorage === undefined) {
     localStorage.setItem("token","")
    }
    return new Promise((resolve, reject) => {
      this.api.post('/auth', { email: email, password: password }).subscribe(
        async (data) => {
          if (data !== []) {
            localStorage.setItem('token', data.token);
            resolve('OK');
          } else {
            console.log('errore nel login');
            reject('ERRORE ' + data.error);
          }
        },
        (e) => {
          reject(e);
        }
      );
    });
  }

  private verifyToken(token: string): boolean {
      this.api.post('/auth/verify', { token }).subscribe((verified) => {
        console.log(verified)
        if(verified.status){
          return true
        }
        return true
      });
      return false
  }

  public logout() {
    localStorage.remove('token');
  }
}
