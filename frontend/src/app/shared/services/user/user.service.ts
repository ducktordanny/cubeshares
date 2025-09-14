import { inject, Injectable, signal } from '@angular/core';

import { ApiService } from '../api';
import { UserResponse } from './user.type';
import {
  catchError,
  of,
  Observable,
  shareReplay,
  switchMap,
  take,
  tap,
  timer,
} from 'rxjs';

@Injectable({ providedIn: 'root' })
export class UserService {
  readonly loggedInUser = signal<UserResponse | null>(null);
  private readonly api = inject(ApiService);

  constructor() {
    this.pollOnReadUserMe();
  }

  readUserMe(): Observable<UserResponse | null> {
    return this.api
      .read<UserResponse>('user/me')
      .pipe(take(1))
      .pipe(
        tap((user) => this.loggedInUser.set(user)),
        catchError(() => {
          this.loggedInUser.set(null);
          return of(null);
        }),
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
