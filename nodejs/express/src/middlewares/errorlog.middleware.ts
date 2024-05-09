import { Request, Response, NextFunction } from "express";
import logger from "../utils/logger.utils";

// error log middleware
//   this middleware will be called when there is an error in the business code 
//   and log the request detail
const errorLogMiddleware = (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  const originalJson = res.json;
  res.json = function (body: any): Response<any, Record<string, any>> {
    // log error if code is not 0
    if (body.code !== 0) {
      logger.error(
        `code: ${body.code}, msg: ${body.message}, req: ${JSON.stringify(
          req.body
        )}, res: ${JSON.stringify(body)}`
      );
    }
    return originalJson.call(this, body);
  };

  next();
};

export default errorLogMiddleware;
