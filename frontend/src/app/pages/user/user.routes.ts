import { Route } from '@angular/router';
import { MePageComponent } from './me/me.page';
import { isLoggedInGuard } from '@cubeshares/shared/services/auth';

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
