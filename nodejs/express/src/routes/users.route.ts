import express from "express";
import expressAsyncHandler from "express-async-handler";
import { getProfile } from "../controllers/users.controller";

const router = express.Router();

router.get("/:userId", expressAsyncHandler(getProfile));

export default router;
