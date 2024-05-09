import { Request, Response } from "express";
import Result from "../utils/result.utils";
import { Errors, CustomError } from "../errors/errors";

class Controller {
  constructor() {}
  success(res: Response, data: object) {
    const result = new Result(
      Errors.Success.code,
      Errors.Success.message,
      data
    );
    res.status(200).json(result);
  }

  fail(res: Response, error: CustomError, message: string = "") {
    message = message || error.message;
    const result = new Result(error.code, message, {});
    res.status(200).json(result);
  }
}

export { Controller };
