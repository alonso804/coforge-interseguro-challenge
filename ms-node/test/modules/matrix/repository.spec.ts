import { strictEqual } from 'node:assert';
import { describe, it } from 'node:test';
import { StatisticsRepository } from 'src/modules/matrix/repository';
import { MockLogger } from 'test/mock/logger';

const repository = new StatisticsRepository({
  logger: new MockLogger(),
});

const MATRIX = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
];

const DIAGONAL_MATRIX = [
  [1, 0, 0],
  [0, 5, 0],
  [0, 0, 9],
];

describe('Statistics Repository', () => {
  describe('getMax', () => {
    it('should return the max value', async () => {
      const expected = 9;

      const result = await repository.getMax(MATRIX);

      strictEqual(result, expected);
    });
  });

  describe('getMin', () => {
    it('should return the min value', async () => {
      const expected = 1;

      const result = await repository.getMin(MATRIX);

      strictEqual(result, expected);
    });
  });

  describe('getSum', () => {
    it('should return the sum of all values given a matrix', async () => {
      const expected = 45;

      const result = await repository.getSum(MATRIX);

      strictEqual(result, expected);
    });
  });

  describe('getMean', () => {
    it('should return mean value given a matrix', async () => {
      const expected = 5;

      const result = await repository.getMean(MATRIX);

      strictEqual(result, expected);
    });
  });

  describe('isDiagonal', () => {
    it('should return true if matrix is diagonal', async () => {
      const expected = true;

      const result = await repository.isDiagonal(DIAGONAL_MATRIX);

      strictEqual(result, expected);
    });

    it('should return true if matrix is not diagonal', async () => {
      const expected = false;

      const result = await repository.isDiagonal(MATRIX);

      strictEqual(result, expected);
    });
  });
});
