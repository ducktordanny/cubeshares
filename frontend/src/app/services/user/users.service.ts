import { inject, Injectable, OnDestroy, signal } from "@angular/core";
import { HttpErrorResponse } from "@angular/common/http";

import { MessageService } from "primeng/api";
import { catchError, finalize, Observable, of, shareReplay, Subject, switchMap, take, takeUntil, tap, timer } from "rxjs";

import { UserResponse } from "./user.type";
import { ApiService } from "../api";

@Injectable()
export class UsersService implements OnDestroy {
  readonly user = signal<UserResponse | null>(null);
  readonly isLoading = signal<boolean>(false);
  private readonly resetPreviousPoll = new Subject<void>();
  private readonly api = inject(ApiService);
  private readonly messageService = inject(MessageService);

  pollReadUserById(id: number): void {
    this.resetPreviousPoll.next();
    timer(0, 60 * 1000)
      .pipe(
        switchMap(() => this.readUserById(id)),
        shareReplay({ bufferSize: 1, refCount: true }),
        takeUntil(this.resetPreviousPoll),
      )
      .subscribe();
  }

  readUserById(id: number): Observable<UserResponse | null> {
    this.isLoading.set(true);
    return this.api.read<UserResponse>(`user/${id}`)
      .pipe(
        take(1),
        tap(user => this.user.set(user)),
        catchError((httpError: HttpErrorResponse) => {
          this.user.set(null);
          const { error } = httpError;
          this.messageService.add({ severity: 'error', summary: 'Error', detail: error?.error || error || 'Unknown error' })
          return of(null);
        }),
        finalize(() => this.isLoading.set(false)),
      )
  }

  ngOnDestroy(): void {
    this.resetPreviousPoll.next();
    this.resetPreviousPoll.complete();
  }
}
