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
  takeUntil,
  Subject,
} from 'rxjs';

import { ApiService } from '../api';
import { UpdateUserBioRequestBody, UserResponse } from './user.type';

@Injectable({ providedIn: 'root' })
export class UserService {
  readonly loggedInUser = signal<UserResponse | null>(null);
  readonly isLoading = signal<boolean>(true);
  private readonly resetPreviousPoll = new Subject<void>();
  private readonly router = inject(Router);
  private readonly api = inject(ApiService);

  constructor() {
    this.pollOnReadUserMe();
  }

  pollOnReadUserMe(): void {
    this.resetPreviousPoll.next();
    timer(0, 60 * 1000)
      .pipe(
        switchMap(() => this.readUserMe()),
        shareReplay({ bufferSize: 1, refCount: true }),
        takeUntil(this.resetPreviousPoll),
      )
      .subscribe();
  }

  updateUserBio(requestBody: UpdateUserBioRequestBody): Observable<void> {
    return this.api.update('user/me/bio', requestBody)
  }

  private readUserMe(): Observable<UserResponse | null> {
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
}
