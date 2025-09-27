import { inject, Injectable, signal } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';

import { MessageService } from 'primeng/api';
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
export class UserMeService {
  readonly loggedInUser = signal<UserResponse | null>(null);
  readonly isLoading = signal<boolean>(true);
  private readonly resetPreviousPoll = new Subject<void>();
  private readonly api = inject(ApiService);
  private readonly messageService = inject(MessageService);

  constructor() {
    this.pollReadUserMe();
  }

  pollReadUserMe(): void {
    this.resetPreviousPoll.next();
    timer(0, 60 * 1000)
      .pipe(
        switchMap(() => this.readUserMe()),
        shareReplay({ bufferSize: 1, refCount: true }),
        takeUntil(this.resetPreviousPoll),
      )
      .subscribe();
  }

  readUserMe(): Observable<UserResponse | null> {
    this.isLoading.set(true);
    return this.api
      .read<UserResponse>('user/me')
      .pipe(
        take(1),
        tap(user => this.loggedInUser.set(user)),
        catchError((httpError: HttpErrorResponse) => {
          this.loggedInUser.set(null);
          if (![401].includes(httpError.status)) {
            const { error } = httpError;
            this.messageService.add({ severity: 'error', summary: 'Error', detail: error?.error || error || 'Unknown error' })
          }
          return of(null);
        }),
        finalize(() => this.isLoading.set(false)),
      );
  }

  updateUserBio(requestBody: UpdateUserBioRequestBody): Observable<void> {
    return this.api.update('user/me/bio', requestBody)
  }
}
