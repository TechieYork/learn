import { createSecretKey } from "crypto";
import { v4 as uuidv4 } from "uuid";
import * as jose from "jose";
import configs from "../configs/configs";
import { Errors } from "../errors/errors";
import logger from "../utils/logger.utils";

const secretKey = createSecretKey(configs.JWT_SECRET ?? "", "utf-8");

class SessionService {
  async create(id: string): Promise<string> {
    // generate new jwt
    const token = await this.generateToken(id);
    logger.info(`Token created for user: ${id}, jwt: ${token}`);
    return token;
  }

  async update(jwt: string): Promise<string> {
    // verify jwt
    const { id, tk } = await this.verifyToken(jwt);

    // generate new jwt with the same id and tk
    const newJwt = await this.generateToken(id, tk);
    logger.info(`Token updated for user: ${id}, jwt: ${newJwt}`);
    return newJwt;
  }

  async delete(id: string): Promise<void> {
    // DO NOTHING HERE
    // in the future, you can save the token to a database
    // so you can delete it by setting the expiration time
    logger.info(`Token deleted for user: ${id}`);
  }

  async verify(jwt: string): Promise<void> {
    const { id, tk } = await this.verifyToken(jwt);
    logger.info(`Token verified for user: ${id}`);
  }

  async generateToken(id: string, tk: string = ""): Promise<string> {
    configs.JWT_ISSUER = configs.JWT_ISSUER ?? "MOOW";
    configs.JWT_AUDIENCE = configs.JWT_AUDIENCE ?? "WebApp";
    configs.JWT_EXPIRATION_TIME = configs.JWT_EXPIRATION_TIME ?? "3h";

    const jwt = await new jose.SignJWT({
      id: id,
      tk: tk || uuidv4(),
    }) // details to  encode in the token
      .setProtectedHeader({
        alg: "HS256",
      }) // algorithm
      .setIssuedAt()
      .setIssuer(configs.JWT_ISSUER) // issuer with nullish coalescing operator
      .setAudience(configs.JWT_AUDIENCE) // audience
      .setExpirationTime(configs.JWT_EXPIRATION_TIME) // token expiration time, e.g., "1 day"
      .sign(secretKey); // secretKey generated from previous step
    return jwt;
  }

  async verifyToken(jwt: string): Promise<{ id: string; tk: string }> {
    try {
      // verify token
      jwt = jwt.replace("Bearer ", "");
      const { payload, protectedHeader } = await jose.jwtVerify(
        jwt,
        secretKey,
        {
          issuer: configs.JWT_ISSUER, // issuer
          audience: configs.JWT_AUDIENCE, // audience
        }
      );

      return { id: payload.id as string, tk: payload.tk as string };
    } catch (error) {
      // jwt verification failed
      logger.warn(`JWT is invalid, error: ${error}`);
      if (error instanceof jose.errors.JWTExpired) {
        throw Errors.Expired;
      }
      throw Errors.AuthFailed;
    }
  }
}

export default SessionService;
