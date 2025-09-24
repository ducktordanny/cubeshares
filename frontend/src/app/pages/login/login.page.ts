import { Component, computed, inject, signal } from '@angular/core';

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
  protected readonly isRedirecting = signal<boolean>(false);
  protected readonly buttonState = computed(() =>
    this.isRedirecting() ? 'Redirecting...' : 'Continue with WCA',
  );

  private readonly auth = inject(AuthService);

  protected onLoginViaWCA(): void {
    this.isRedirecting.set(true);
    this.auth.wcaOAuthLogin();
  }
}
