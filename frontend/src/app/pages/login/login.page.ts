import { Component, inject } from '@angular/core';

import { ButtonModule } from 'primeng/button';

import { AuthService } from '@cubeshares/shared/services/auth';

@Component({
  selector: 'cubeshares-login-page',
  templateUrl: 'login.page.html',
  imports: [ButtonModule],
})
export class LoginPageComponent {
  private readonly auth = inject(AuthService);

  protected onLoginViaWCA(): void {
    this.auth.wcaOAuthLogin();
  }
}
