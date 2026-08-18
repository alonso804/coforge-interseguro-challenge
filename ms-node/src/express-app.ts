import { createServer, type Server as HttpServer } from 'node:http';
import cors from 'cors';
import express, { type Express } from 'express';
import errorHandler from './middlewares/error-handler';
import incomeLog from './middlewares/income-log';
import { router } from './modules/matrix/controllers';

export const expressApp = (app: Express): HttpServer => {
  const server = createServer(app);

  app.use(express.json());
  app.use(express.urlencoded({ extended: true, limit: '1mb' }));
  app.use(cors());

  app.use(incomeLog);

  // API
  app.use('/api', router);

  app.get('/health-check', (_req, res) => {
    res.status(200).send('[MS-NODE] Health check passed!');
  });

  app.use(errorHandler);

  return server;
};
