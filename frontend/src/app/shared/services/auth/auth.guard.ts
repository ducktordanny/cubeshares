import { inject } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanActivateFn,
  Router,
  RouterStateSnapshot,
} from '@angular/router';
import { UserService } from '../user';
import { map, tap } from 'rxjs';

export const isLoggedInGuard: CanActivateFn = (
  _route: ActivatedRouteSnapshot,
  _state: RouterStateSnapshot,
) => {
  const userService = inject(UserService);
  const router = inject(Router);
  return userService.readUserMe().pipe(
    map((user) => user !== null),
    tap((canActivate) => {
      if (!canActivate) void router.navigate(['/login']);
    }),
  );
};
