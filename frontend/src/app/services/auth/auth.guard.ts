import { computed, inject } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import {
  ActivatedRouteSnapshot,
  CanActivateFn,
  Router,
  RouterStateSnapshot,
} from '@angular/router';

import { filter, map, tap } from 'rxjs';

import { UserMeService } from '../user';

const accessState = () => {
  const service = inject(UserMeService);
  return computed(() => ({
    loggedInUser: service.loggedInUser(),
    isLoading: service.isLoading(),
    error: service.error(),
  }));
}


export const isLoggedInGuard: CanActivateFn = (
  _route: ActivatedRouteSnapshot,
  _state: RouterStateSnapshot,
) => {
  const router = inject(Router);

  return toObservable(accessState()).pipe(
    filter(state => !state.isLoading),
    map(state => state.loggedInUser !== null && state.error === null),
    tap(access => {
      if (!access && !router.url.startsWith('/login')) void router.navigate(['/login']);
    }),
  );
};
