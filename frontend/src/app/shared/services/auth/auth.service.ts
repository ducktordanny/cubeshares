import { inject, Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { HttpResponse } from '@angular/common/http';

import { catchError, of, finalize, take, tap } from 'rxjs';

import { api } from '@cubeshares/shared/utils';

import { ApiService } from '../api';
import { UserService } from '../user';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private readonly router = inject(Router);
  private readonly api = inject(ApiService);
  private readonly userService = inject(UserService);

  wcaOAuthLogin(): void {
    window.location.href = api('auth/login');
  }

  logout(): void {
    this.api
      .create<HttpResponse<unknown>>('auth/logout', null)
      .pipe(
        take(1),
        tap(() => this.userService.loggedInUser.set(null)),
        catchError(() => of(null)),
        finalize(() => {
          this.userService.loggedInUser.set(null);
          this.router.navigate(['/login']);
        }),
      )
      .subscribe();
  }
}
