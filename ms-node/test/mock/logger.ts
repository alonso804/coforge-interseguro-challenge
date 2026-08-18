// biome-ignore-all lint/suspicious/noExplicitAny: It is not necessary to specify the type of the message parameter, as it can be of any type.

import type { LoggerFactory } from 'src/commons/logger/factory';

export class MockLogger implements LoggerFactory {
  info = (message: any) => {
    console.log(message);

    return this;
  };

  error = (message: any) => {
    console.error(message);

    return this;
  };

  warn = (message: any) => {
    console.warn(message);

    return this;
  };

  debug = (message: any) => {
    console.debug(message);

    return this;
  };

  child = (_options: object) => {
    return this;
  };
}
