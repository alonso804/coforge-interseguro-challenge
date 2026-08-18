import type { LoggerFactory } from './commons/logger/factory';
import { WinstonLogger } from './commons/logger/winston';

export const logger: LoggerFactory = new WinstonLogger();
