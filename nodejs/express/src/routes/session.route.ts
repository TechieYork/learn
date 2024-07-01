import express from "express";
import expressAsyncHandler from "express-async-handler";
import {
  createSession,
  updateSession,
  deleteSession,
} from "../controllers/session.controller";
import {
  createSessionValidator,
  updateSessionValidator,
  deleteSessionValidator,
} from "../validators/session.validator";

const router = express.Router();

router.post("/", createSessionValidator, expressAsyncHandler(createSession));
router.patch("/", updateSessionValidator, expressAsyncHandler(updateSession));
router.delete("/", deleteSessionValidator, expressAsyncHandler(deleteSession));

export default router;
