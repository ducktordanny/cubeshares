import { Component } from '@angular/core';

import { ButtonModule } from 'primeng/button';

import { UserService } from '@cubeshares/services/user';

import { UserDetailsCardComponent } from '../components/user-details-card/user-details-card.component';

@Component({
  selector: 'cubeshares-me-page',
  templateUrl: 'me.page.html',
  imports: [ButtonModule, UserDetailsCardComponent],
})
export class MePageComponent {
  protected readonly user = this.userService.loggedInUser;

  constructor(private readonly userService: UserService) { }
}
