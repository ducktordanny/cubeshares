import { Component, effect, untracked } from '@angular/core';
import { Router } from '@angular/router';

import { ButtonModule } from 'primeng/button';

import { UserMeService } from '@cubeshares/services/user';

import { UserDetailsCardComponent } from '../components/user-details-card/user-details-card.component';

@Component({
  selector: 'cubeshares-me-page',
  templateUrl: 'me.page.html',
  imports: [ButtonModule, UserDetailsCardComponent],
})
export class MePageComponent {
  protected readonly user = this.userMeService.loggedInUser;

  constructor(private readonly userMeService: UserMeService, private readonly router: Router) {
    effect(() => {
      const user = this.user();
      untracked(() => !user && void this.router.navigate(['/login']))
    })
  }
}
