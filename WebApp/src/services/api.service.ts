import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private root:string;
  private port:string;
  constructor(private http: HttpClient) {
    this.root = 'http://localhost';
    this.port = ':5555';
  }
}
