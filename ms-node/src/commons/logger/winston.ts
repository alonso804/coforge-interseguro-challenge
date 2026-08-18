// biome-ignore-all lint/suspicious/noExplicitAny: It is not necessary to specify the type of the message parameter, as it can be of any type.

import type { Logger } from 'winston';
import { createLogger, format, transports } from 'winston';
import type { LoggerFactory } from './factory';

const devLogger = createLogger({
  level: 'debug',
  format: format.combine(
    format.timestamp(),
    format.json(),
    format.prettyPrint(),
    format.colorize({ all: true }),
  ),
  transports: [new transports.Console()],
});

const prodLogger = createLogger({
  level: 'info',
  format: format.combine(format.timestamp(), format.json()),
  transports: [new transports.Console()],
});

const logger = process.env.NODE_ENV === 'local' ? devLogger : prodLogger;

export class WinstonLogger implements LoggerFactory {
  #logger: Logger;

  constructor() {
    this.#logger = logger;
  }

  info = (message: any) => this.#logger.info(message);
  error = (message: any) => this.#logger.error(message);
  warn = (message: any) => this.#logger.warn(message);
  debug = (message: any) => this.#logger.debug(message);
  child = (options: object) => this.#logger.child(options);
}
