import { Router } from 'express';
import { handler } from 'src/commons/handler';
import { validateSchema } from 'src/middlewares/validate-schema';
import { getStatisticsSchemas } from '../dto/get-statistics';
import { CTRL, container } from '../module';

const router = Router();

router.post(
  '/get-statistics',
  validateSchema(getStatisticsSchemas),
  handler(container, CTRL.getStatistics),
);

export { router };
