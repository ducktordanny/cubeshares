import { Route } from '@angular/router';

import { isLoggedInGuard } from '@cubeshares/services/auth';

import { MePageComponent } from './me/me.page';

export const USER_ROUTES: Route[] = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'me',
  },
  {
    path: 'me',
    canActivate: [isLoggedInGuard],
    component: MePageComponent,
  },
];
