import type { RequestHandler } from 'express';

export type Empty = unknown;

export type GetHandler<Params, Query, Response> = RequestHandler<
  Params,
  Response,
  unknown,
  Query
>;

export type PostHandler<Body, Response> = RequestHandler<
  unknown,
  Response,
  Body,
  unknown
>;

export type DeleteHandler<Params, Response> = RequestHandler<
  Params,
  Response,
  unknown,
  unknown
>;

export type PutHandler<Params, Body, Response> = RequestHandler<
  Params,
  Response,
  Body,
  unknown
>;

export type PatchHandler<Params, Body, Response> = RequestHandler<
  Params,
  Response,
  Body,
  unknown
>;
