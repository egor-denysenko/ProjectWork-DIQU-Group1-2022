import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { AuthGuard } from "src/services/auth.guard";
import { AuthComponent } from "./layouts/auth/auth.component";
import { DashboardComponent } from "./views/dashboard/dashboard.component";
import { LoginComponent } from "./views/login/login.component";
import { RegisterComponent } from "./views/register/register.component";
import { TrainListComponent } from "./views/train-list/train-list.component";


const routes: Routes = [
  {
    path:'dashboard',
    component: TrainListComponent,
  },
  {path:'TrainData/:id',
component:DashboardComponent,
},
  // auth views
  {
    path: "auth",
    component: AuthComponent,
    children: [
      { path: "login", component:LoginComponent},
      { path: "register", component:RegisterComponent},
      { path: "", redirectTo: "login", pathMatch: "full" },
      { path: "**", redirectTo: "", pathMatch: "full" },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})

export class AppRoutingModule {}
