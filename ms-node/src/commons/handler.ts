import { randomUUID } from 'node:crypto';
import { type AwilixContainer, asValue } from 'awilix';
import type { NextFunction, Request, RequestHandler, Response } from 'express';
import type { MatrixCtrl } from 'src/modules/matrix/module';
import type { LoggerFactory } from './logger/factory';
import type { HttpCtrl } from './types/business';

type Instance = MatrixCtrl;

const getControllerInstance = (
  container: AwilixContainer,
  instance: Instance,
  req: Request,
  res: Response,
  next: NextFunction,
) => {
  const logger = container.resolve<LoggerFactory>('logger');

  const requestLogger = logger.child({
    traceId: req.headers['x-trace-id'] ?? randomUUID(),
  });

  const requestContainer = container.createScope();

  requestContainer.register({
    logger: asValue(requestLogger),
  });

  const controllerInstance = requestContainer.resolve<HttpCtrl>(instance);

  return (controllerInstance.run(req, res, next) as Promise<void>).catch(next);
};

export const handler =
  (container: AwilixContainer, instance: Instance): RequestHandler =>
  async (req, res, next) => {
    return getControllerInstance(container, instance, req, res, next);
  };
