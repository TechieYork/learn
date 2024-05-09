import { createLogger, transports, format } from "winston";
import "winston-daily-rotate-file";

const consoleLogFormat = format.combine(
  format.timestamp({ format: "MMM-DD-YYYY HH:mm:ss" }),
  format.printf((i) => `${[i.timestamp]} -> ${i.level}: ${i.message}`)
);

const fileLogFormat = format.combine(format.json(), format.timestamp());

const logger = createLogger({
  transports: [
    // console logger
    new transports.Console(),

    // file logger with rotation
    new transports.DailyRotateFile({
      filename: "logs/app-%DATE%.log",
      level: "info",
      format: fileLogFormat,
      datePattern: "YYYY-MM-DD",
      zippedArchive: true,
      maxSize: "60m",
      maxFiles: "14d",
    }),
  ],
  format: consoleLogFormat,
});

export default logger;
