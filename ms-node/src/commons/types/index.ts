import type { STATUS_CODE } from '../constants/http';

export type StatusCode = (typeof STATUS_CODE)[keyof typeof STATUS_CODE];
