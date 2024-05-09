import { Request, Response, NextFunction } from "express";
import { CustomError } from "../errors/errors";
import Result from "../utils/result.utils";

// error middleware
//   this middleware will be called when there is an error in the business code 
//   and return 200 OK with business error in JSON body
const errorMiddleware = (
  error: CustomError,
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const result = new Result(
    error.code,
    error.message,
    {}
  );
  res.status(200).json(result);
};

export default errorMiddleware;
