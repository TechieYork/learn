import express, { Express } from "express";
import morganBody from "morgan-body";
import configs from "./configs/configs";
import constants from "./configs/constants";
import logger from "./utils/logger.utils";
import routes from "./routes/routes";
import sessionRoutes from "./routes/session.route";
import usersRoutes from "./routes/users.route";
import errorLogMiddleware from "./middlewares/errorlog.middleware";
import errorMiddleware from "./middlewares/error.middleware";

// init service
const app: Express = express();
const port = configs.PORT || 3000;
const node_env = configs.NODE_ENV || constants.NODE_ENV_PROD;

// configure middleware
app.use(express.json()); // this will parse your data as JSON format
app.use(express.urlencoded({ extended: true }));
if (node_env !== "prod") {
  morganBody(app); // NOTE: this morgan body middleware should be called before register routes
  app.use(errorLogMiddleware);
}

// configure routes
app.use("/", routes);
app.use("/session", sessionRoutes);
app.use("/users", usersRoutes);

// configure error handling, this must be configured at the end after the routes
app.use(errorMiddleware);

// run service
app.listen(port, () => {
  logger.info(`Server is running at port:${port}`);
});
