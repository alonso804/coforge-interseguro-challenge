import { STATUS_CODE } from 'src/commons/constants/http';
import type { LoggerFactory } from 'src/commons/logger/factory';
import type { HttpCtrl } from 'src/commons/types/business';
import type { PostHandler } from 'src/commons/types/express';
import type {
  GetStatisticsPayload,
  GetStatisticsResponse,
} from '../dto/get-statistics';
import type { StatisticsService } from '../service';

export class GetStatisticsController implements HttpCtrl {
  #statisticsService: StatisticsService;
  #logger: LoggerFactory;

  constructor(dependencies: {
    statisticsService: StatisticsService;
    logger: LoggerFactory;
  }) {
    this.#statisticsService = dependencies.statisticsService;
    this.#logger = dependencies.logger;
  }

  run: PostHandler<GetStatisticsPayload, GetStatisticsResponse> = async (
    req,
    res,
  ) => {
    const response = await this.#statisticsService.getStatistics(req.body);

    this.#logger.info(`Response: ${JSON.stringify(response)}`);

    res.status(STATUS_CODE.OK_200).json(response);
  };
}
