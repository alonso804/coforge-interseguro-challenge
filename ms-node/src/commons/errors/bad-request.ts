import { STATUS_CODE } from '../constants/http';
import BaseError from './base-error';

export default class BadRequest extends BaseError {
  constructor(errors: Record<string, string[]>) {
    super(STATUS_CODE.BAD_REQUEST_400, JSON.stringify(errors));
  }
}
