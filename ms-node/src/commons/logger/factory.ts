// biome-ignore-all lint/suspicious/noExplicitAny: It is not necessary to specify the type of the message parameter, as it can be of any type.

export interface LoggerFactory {
  info: (message: any) => LoggerFactory;

  error: (message: any) => LoggerFactory;

  warn: (message: any) => LoggerFactory;

  debug: (message: any) => LoggerFactory;

  child: (options: object) => LoggerFactory;
}
