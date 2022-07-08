import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
@Injectable({
  providedIn: "root",
})
export class ApiService {
  private root: string;
  private port: string;
  constructor(private http: HttpClient) {
    this.root = "http://20.23.39.128";
    this.port = ":5555";
  }

  public get(link: string) {
    return this.http.get<any>(this.root + this.port + link);
  }

  public post(link, obj) {
    return this.http.post<any>(this.root + this.port + link, obj);
  }
}
