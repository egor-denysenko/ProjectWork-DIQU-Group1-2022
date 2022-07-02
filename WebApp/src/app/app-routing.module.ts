import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { AuthComponent } from "./layouts/auth/auth.component";
import { LoginComponent } from "./views/login/login.component";
import { RegisterComponent } from "./views/register/register.component";


const routes: Routes = [
  /*{
    path:'dashboard',
    component: ,
    children: [
      { path: "Dashboard", component: },
      { path: "TraindData", component:},
      { path: "tables", component: },
      { path: "", redirectTo: "dashboard", pathMatch: "full" },
    ],
  }
  // admin views
  {
    path: "admin",
    component: ,
    children: [
      { path: "TrainCommand", component:  },
    ],
  },*/
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
