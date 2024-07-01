import { body } from "express-validator";

export const createSessionValidator = [
  body("username").isString().notEmpty(),
  body("type").isString().notEmpty().isIn(["email", "phone"]),
  body("password").isString().notEmpty(),
];

export const updateSessionValidator = [];

export const deleteSessionValidator = [];
