// biome-ignore-all lint/suspicious/noExplicitAny: It is not necessary to specify the type of the message parameter, as it can be of any type.
import { strictEqual } from 'node:assert';
import { afterEach, beforeEach, describe, it, mock } from 'node:test';
import { StatisticsRepository } from 'src/modules/matrix/repository';
import { StatisticsService } from 'src/modules/matrix/service';
import { MockLogger } from 'test/mock/logger';

const repository = new StatisticsRepository({
  logger: new MockLogger(),
});

const service = new StatisticsService({
  statisticsRepository: repository,
  logger: new MockLogger(),
});

const MATRIX = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
];

describe('Statistics Service', () => {
  let minSpy: ReturnType<typeof mock.method>;
  let maxSpy: ReturnType<typeof mock.method>;
  let meanSpy: ReturnType<typeof mock.method>;
  let sumSpy: ReturnType<typeof mock.method>;
  let isDiagonalSpy: ReturnType<typeof mock.method>;

  beforeEach(() => {
    minSpy = mock.method(repository, 'getMin');
    maxSpy = mock.method(repository, 'getMax');
    meanSpy = mock.method(repository, 'getMean');
    sumSpy = mock.method(repository, 'getSum');
    isDiagonalSpy = mock.method(repository, 'isDiagonal');
  });

  afterEach(() => {
    minSpy.mock.resetCalls();
    maxSpy.mock.resetCalls();
    meanSpy.mock.resetCalls();
    sumSpy.mock.resetCalls();
    isDiagonalSpy.mock.resetCalls();
  });

  it('should call all repository methods once', async () => {
    await service.getStatistics({ matrix: MATRIX });

    strictEqual(minSpy.mock.callCount(), 1);
    strictEqual(maxSpy.mock.callCount(), 1);
    strictEqual(meanSpy.mock.callCount(), 1);
    strictEqual(sumSpy.mock.callCount(), 2); // It is 2 because mean uses sum internally
    strictEqual(isDiagonalSpy.mock.callCount(), 1);
  });

  it('should throw an error if any repository method fails', async () => {
    minSpy.mock.mockImplementationOnce(() => {
      throw new Error('Repository error');
    });

    try {
      await service.getStatistics({ matrix: MATRIX });
    } catch (error) {
      if (error instanceof Error) {
        strictEqual(error.message, 'Repository error');
      }
    }
  });
});
