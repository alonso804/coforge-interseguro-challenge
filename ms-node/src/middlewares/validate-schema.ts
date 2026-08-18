import type { NextFunction, Request, Response } from 'express';
import type { z } from 'zod';
import BadRequest from '../commons/errors/bad-request';

export const validateSchema =
  (schema: z.ZodTypeAny) =>
  (req: Request, _res: Response, next: NextFunction): void => {
    const result = schema.safeParse(req);

    if (!result.success) {
      const errors: Record<string, string[]> = {};

      result.error.issues.forEach(({ path, message }) => {
        const key = path.join('.');

        if (!errors[key]) {
          errors[key] = [];
        }

        errors[key].push(message);
      });

      throw new BadRequest(errors);
    }

    // eslint-disable-next-line no-param-reassign
    req = result.data as Request;

    next();
  };
