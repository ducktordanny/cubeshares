import { Routes } from '@angular/router';

export const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: '/user' },
  {
    path: 'login',
    loadChildren: () =>
      import('@cubeshares/pages/login').then((m) => m.LOGIN_ROUTES),
  },
  {
    path: 'user',
    loadChildren: () =>
      import('@cubeshares/pages/user').then((m) => m.USER_ROUTES),
  },
];
