import z from 'zod';

export const getStatisticsSchemas = z.object({
  body: z.object({
    matrix: z.array(z.array(z.number())).nonempty('Matrix cannot be empty'),
  }),
});

export type GetStatisticsPayload = z.infer<typeof getStatisticsSchemas>['body'];

export type GetStatisticsResponse = {
  max: number;
  min: number;
  mean: number;
  sum: number;
  isDiagonal: boolean;
};
