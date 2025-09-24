import { Component } from '@angular/core';

import { MenuItem } from 'primeng/api';
import { Avatar } from 'primeng/avatar';
import { Menu } from 'primeng/menu';
import { Toolbar } from 'primeng/toolbar';

import { AuthService } from '@cubeshares/services/auth';
import { UserService } from '@cubeshares/services/user';

@Component({
  selector: 'cubeshares-header',
  templateUrl: 'header.component.html',
  styleUrl: 'header.component.scss',
  imports: [Avatar, Menu, Toolbar],
})
export class HeaderComponent {
  protected readonly user = this.userService.loggedInUser;
  protected readonly userMenuItems: MenuItem[] = [
    {
      label: 'Sign out',
      icon: 'pi pi-sign-out',
      command: () => this.authService.logout(),
    },
  ];

  constructor(
    private readonly userService: UserService,
    private readonly authService: AuthService,
  ) {}
}
