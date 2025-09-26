import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';

import { api } from '@cubeshares/utils';

import { HttpOptions } from './api.type';

@Injectable({ providedIn: 'root' })
export class ApiService {
  private readonly http = inject(HttpClient);

  read<T>(url: string, options?: HttpOptions<'get'>) {
    return this.http.get<T>(api(url), options);
  }

  create<T>(url: string, body: any | null, options?: HttpOptions<'post'>) {
    return this.http.post<T>(api(url), body, options);
  }
}
