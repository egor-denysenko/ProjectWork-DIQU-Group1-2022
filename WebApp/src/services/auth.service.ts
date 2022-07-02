import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { ApiService } from './api.service';
const jwtHelper = new JwtHelperService();

@Injectable()
export class AuthService {
  private localStorage: Storage; //| null = null;
  constructor(
    private api: ApiService,
    private storage: Storage,
  ) {}

  async onInit() {
    // If using, define drivers here: await this.storage.defineDriver(/*...*/);
    const storage = await this.storage.create();
    this.localStorage = storage;
  }

  public async isAuthenticated(): Promise<any> {
    await this.onInit();
    // Get token from localstorage
    const token = await this.localStorage.get('token');
    // Check if token is null or empty
    if (token !== null) {
      // Check whether the token is expired and return
      // true or false
      const verified = await this.verifyToken(token);
      console.log(verified);
      if (verified) {
        const { id, username } = jwtHelper.decodeToken(token);
        return !jwtHelper.isTokenExpired(token);
      } else {
        return false;
      }
    } else {
      return false;
    }
  }

  public async login(email, password) {
    if (this.localStorage === undefined) {
      await this.onInit();
    }
    return new Promise((resolve, reject) => {
      this.api.post('/login', { email: email, password: password }).subscribe(
        async (data) => {
          if (data !== []) {
            console.log(data.token);
            await this.localStorage.set('token', data.token);
            resolve('OK');
          } else {
            console.log('errore nel login');
            console.log(data.error);
            reject('ERRORE ' + data.error);
          }
        },
        (e) => {
          reject(e);
        }
      );
    });
  }

  private async verifyToken(token) {
    return new Promise((resolve, reject) => {
      this.api.post('/login/verify', { token }).subscribe((verified) => {
        if (verified.status) resolve(true);
        else reject(false + verified.error);
      });
    });
  }

  public logout() {
    this.localStorage.remove('token');
  }
}
