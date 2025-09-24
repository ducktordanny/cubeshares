import { Component } from '@angular/core';
import { JsonPipe } from '@angular/common';

import { ButtonModule } from 'primeng/button';

import { UserService } from '@cubeshares/services/user';
import { AuthService } from '@cubeshares/services/auth';

@Component({
  selector: 'cubeshares-me-page',
  templateUrl: 'me.page.html',
  imports: [ButtonModule, JsonPipe],
})
export class MePageComponent {
  protected readonly user = this.userService.loggedInUser;

  constructor(
    private readonly authService: AuthService,
    private readonly userService: UserService,
  ) {}

  protected onLogout(): void {
    this.authService.logout();
  }
}
