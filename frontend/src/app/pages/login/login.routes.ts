import { Route } from '@angular/router';

import { LoginPageComponent } from './login.page';

export const LOGIN_ROUTES: Route[] = [
  {
    path: '',
    children: [
      {
        path: '',
        component: LoginPageComponent,
      },
    ],
  },
];
