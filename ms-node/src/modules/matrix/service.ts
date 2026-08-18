import type { LoggerFactory } from 'src/commons/logger/factory';
import type {
  GetStatisticsPayload,
  GetStatisticsResponse,
} from './dto/get-statistics';
import type { StatisticsRepository } from './repository';

export class StatisticsService {
  #statisticsRepository: StatisticsRepository;
  #logger: LoggerFactory;

  constructor(dependencies: {
    statisticsRepository: StatisticsRepository;
    logger: LoggerFactory;
  }) {
    this.#statisticsRepository = dependencies.statisticsRepository;
    this.#logger = dependencies.logger;
  }

  async getStatistics(
    payload: GetStatisticsPayload,
  ): Promise<GetStatisticsResponse> {
    this.#logger.info(`Payload: ${JSON.stringify(payload)}`);

    const [max, min, mean, sum, isDiagonal] = await Promise.all([
      this.#statisticsRepository.getMax(payload.matrix),
      this.#statisticsRepository.getMin(payload.matrix),
      this.#statisticsRepository.getMean(payload.matrix),
      this.#statisticsRepository.getSum(payload.matrix),
      this.#statisticsRepository.isDiagonal(payload.matrix),
    ]);

    return {
      max,
      min,
      mean,
      sum,
      isDiagonal,
    };
  }
}
