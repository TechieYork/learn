import { NextFunction, Request, Response } from "express";
import logger from "../utils/logger.utils";
import UsersService from "../services/users.service";
import { CustomError, Errors } from "../errors/errors";
import Result from "../utils/result.utils";

async function getProfile(req: Request, res: Response, next: NextFunction) {
  // param check
  logger.info(`user ID to get: ${req.params.userId}`);
  if (req.params.userId === "") {
    throw Errors.ParamsInvalid;
  }

  // business code
  let users = new UsersService();
  const profile = await users.getProfile(req.params.userID);
  res
    .status(200)
    .json(new Result(Errors.Success.code, Errors.Success.message, { profile }));
}

export { getProfile };
