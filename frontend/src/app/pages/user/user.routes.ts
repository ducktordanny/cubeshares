import { Route } from '@angular/router';

import { isLoggedInGuard } from '@cubeshares/services/auth';

import { MePageComponent } from './me/me.page';
import { LayoutComponent } from '@cubeshares/components/layout/layout.component';

export const USER_ROUTES: Route[] = [
  {
    path: '',
    component: LayoutComponent,
    children: [
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
    ],
  },
];
