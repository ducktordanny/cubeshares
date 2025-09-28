import { Component, computed } from '@angular/core';
import { Router } from '@angular/router';

import { Avatar } from 'primeng/avatar';
import { ButtonModule } from 'primeng/button';
import { Menu } from 'primeng/menu';
import { MenuItem } from 'primeng/api';
import { Toolbar } from 'primeng/toolbar';

import { AuthService } from '@cubeshares/services/auth';
import { UserMeService } from '@cubeshares/services/user';
import { urlPath } from '@cubeshares/signals';

@Component({
  selector: 'cubeshares-header',
  templateUrl: 'header.component.html',
  styleUrl: 'header.component.scss',
  imports: [Avatar, ButtonModule, Menu, Toolbar],
})
export class HeaderComponent {
  protected readonly user = this.userMeService.loggedInUser;
  protected readonly isLoading = this.userMeService.isLoading;
  protected readonly userMenuItems = computed<MenuItem[]>(() => [
    {
      label: 'My profile',
      icon: 'pi pi-user',
      disabled: this.url().startsWith('/user/me'),
      command: () => this.router.navigate(['/user/me'])
    },
    {
      label: 'Sign out',
      icon: 'pi pi-sign-out',
      command: () => this.authService.logout(),
    },
  ]);
  private readonly url = urlPath();

  constructor(
    private readonly userMeService: UserMeService,
    private readonly authService: AuthService,
    private readonly router: Router,
  ) { }

  protected onSignUp(): void {
    void this.router.navigate(['/login']);
  }
}
