import { NextFunction, Request, Response } from "express";
import { validationResult } from "express-validator";
import logger from "../utils/logger.utils";
import SessionService from "../services/session.service";
import { Errors } from "../errors/errors";
import Result from "../utils/result.utils";

async function createSession(req: Request, res: Response, next: NextFunction) {
  // param check
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    logger.error(
      `create session validation error: ${JSON.stringify(errors.array())}`
    );
    throw Errors.ParamsInvalid;
  }

  // TODO: get user ID from database
  const userId = req.body.username;

  // TODO: verify password

  // create JWT
  let sessions = new SessionService();
  const token = await sessions.create(userId);
  res
    .status(200)
    .json(new Result(Errors.Success.code, Errors.Success.message, { token }));
}

async function updateSession(req: Request, res: Response, next: NextFunction) {
  // get JWT
  let jwt = req.headers.authorization || "";
  if (jwt === "") {
    throw Errors.ParamsInvalid;
  }

  // update JWT
  let sessions = new SessionService();
  const newJwt = await sessions.update(jwt);

  // return new JWT
  let result = new Result(Errors.Success.code, Errors.Success.message, {
    newJwt,
  });
  res.status(200).json(result);
}

async function deleteSession(req: Request, res: Response, next: NextFunction) {
  res.status(200).json({});
}

export { createSession, updateSession, deleteSession };
