import { inject, Injectable, signal } from '@angular/core';
import { Router } from '@angular/router';

import {
  catchError,
  of,
  Observable,
  shareReplay,
  switchMap,
  take,
  tap,
  timer,
  finalize,
} from 'rxjs';

import { ApiService } from '../api';
import { UserResponse } from './user.type';

@Injectable({ providedIn: 'root' })
export class UserService {
  readonly loggedInUser = signal<UserResponse | null>(null);
  readonly isLoading = signal<boolean>(true);
  private readonly router = inject(Router);
  private readonly api = inject(ApiService);

  constructor() {
    this.pollOnReadUserMe();
  }

  readUserMe(): Observable<UserResponse | null> {
    this.isLoading.set(true);
    return this.api
      .read<UserResponse>('user/me')
      .pipe(take(1))
      .pipe(
        tap((user) => this.loggedInUser.set(user)),
        catchError(() => {
          this.loggedInUser.set(null);
          this.router.navigate(['/login']);
          return of(null);
        }),
        finalize(() => this.isLoading.set(false)),
      );
  }

  private pollOnReadUserMe(): void {
    timer(0, 60 * 1000)
      .pipe(
        switchMap(() => this.readUserMe()),
        shareReplay({ bufferSize: 1, refCount: true }),
      )
      .subscribe();
  }
}
