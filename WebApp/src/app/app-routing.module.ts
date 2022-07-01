import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";


const routes: Routes = [
  {
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
  },
  // auth views
  {
    path: "auth",
    component: ,
    children: [
      { path: "login", component:},
      { path: "register", component:},
      { path: "", redirectTo: "login", pathMatch: "full" },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
