import { inject } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import {
  ActivatedRouteSnapshot,
  CanActivateFn,
  Router,
  RouterStateSnapshot,
} from '@angular/router';

import { filter, map, tap } from 'rxjs';

import { UserService } from '../user';

export const isLoggedInGuard: CanActivateFn = (
  _route: ActivatedRouteSnapshot,
  _state: RouterStateSnapshot,
) => {
  const userService = inject(UserService);
  const router = inject(Router);

  return toObservable(userService.loggedInUser).pipe(
    filter(() => !userService.isLoading()),
    map((user) => user !== null),
    tap((canActivate) => {
      if (!canActivate) void router.navigate(['/login']);
    }),
  );
};
