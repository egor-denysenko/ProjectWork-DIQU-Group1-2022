import { BrowserModule } from "@angular/platform-browser";
import { NgModule } from "@angular/core";

import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";
import { LoginComponent } from "./views/login/login.component";
import { RegisterComponent } from "./views/register/register.component";
import { AuthComponent } from './layouts/auth/auth.component';
import { FooterSmallComponent } from "./components/footer-small/footer-small.component";
import { AuthNavbarComponent } from "./components/auth-navbar/auth-navbar.component";

@NgModule({
  declarations: [LoginComponent,RegisterComponent,AppComponent, AuthComponent,FooterSmallComponent,AuthNavbarComponent],
  imports: [BrowserModule, AppRoutingModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
