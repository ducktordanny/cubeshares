import { inject, Injector, Signal } from "@angular/core";
import { NavigationEnd, Router } from "@angular/router";
import { toSignal } from '@angular/core/rxjs-interop';

import { filter, map } from 'rxjs';

export interface UrlPathSignalOptions {
  injector?: Injector;
}

export function urlPath(options?: UrlPathSignalOptions): Signal<string> {
  const router = options?.injector ? options.injector.get(Router) : inject(Router);
  return toSignal(
    router.events.pipe(
      filter(event => event instanceof NavigationEnd),
      map(() => router.url),
    ),
    { injector: options?.injector, initialValue: router.url }
  );
}
