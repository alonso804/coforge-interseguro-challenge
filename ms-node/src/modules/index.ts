import { Router } from 'express';
import { handler } from 'src/commons/handler';
import { validateSchema } from 'src/middlewares/validate-schema';
import { CTRL, container } from '../dependencies';
import { operateMatrixSchema } from './operate-matrix';

const router = Router();

router.post(
  '/operate-matrix',
  validateSchema(operateMatrixSchema),
  handler(container, CTRL.operateMatrix),
);

export { router };
