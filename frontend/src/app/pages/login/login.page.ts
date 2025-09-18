import { Component, inject } from '@angular/core';

import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';

import { AuthService } from '@cubeshares/shared/services/auth';

@Component({
  selector: 'cubeshares-login-page',
  templateUrl: 'login.page.html',
  styleUrl: 'login.page.scss',
  imports: [ButtonModule, CardModule],
})
export class LoginPageComponent {
  private readonly auth = inject(AuthService);

  protected onLoginViaWCA(): void {
    this.auth.wcaOAuthLogin();
  }
}
