import { HttpClient } from '@angular/common/http';

type HttpMethodsWithOptions =
  | 'get'
  | 'post'
  | 'put'
  | 'patch'
  | 'head'
  | 'options'
  | 'delete';

type HttpMethodWithBody = 'post' | 'put' | 'patch';

export type HttpOptions<T extends HttpMethodsWithOptions> = Parameters<
  HttpClient[T]
>[T extends HttpMethodWithBody ? 2 : 1];
