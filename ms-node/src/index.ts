import express from 'express';
import { ENV } from './env';
import { expressApp } from './express-app';

function startServer() {
  const app = express();

  const server = expressApp(app);

  app
    .listen(ENV.PORT, () => {
      console.log(`🚀 Server is running on port ${ENV.PORT}`);
    })
    .on('error', (err) => {
      console.error('Error starting server:', err);
      process.exit(1);
    });

  return server;
}

startServer();
