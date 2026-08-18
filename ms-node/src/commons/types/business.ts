import type { RequestHandler } from 'express';

export interface HttpCtrl {
  // biome-ignore lint/suspicious/noExplicitAny: It is not necessary
  run: RequestHandler<any, any, any, any>;
}
