import {Component, OnInit} from "@angular/core";
import {Auth} from "../auth/auth.service";
import {Router} from "@angular/router";

@Component({
  selector: "kubermatic-frontpage",
  templateUrl: "./frontpage.component.html",
  styleUrls: ["./frontpage.component.scss"]
})
export class FrontpageComponent implements OnInit {

  constructor(private auth: Auth, private router: Router) {
  }

  ngOnInit(): void {
    if (this.auth.authenticated()) {
      this.router.navigate(["dashboard/clusters"]);
    } else {
      localStorage.setItem("redirect_url", "dashboard/clusters");
      this.auth.login();
    }
  }
}
