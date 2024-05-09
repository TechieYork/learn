import { UserProfile, UserProfileOptions, UserStorage } from "../models/users.model";
import { Errors } from "../errors/errors";

class UsersService {
  async getProfile(id: string): Promise<UserProfile | undefined> {
    let storage = new UserStorage();
    // Your implementation here
    return storage.get(id)
  }
}

export default UsersService;
