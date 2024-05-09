import express from "express";
import expressAsyncHandler from "express-async-handler";
import { Errors } from "../errors/errors";

const router = express.Router();

router.get(
  "/",
  expressAsyncHandler(async (req, res) => {
    throw Errors.NotFound;
  })
);

export default router;
