import type { LoggerFactory } from 'src/commons/logger/factory';

export class StatisticsRepository {
  #logger: LoggerFactory;

  constructor(dependencies: {
    logger: LoggerFactory;
  }) {
    this.#logger = dependencies.logger;
  }

  async getMin(matrix: number[][]): Promise<number> {
    return Math.min(...matrix.flat());
  }

  async getMax(matrix: number[][]): Promise<number> {
    return Math.max(...matrix.flat());
  }

  async getSum(matrix: number[][]): Promise<number> {
    return matrix.flat().reduce((acc, val) => acc + val, 0);
  }

  async getMean(matrix: number[][]): Promise<number> {
    const count = matrix.flat().length;
    const sum = await this.getSum(matrix);

    return sum / count;
  }

  async isDiagonal(matrix: number[][]): Promise<boolean> {
    for (let i = 0; i < matrix.length; i++) {
      for (let j = 0; j < matrix[i].length; j++) {
        if (i !== j && matrix[i][j] !== 0) {
          return false;
        }
      }
    }

    return true;
  }
}
