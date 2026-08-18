import { asClass, createContainer } from 'awilix';
import { WinstonLogger } from 'src/commons/logger/winston';
import { GetStatisticsController } from './controllers/get-statistics';
import { StatisticsRepository } from './repository';
import { StatisticsService } from './service';

export const container = createContainer({
  injectionMode: 'PROXY',
});

export const CTRL = {
  getStatistics: 'getStatistics',
} as const;

export type MatrixCtrl = (typeof CTRL)[keyof typeof CTRL];

container.register({
  logger: asClass(WinstonLogger),
});

// Implementations
container.register({
  statisticsRepository: asClass(StatisticsRepository),
});

// Services
container.register({
  statisticsService: asClass(StatisticsService),
});

// Controllers
container.register({
  [CTRL.getStatistics]: asClass(GetStatisticsController),
});
