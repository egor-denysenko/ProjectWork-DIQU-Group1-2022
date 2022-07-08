import { BrowserModule } from "@angular/platform-browser";
import { NgModule } from "@angular/core";

import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";
import { LoginComponent } from "./views/login/login.component";
import { RegisterComponent } from "./views/register/register.component";
import { AuthComponent } from "./layouts/auth/auth.component";
import { FooterSmallComponent } from "./components/footer-small/footer-small.component";
import { AuthNavbarComponent } from "./components/auth-navbar/auth-navbar.component";
import { HttpClientModule } from "@angular/common/http";

import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { TrainListComponent } from "./views/train-list/train-list.component";
import { CardTableComponent } from "./components/card-table/card-table.component";
import { AuthService } from "src/services/auth.service";
import { TemperatureLineChartComponent } from "./components/temperature-line-chart/temperature-card-line-chart.component";
import { HumidityLineChartComponent } from "./components/humidity-line-chart/humidity-line-chart.component";
import { DashboardComponent } from "./views/dashboard/dashboard.component";

@NgModule({
  declarations: [
    LoginComponent,
    RegisterComponent,
    AppComponent,
    AuthComponent,
    FooterSmallComponent,
    AuthNavbarComponent,
    TrainListComponent,
    CardTableComponent,
    TemperatureLineChartComponent,
    HumidityLineChartComponent,
    DashboardComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  providers: [AuthService],
  bootstrap: [AppComponent],
})
export class AppModule {}
