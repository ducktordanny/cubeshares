import { Component } from '@angular/core';

import { ButtonModule } from 'primeng/button';

import { UserService } from '@cubeshares/services/user';

@Component({
  selector: 'cubeshares-me-page',
  templateUrl: 'me.page.html',
  styleUrl: 'me.page.scss',
  imports: [ButtonModule],
})
export class MePageComponent {
  protected readonly user = this.userService.loggedInUser;

  constructor(private readonly userService: UserService) {}
}
