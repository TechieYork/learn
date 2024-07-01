import dotenv from "dotenv";

// get config
dotenv.config({ path: "" });

// check if the proper values are set
if (!process.env.PORT) {
  throw new Error("PORT is not set in .env file");
}

if (!process.env.NODE_ENV) {
  throw new Error("NODE_ENV is not set in .env file");
}

if (!process.env.JWT_SECRET) {
  throw new Error("One of the JWT settings is not set in .env file");
}

export default process.env;
