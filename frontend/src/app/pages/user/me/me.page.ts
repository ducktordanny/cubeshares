import { Component } from '@angular/core';

import { ButtonModule } from 'primeng/button';

import { UserService } from '@cubeshares/shared/services/user';
import { AuthService } from '@cubeshares/shared/services/auth';
import { JsonPipe } from '@angular/common';

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
